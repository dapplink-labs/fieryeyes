package client

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/savour-labs/fieryeyes/fe-scrapy/protobuf"
	"github.com/savour-labs/fieryeyes/fe-service/services/common"
	"google.golang.org/grpc"
	"time"
)

type ScrapyClientConfig struct {
	ScrapySocket string `json:"scrapy_socket"`
}

type ScrapyClient struct {
	Ctx    context.Context
	Cfg    *ScrapyClientConfig
	Client protobuf.GiantWhaleServiceClient
	Cancel func()
}

func NewScrapyClient(cfg *ScrapyClientConfig) (*ScrapyClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(common.DefaultTimeout))
	defer cancel()
	conn, err := grpc.Dial(cfg.ScrapySocket, grpc.WithInsecure())
	if err != nil {
		log.Error("Cannot connect to fe Scrapy", "ScrapySocket", cfg.ScrapySocket)
		return nil, err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	return &ScrapyClient{
		Ctx:    ctx,
		Cfg:    cfg,
		Client: protobuf.NewGiantWhaleServiceClient(conn),
		Cancel: cancel,
	}, nil
}
