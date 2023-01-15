package services

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/savour-labs/fieryeyes/fe-law/law"
	"github.com/savour-labs/fieryeyes/fe-law/protobuf"
	chain "github.com/savour-labs/fieryeyes/indexer/blockchain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	"runtime/debug"
	"strings"
	"sync"
)

type ILawRpcServices interface {
	GetGiantWhaleWalletAddressLaw(ctx context.Context, req *protobuf.GiantWhaleWalletAddressLawReq) (*protobuf.GiantWhaleWalletAddressLawRep, error)
	GetNftCollectionsLaw(ctx context.Context, req *protobuf.NftCollectionsLawReq) (*protobuf.NftCollectionsLawRep, error)
	GetSingleNftLaw(ctx context.Context, req *protobuf.SingleNftLawReq) (*protobuf.SingleNftLawRep, error)
}

type CommonRequest interface {
	GetConsumerToken() string
}

type LawRPCConfig struct {
	RpcHost string
	RpcPort string
	FeLaw   *law.FeLaw
}

type LawRPCServices struct {
	Ctx    context.Context
	Cfg    *LawRPCConfig
	cancel func()
	wg     sync.WaitGroup
	ILawRpcServices
}

func NewLawRPCServices(ctx context.Context, cfg *LawRPCConfig) (*LawRPCServices, error) {
	ctxt, cancel := context.WithTimeout(ctx, chain.DefaultTimeout)
	defer cancel()
	return &LawRPCServices{
		Ctx:    ctxt,
		Cfg:    cfg,
		cancel: cancel,
	}, nil
}

func (rpc *LawRPCServices) GetGiantWhaleWalletAddressLaw(ctx context.Context, req *protobuf.GiantWhaleWalletAddressLawReq) (*protobuf.GiantWhaleWalletAddressLawRep, error) {
	gwwAddressLaw, _ := rpc.Cfg.FeLaw.GiantWhaleWalletAddressLaw()
	return &protobuf.GiantWhaleWalletAddressLawRep{
		Code:          protobuf.ReturnCode_SUCCESS,
		Msg:           "request success",
		TokenValue:    gwwAddressLaw.TokenValue,
		NftValue:      gwwAddressLaw.NftValue,
		TokenActivity: gwwAddressLaw.TokenActivity,
		NftActivity:   gwwAddressLaw.NftActivity,
		TotalToken:    gwwAddressLaw.TotalToken,
		TotalNft:      gwwAddressLaw.TotalNft,
	}, nil
}

func (rpc *LawRPCServices) GetNftCollectionsLaw(ctx context.Context, req *protobuf.NftCollectionsLawReq) (*protobuf.NftCollectionsLawRep, error) {
	collectionsLaw, _ := rpc.Cfg.FeLaw.NftCollectionsLaw()
	return &protobuf.NftCollectionsLawRep{
		Code:                              protobuf.ReturnCode_SUCCESS,
		Msg:                               "request success",
		CollectionAverageValue:            collectionsLaw.AverageValue,
		CollectionAverageTransactions:     collectionsLaw.AverageTransactions,
		CollectionDailyTransactions:       collectionsLaw.DailyTransactions,
		CollectionHolderAddress:           collectionsLaw.HolderAddress,
		CollectionAverageTransactionPrice: collectionsLaw.AverageTransactionPrice,
		CollectionDailyTransactionPrice:   collectionsLaw.DailyTransactionPrice,
	}, nil
}

func (rpc *LawRPCServices) GetSingleNftLaw(ctx context.Context, req *protobuf.SingleNftLawReq) (*protobuf.SingleNftLawRep, error) {
	nftLaw, _ := rpc.Cfg.FeLaw.SingleNftLaw()
	return &protobuf.SingleNftLawRep{
		Code:                       protobuf.ReturnCode_SUCCESS,
		Msg:                        "request success",
		NftTotalTransactions:       nftLaw.TotalTransactions,
		NftDailyTransactions:       nftLaw.DailyTransactions,
		NftLatestPrice:             nftLaw.LatestPrice,
		NftAverageTransactionPrice: nftLaw.AverageTransactionPrice,
		NftDailyTransactionPrice:   nftLaw.DailyTransactionPrice,
	}, nil
}

func (rpc *LawRPCServices) interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
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

func (rpc *LawRPCServices) Start() error {
	defer rpc.wg.Done()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(rpc.interceptor))
	defer grpcServer.GracefulStop()
	protobuf.RegisterLawRpcServiceServer(grpcServer, rpc)
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

func (rpc *LawRPCServices) Stop() {
	rpc.cancel()
	rpc.wg.Wait()
}
