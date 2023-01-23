package chain

import (
	"context"
	"crypto/tls"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/savour-labs/fieryeyes/indexer/db"
	"github.com/savour-labs/fieryeyes/indexer/models"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	DefaultTimeout = 5 * time.Second
)

type EvmChainConfig struct {
	EthRpc          string
	DisableHTTP2    bool
	SyncBlockHeight uint64
	LoopInterval    time.Duration
	Database        *db.Database
}

type EvmChainClient struct {
	Ctx       context.Context
	EthClient *ethclient.Client
	Cfg       *EvmChainConfig
	cancel    func()
	wg        sync.WaitGroup
}

func NewEvmChainClient(ctx context.Context, cfg *EvmChainConfig) (*EvmChainClient, error) {
	ctxt, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()
	var ethClient *ethclient.Client
	if strings.HasPrefix(cfg.EthRpc, "http") {
		httpClient := new(http.Client)
		if cfg.DisableHTTP2 {
			log.Info("Disabled HTTP/2 support in L1 eth client")
			httpClient.Transport = &http.Transport{
				TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
			}
		}
		rpcClient, err := rpc.DialHTTPWithClient(cfg.EthRpc, httpClient)
		if err != nil {
			return nil, err
		}
		ethClient = ethclient.NewClient(rpcClient)
	}
	ethClient, _ = ethclient.DialContext(ctxt, cfg.EthRpc)
	return &EvmChainClient{
		Ctx:       ctx,
		EthClient: ethClient,
		Cfg:       cfg,
		cancel:    cancel,
	}, nil
}

func (ecc *EvmChainClient) InitBlock() error {
	var blocks models.Blocks
	latestBlock, err := ecc.EthClient.BlockNumber(ecc.Ctx)
	log.Info("latestBlock", "latestBlock", latestBlock)
	if err != nil {
		log.Error("get latest block number fail", "err", err)
		return err
	}
	if !blocks.ExistBlock(ecc.Cfg.Database.Db) {
		initBlock := &models.Blocks{
			BlockHeight:       ecc.Cfg.SyncBlockHeight,
			LatestBlockHeight: latestBlock,
		}
		err = initBlock.SelfInsert(ecc.Cfg.Database.Db)
		if err != nil {
			log.Error("insert block fail", "err", err)
			return err
		}
	}
	return nil
}

func (ecc *EvmChainClient) ProcessBlock(block *types.Block) error {
	for _, tx := range block.Transactions() {
		receipt, err := ecc.EthClient.TransactionReceipt(ecc.Ctx, tx.Hash())
		if err != nil {
			log.Error("get transaction fail", "err", err)
		}
		for _, rLog := range receipt.Logs {
			err = ecc.ProcessTransactionEvent(rLog, receipt.Status)
			if err != nil {
				log.Error("process transaction event fail", "err", err)
			}
		}
		err = ecc.ProcessTransaction(tx, block.Time(), block.Number(), receipt.Status)
		if err != nil {
			log.Error("process transaction fail", "err", err)
		}
	}
	return nil
}

func (ecc *EvmChainClient) ProcessTransactionEvent(rLog *types.Log, status uint64) error {
	log.Info("ProcessTransactionEvent", "address", rLog.Address, "data", rLog.Data)
	event := &models.Events{
		Address:     rLog.Address.String(),
		Data:        "",
		BlockNumber: rLog.BlockNumber,
		TxHash:      rLog.TxHash.String(),
		TxIndex:     rLog.TxIndex,
		BlockHash:   rLog.BlockHash.String(),
		LogIndex:    rLog.Index,
		Removed:     rLog.Removed,
	}
	err := event.SelfInsert(ecc.Cfg.Database.Db)
	if err != nil {
		log.Error("insert event fail", "err", err)
		return err
	}
	evt, err := event.GetEventByTxHash(ecc.Cfg.Database.Db)
	if err != nil {
		log.Error("get event by tx hash fail", "err", err)
		return err
	}
	log.Info("Topics", "topic", rLog.Topics)
	for _, tp := range rLog.Topics {
		topic := &models.Topics{
			EventId: evt.Id,
			Topic:   tp.String(),
		}
		err := topic.SelfInsert(ecc.Cfg.Database.Db)
		if err != nil {
			log.Error("insert topic fail", "err", err)
			return err
		}
	}
	return nil
}

