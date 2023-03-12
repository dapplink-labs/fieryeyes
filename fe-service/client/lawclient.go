package client

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/savour-labs/fieryeyes/fe-law/protobuf"
	"github.com/savour-labs/fieryeyes/fe-service/services/common"
	"google.golang.org/grpc"
	"time"
)

type LawClientConfig struct {
	LawSocket string `json:"law_socket"`
}

type LawClient struct {
	Ctx    context.Context
	Cfg    *LawClientConfig
	Client protobuf.LawRpcServiceClient
	Cancel func()
}

func NewLawClient(cfg *LawClientConfig) (*LawClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(common.DefaultTimeout))
	defer cancel()
	conn, err := grpc.Dial(cfg.LawSocket, grpc.WithInsecure())
	if err != nil {
		log.Error("Cannot connect to fe law", "LawSocket", cfg.LawSocket)
		return nil, err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	return &LawClient{
		Ctx:    ctx,
		Cfg:    cfg,
		Client: protobuf.NewLawRpcServiceClient(conn),
		Cancel: cancel,
	}, nil
}

func (lc *LawClient) GetGiantWhaleWalletAddressLaw() (*protobuf.GiantWhaleWalletAddressLawRep, error) {
	gwwReq := &protobuf.GiantWhaleWalletAddressLawReq{}
	gwLaw, err := lc.Client.GetGiantWhaleWalletAddressLaw(lc.Ctx, gwwReq)
	if err != nil {
		log.Error("get giant whale wallet address law fail")
		return nil, err
	}
	return gwLaw, nil
}

func (lc *LawClient) GetNftCollectionsLaw() (*protobuf.NftCollectionsLawRep, error) {
	nclReq := &protobuf.NftCollectionsLawReq{}
	nftCollectionLaw, err := lc.Client.GetNftCollectionsLaw(lc.Ctx, nclReq)
	if err != nil {
		log.Error("get nft collection law fail")
		return nil, err
	}
	return nftCollectionLaw, nil
}

func (lc *LawClient) GetSingleNftLaw() (*protobuf.SingleNftLawRep, error) {
	singleReq := &protobuf.SingleNftLawReq{}
	singleNftLaw, err := lc.Client.GetSingleNftLaw(lc.Ctx, singleReq)
	if err != nil {
		log.Error("get single nft law fail")
		return nil, err
	}
	return singleNftLaw, nil
}
