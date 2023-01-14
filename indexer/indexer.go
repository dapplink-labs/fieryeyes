package indexer

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	chain "github.com/savour-labs/fieryeyes/indexer/blockchain"
	"github.com/savour-labs/fieryeyes/indexer/db"
	"github.com/savour-labs/fieryeyes/indexer/metrics"
	"github.com/savour-labs/fieryeyes/indexer/services"
	"github.com/urfave/cli"
	"os"
	"strconv"
)

func Main(gitVersion string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		cfg, err := NewConfig(ctx)
		if err != nil {
			return err
		}

		log.Info("Initializing indexer")

		indexer, err := NewIndexer(cfg)
		if err != nil {
			log.Error("Unable to create indexer", "error", err)
			return err
		}

		log.Info("Starting indexer")

		if err := indexer.Start(); err != nil {
			return err
		}
		defer indexer.Stop()

		log.Info("Indexer started")

		<-(chan struct{})(nil)

		return nil
	}
}

type Indexer struct {
	ctx                context.Context
	cfg                Config
	ethClient          *chain.EvmChainClient
	indexerRpcServices *services.IndexerRPCServices
	db                 *db.Database
	metrics            *metrics.Metrics
}

func NewIndexer(cfg Config) (*Indexer, error) {
	ctx := context.Background()

	var logHandler log.Handler
	if cfg.LogTerminal {
		logHandler = log.StreamHandler(os.Stdout, log.TerminalFormat(true))
	} else {
		logHandler = log.StreamHandler(os.Stdout, log.JSONFormat())
	}

	logLevel, err := log.LvlFromString(cfg.LogLevel)
	if err != nil {
		return nil, err
	}
	log.Root().SetHandler(log.LvlFilterHandler(logLevel, logHandler))

	dbConfig := &db.DatabaseConfig{
		Username: cfg.DbUsername,
		Password: cfg.DbPassword,
		Host:     cfg.DbHost,
		Port:     cfg.DbPort,
		DbName:   cfg.DbName,
	}
	dbSelf, err := db.NewDatabase(ctx, dbConfig)
	if err != nil {
		log.Error("new database fail", "err", err)
		return nil, err
	}
	err = dbSelf.MigrateDb()
	if err != nil {
		log.Error("migrate db fail", "err", err)
		return nil, err
	}
	evmConfig := &chain.EvmChainConfig{
		EthRpc:          cfg.EthRpc,
		DisableHTTP2:    cfg.DisableHTTP2,
		SyncBlockHeight: cfg.SyncBlockHeight,
		LoopInterval:    cfg.LoopInterval,
		Database:        dbSelf,
	}
	ethCli, err := chain.NewEvmChainClient(ctx, evmConfig)
	if err != nil {
		log.Error("new evm chain client fail", "err", err)
		return nil, err
	}
	err = ethCli.InitBlock()
	if err != nil {
		log.Error("init block fail", "err", err)
		return nil, err
	}

	m := metrics.NewMetrics(nil)

	if cfg.MetricsServerEnable {
		go func() {
			_, err := m.Serve(cfg.MetricsHostname, cfg.MetricsPort)
			if err != nil {
				log.Error("metrics server failed to start", "err", err)
			}
		}()
		log.Info("metrics server enabled", "host", cfg.MetricsHostname, "port", cfg.MetricsPort)
	}

	iRpcConfig := &services.IndexerRPCConfig{
		RpcHost:  cfg.RpcHost,
		RpcPort:  strconv.FormatUint(cfg.RpcPort, 10),
		Database: dbSelf,
	}
	iRpcServices, err := services.NewIndexerRPCServices(ctx, iRpcConfig)
	if err != nil {
		log.Error("new indexer rpc services fail", "err", err)
		return nil, err
	}

	return &Indexer{
		ctx:                ctx,
		cfg:                cfg,
		ethClient:          ethCli,
		indexerRpcServices: iRpcServices,
		db:                 dbSelf,
		metrics:            m,
	}, nil

}

func (i Indexer) Start() error {
	log.Info("indexer start success")
	i.ethClient.Start()
	i.indexerRpcServices.Start()
	return nil
}

func (i Indexer) Stop() {
	i.ethClient.Stop()
	i.indexerRpcServices.Stop()
	log.Info("indexer stop success")
}
