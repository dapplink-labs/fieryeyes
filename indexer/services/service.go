package services

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/savour-labs/fieryeyes/indexer/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime/debug"
	"strings"
)

type IIndexerRpcServices interface {
	//GetSupportChain(ctx context.Context, req *sav_scrapy.SupportChainReq) (*sav_scrapy.SupportChainRep, error)
}

type CommonRequest interface {
	GetChain() string
}

type IndexerRPCConfig struct {
	RpcHost string
	RpcPort string
}

type IndexerRPCServices struct {
	Ctx       context.Context
	Db        *db.Database
	RPCConfig *IndexerRPCConfig
	IIndexerRpcServices
}

func NewIndexerRPCServices(ctx context.Context, db *db.Database, cfg *IndexerRPCConfig) (*IndexerRPCServices, error) {
	return &IndexerRPCServices{
		Ctx:       ctx,
		Db:        db,
		RPCConfig: cfg,
	}, nil
}

//func (rpc *IndexerRPCServices) GetSupportChain(ctx context.Context, req *sav_scrapy.SupportChainReq) (*sav_scrapy.SupportChainRep, error) {
//	return nil, nil
//}

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
	chain := req.(CommonRequest).GetChain()
	log.Info(method, "chain", chain, "req", req)
	resp, err = handler(ctx, req)
	log.Debug("Finish handling", "resp", resp, "err", err)
	return
}

func (rpc *IndexerRPCServices) Start() error {
	//grpcServer := grpc.NewServer(grpc.UnaryInterceptor(rpc.interceptor))
	//defer grpcServer.GracefulStop()
	//sav_scrapy.RegisterGiantWhaleServiceServer(grpcServer, rpc)
	//listen, err := net.Listen("tcp", ":"+rpc.RPCPort)
	//if err != nil {
	//	log.Error("net listen failed", "err", err)
	//	return err
	//}
	//reflection.Register(grpcServer)
	//log.Info("savour dao start success", "port", rpc.RPCPort)
	//if err := grpcServer.Serve(listen); err != nil {
	//	log.Error("grpc server serve failed", "err", err)
	//	return err
	//}
	return nil
}
