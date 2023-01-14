package services

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	chain "github.com/savour-labs/fieryeyes/indexer/blockchain"
	"github.com/savour-labs/fieryeyes/indexer/db"
	"github.com/savour-labs/fieryeyes/indexer/models"
	"github.com/savour-labs/fieryeyes/indexer/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	"runtime/debug"
	"strings"
	"sync"
)

type IIndexerRpcServices interface {
	GetLatestBlock(ctx context.Context, req *protobuf.LatestBlock) (*protobuf.LatestBlockRep, error)
}

type CommonRequest interface {
	GetConsumerToken() string
}

type IndexerRPCConfig struct {
	RpcHost  string
	RpcPort  string
	Database *db.Database
}

type IndexerRPCServices struct {
	Ctx    context.Context
	Cfg    *IndexerRPCConfig
	cancel func()
	wg     sync.WaitGroup
	IIndexerRpcServices
}

func NewIndexerRPCServices(ctx context.Context, cfg *IndexerRPCConfig) (*IndexerRPCServices, error) {
	ctxt, cancel := context.WithTimeout(ctx, chain.DefaultTimeout)
	defer cancel()
	return &IndexerRPCServices{
		Ctx:    ctxt,
		Cfg:    cfg,
		cancel: cancel,
	}, nil
}

func (rpc *IndexerRPCServices) GetLatestBlock(ctx context.Context, req *protobuf.LatestBlock) (*protobuf.LatestBlockRep, error) {
	var blocks models.Blocks
	block, err := blocks.GetFirstColumn(rpc.Cfg.Database.Db)
	if err != nil {
		log.Error("get db block number fail", "err", err)
		return &protobuf.LatestBlockRep{
			Code: protobuf.ReturnCode_SUCCESS,
			Msg:  "get latest block number fail",
		}, nil
	}
	return &protobuf.LatestBlockRep{
		Code:        protobuf.ReturnCode_SUCCESS,
		Msg:         "request success",
		BlockNumber: block.LatestBlockHeight,
	}, nil
}

func (rpc *IndexerRPCServices) interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Error("panic error", "msg", e)
			log.Debug(string(debug.Stack()))
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()
	pos := strings.LastIndex(info.FullMethod, "/")
	method := info.FullMethod[pos+1:]
	token := req.(CommonRequest).GetConsumerToken()
	log.Info(method, "token", token, "req", req)
	resp, err = handler(ctx, req)
	log.Debug("Finish handling", "resp", resp, "err", err)
	return
}

func (rpc *IndexerRPCServices) Start() error {
	defer rpc.wg.Done()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(rpc.interceptor))
	defer grpcServer.GracefulStop()
	protobuf.RegisterIndexerRpcServiceServer(grpcServer, rpc)
	listen, err := net.Listen("tcp", ":"+rpc.Cfg.RpcPort)
	if err != nil {
		log.Error("net listen failed", "err", err)
		return err
	}
	reflection.Register(grpcServer)
	log.Info("savour dao start success", "port", rpc.Cfg.RpcPort)
	if err := grpcServer.Serve(listen); err != nil {
		log.Error("grpc server serve failed", "err", err)
		return err
	}
	return nil
}

func (rpc *IndexerRPCServices) Stop() {
	rpc.cancel()
	rpc.wg.Wait()
}
