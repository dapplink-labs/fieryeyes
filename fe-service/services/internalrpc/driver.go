package internalrpc

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/savour-labs/fieryeyes/fe-service/db"
	"github.com/savour-labs/fieryeyes/fe-service/protobuf"
	"github.com/savour-labs/fieryeyes/fe-service/services/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	"runtime/debug"
	"strings"
	"sync"
)

type IInternalRpcServices interface {
	GetAddressInfo(ctx context.Context, req *protobuf.AddressInfoReq) (*protobuf.AddressInfoRep, error)
	GetNftCollectionsInfo(ctx context.Context, req *protobuf.NftCollectionsInfoReq) (*protobuf.NftCollectionsInfoRep, error)
	GetNftInfo(ctx context.Context, req *protobuf.NftInfoReq) (*protobuf.NftInfoRep, error)
}

type CommonRequest interface {
	GetConsumerToken() string
}

type InternalRpcConfig struct {
	RpcHost  string
	RpcPort  string
	Database *db.Database
}

type InternalRpcServices struct {
	Ctx    context.Context
	Cfg    *InternalRpcConfig
	cancel func()
	wg     sync.WaitGroup
	IInternalRpcServices
}

func NewIndexerRPCServices(ctx context.Context, cfg *InternalRpcConfig) (*InternalRpcServices, error) {
	ctxt, cancel := context.WithTimeout(ctx, common.DefaultTimeout)
	defer cancel()
	return &InternalRpcServices{
		Ctx:    ctxt,
		Cfg:    cfg,
		cancel: cancel,
	}, nil
}

func (rpc *InternalRpcServices) GetAddressInfo(ctx context.Context, req *protobuf.AddressInfoReq) (*protobuf.AddressInfoRep, error) {
	return &protobuf.AddressInfoRep{
		Code: protobuf.ReturnCode_SUCCESS,
		Msg:  "request address success",
	}, nil
}

func (rpc *InternalRpcServices) GetNftCollectionsInfo(ctx context.Context, req *protobuf.NftCollectionsInfoReq) (*protobuf.NftCollectionsInfoRep, error) {
	return &protobuf.NftCollectionsInfoRep{
		Code: protobuf.ReturnCode_SUCCESS,
		Msg:  "request nft collection success",
	}, nil
}

func (rpc *InternalRpcServices) GetNftInfo(ctx context.Context, req *protobuf.NftInfoReq) (*protobuf.NftInfoRep, error) {
	return &protobuf.NftInfoRep{
		Code: protobuf.ReturnCode_SUCCESS,
		Msg:  "request nft success",
	}, nil
}

func (rpc *InternalRpcServices) interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
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

func (rpc *InternalRpcServices) Start() error {
	defer rpc.wg.Done()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(rpc.interceptor))
	defer grpcServer.GracefulStop()
	protobuf.RegisterInternalRpcServiceServer(grpcServer, rpc)
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

func (rpc *InternalRpcServices) Stop() {
	rpc.cancel()
	rpc.wg.Wait()
}
