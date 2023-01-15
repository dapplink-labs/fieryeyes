package law

import (
	"context"
)

type GiantWhaleWalletAddress struct {
	TokenValue    string
	NftValue      string
	TokenActivity uint64
	NftActivity   uint64
	TotalToken    uint64
	TotalNft      uint64
}

type NftCollections struct {
	AverageValue            string
	AverageTransactions     uint64
	DailyTransactions       uint64
	HolderAddress           uint64
	AverageTransactionPrice string
	DailyTransactionPrice   string
}

type SingleNft struct {
	TotalTransactions       uint64
	DailyTransactions       uint64
	LatestPrice             string
	AverageTransactionPrice string
	DailyTransactionPrice   string
}

type FeLawConfig struct {
	giantWhaleWalletAddress *GiantWhaleWalletAddress
	nftCollections          *NftCollections
	singleNft               *SingleNft
}

type FeLaw struct {
	Ctx      context.Context
	FlConfig *FeLawConfig
}

func NewFeLaw(ctx context.Context, Conf *FeLawConfig) (*FeLaw, error) {
	return &FeLaw{
		Ctx:      ctx,
		FlConfig: Conf,
	}, nil
}

func (fw *FeLaw) GiantWhaleWalletAddressLaw() (*GiantWhaleWalletAddress, error) {
	return fw.FlConfig.giantWhaleWalletAddress, nil
}

func (fw *FeLaw) NftCollectionsLaw() (*NftCollections, error) {
	return fw.FlConfig.nftCollections, nil
}

func (fw *FeLaw) SingleNftLaw() (*SingleNft, error) {
	return fw.FlConfig.singleNft, nil
}
