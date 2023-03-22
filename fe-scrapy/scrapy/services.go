package scrapy

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/savour-labs/fieryeyes/fe-scrapy/db"
	"github.com/savour-labs/fieryeyes/fe-scrapy/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	"runtime/debug"
	"strings"
)

type ScrapyRpcServices interface {
	GetSupportChain(ctx context.Context, req *protobuf.SupportChainReq) (*protobuf.SupportChainRep, error)
	SetGiantWhale(ctx context.Context, in *protobuf.SetGiantWhaleReq) (*protobuf.SetGiantWhaleRep, error)
	GetGiantWhale(ctx context.Context, in *protobuf.GetGiantWhaleReq) (*protobuf.GetGiantWhaleRep, error)
	RemoveGiantWhale(ctx context.Context, in *protobuf.RemoveGiantWhaleReq) (*protobuf.RemoveGiantWhaleRep, error)
}

type CommonRequest interface {
	GetChain() string
}

type RPCServices struct {
	Db      *db.Database
	RPCHost string
	RPCPort string
	ScrapyRpcServices
}

func NewRPCServices(db *db.Database, rpcHost, rpcPort string) (*RPCServices, error) {
	return &RPCServices{
		Db:      db,
		RPCHost: rpcHost,
		RPCPort: rpcPort,
	}, nil
}

func (rpc *RPCServices) GetSupportChain(ctx context.Context, req *protobuf.SupportChainReq) (*protobuf.SupportChainRep, error) {
	return nil, nil
}

func (rpc *RPCServices) SetGiantWhale(ctx context.Context, req *protobuf.SetGiantWhaleReq) (*protobuf.SetGiantWhaleRep, error) {
	return &protobuf.SetGiantWhaleRep{
		Code: protobuf.ReturnCode_SUCCESS,
		Msg:  "Success for api",
	}, nil
}

func (rpc *RPCServices) GetGiantWhale(ctx context.Context, req *protobuf.GetGiantWhaleReq) (*protobuf.GetGiantWhaleRep, error) {
	return nil, nil
}

func (rpc *RPCServices) RemoveGiantWhale(ctx context.Context, req *protobuf.RemoveGiantWhaleReq) (*protobuf.RemoveGiantWhaleRep, error) {
	return nil, nil
}

func (rpc *RPCServices) interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
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

func (rpc *RPCServices) Start() error {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(rpc.interceptor))
	defer grpcServer.GracefulStop()
	protobuf.RegisterGiantWhaleServiceServer(grpcServer, rpc)
	listen, err := net.Listen("tcp", ":"+rpc.RPCPort)
	if err != nil {
		log.Error("net listen failed", "err", err)
		return err
	}
	reflection.Register(grpcServer)
	log.Info("savour dao start success", "port", rpc.RPCPort)
	if err := grpcServer.Serve(listen); err != nil {
		log.Error("grpc server serve failed", "err", err)
		return err
	}
	return nil
}