func (ecc *EvmChainClient) ProcessTransaction(tx *types.Transaction, blockTime uint64, blockNumber *big.Int, status uint64) error {
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		log.Error("Failed to read the sender address", "TxHash", tx.Hash(), "err", err)
		return err
	}
	log.Info("hand transaction", "txHash", tx.Hash().String())
	transaction := &models.Transaction{
		BlockNumber: blockNumber.Uint64(),
		TxHash:      tx.Hash().Hex(),
		From:        from.Hex(),
		Value:       tx.Value().String(),
		Status:      status,
		Timestamp:   time.Unix(int64(blockTime), 0),
		InputData:   tx.Data(),
	}
	if tx.To() == nil {
		log.Info("Contract creation found", "Sender", transaction.From, "TxHash", transaction.TxHash)
		toAddress := crypto.CreateAddress(from, tx.Nonce()).Hex()
		transaction.Contract = toAddress
		token := &models.Token{
			Address: toAddress,
		}
		err = token.SelfInsert(ecc.Cfg.Database.Db)
		if err != nil {
			log.Error("insert transaction fail", "err", err)
			return err
		}
	} else {
		to := tx.To().String()
		token := &models.Token{
			Address: to,
		}
		ok := token.ExistToken(ecc.Cfg.Database.Db)
		if ok { // to is contract, call contract
			token.TotalTransactions += 1
			err = token.SelfUpdate(ecc.Cfg.Database.Db)
			if err != nil {
				log.Error("token update fail", "err", err)
				return err
			}
		} else { // to is not contract, normal transaction
			address := &models.Addresses{
				Address: to,
			}
			ok := address.ExistAddress(ecc.Cfg.Database.Db)
			if !ok {
				newAddress := models.Addresses{
					Address: tx.To().String(),
					Balance: tx.Value().String(),
				}
				err := newAddress.SelfInsert(ecc.Cfg.Database.Db)
				if err != nil {
					log.Error("address insert fail", "err", err)
					return err
				}
			} else {
				bigInt := new(big.Int)
				sBalance, ok := bigInt.SetString(address.Balance, 10)
				if !ok {
					log.Error("string to big int fail", "err", err)
					return err
				}
				newBalance := new(big.Int).Add(sBalance, tx.Value())
				address.Balance = newBalance.String()
				err := address.SelfUpdate(ecc.Cfg.Database.Db)
				if err != nil {
					log.Error("update balance fail", "err", err)
					return err
				}
			}
		}
		transaction.Contract = tx.To().Hex()
	}
	err = transaction.SelfInsert(ecc.Cfg.Database.Db)
	if err != nil {
		log.Error("insert transaction fail", "err", err)
		return err
	}
	return nil
}

func (ecc *EvmChainClient) Start() error {
	ecc.wg.Add(1)
	go ecc.SyncLoop()
	return nil
}

func (ecc *EvmChainClient) Stop() {
	ecc.cancel()
	ecc.wg.Wait()
}

func (ecc *EvmChainClient) SyncLoop() {
	defer ecc.wg.Done()
	ticker := time.NewTicker(ecc.Cfg.LoopInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			log.Info("start sync task")
			var blocks models.Blocks
			latestBlock, err := ecc.EthClient.BlockNumber(ecc.Ctx)
			if err != nil {
				log.Error("get latest block number fail", "err", err)
				return
			}
			block, err := blocks.GetFirstColumn(ecc.Cfg.Database.Db)
			if err != nil {
				log.Error("get db block number fail", "err", err)
				return
			}
			if block.BlockHeight >= latestBlock {
				log.Info("chain latest block is equal db block", "chain block", latestBlock, "db block", block.LatestBlockHeight)
				continue
			}
			tBlock, err := ecc.EthClient.BlockByNumber(ecc.Ctx, big.NewInt(int64(block.BlockHeight)))
			if err != nil {
				log.Error("get block by number fail", "err", err)
			}
			log.Info("tBlock", "blockNumber", tBlock.Number(), "blockHash", tBlock.Hash())

			err = ecc.ProcessBlock(tBlock)
			if err != nil {
				log.Info("process block fail", "err", err)
				return
			}
			block.BlockHash = tBlock.Hash().String()
			err = block.SelfUpdate(ecc.Cfg.Database.Db)
			if err != nil {
				log.Error("update block hash fail", "err", err)
			}
			newBlock := models.Blocks{
				BlockHeight:       block.BlockHeight + 1,
				BlockHash:         "",
				ParentHash:        tBlock.ParentHash().String(),
				LatestBlockHeight: latestBlock,
			}
			err = newBlock.SelfInsert(ecc.Cfg.Database.Db)
			if err != nil {
				log.Error("insert new block fail", "err", err)
			}
		case err := <-ecc.Ctx.Done():
			log.Error("Sync loop exit, fail reason", "err", err)
			return
		}
	}
}
