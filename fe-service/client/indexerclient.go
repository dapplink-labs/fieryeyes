package client

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/savour-labs/fieryeyes/fe-service/services/common"
	"github.com/savour-labs/fieryeyes/indexer/protobuf"
	"google.golang.org/grpc"
	"time"
)

type IndexerClientConfig struct {
	IndexerSocket string `json:"indexer_socket"`
}

type IndexerClient struct {
	Ctx    context.Context
	Cfg    *IndexerClientConfig
	Client protobuf.IndexerRpcServiceClient
	Cancel func()
}

func NewIndexerClient(cfg *IndexerClientConfig) (*IndexerClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(common.DefaultTimeout))
	defer cancel()
	conn, err := grpc.Dial(cfg.IndexerSocket, grpc.WithInsecure())
	if err != nil {
		log.Error("Cannot connect to index", "IndexerSocket", cfg.IndexerSocket)
		return nil, err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	return &IndexerClient{
		Ctx:    ctx,
		Cfg:    cfg,
		Client: protobuf.NewIndexerRpcServiceClient(conn),
		Cancel: cancel,
	}, nil
}

func (ic IndexerClient) GetLatestBlock() (*protobuf.LatestBlockRep, error) {
	return ic.Client.GetLatestBlock(ic.Ctx, nil)
}
