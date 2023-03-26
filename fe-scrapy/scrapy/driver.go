package scrapy

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/savour-labs/fieryeyes/fe-scrapy/db"
	"github.com/savour-labs/fieryeyes/fe-scrapy/models"
	"github.com/savour-labs/fieryeyes/fe-scrapy/website/whale/dune"
	"github.com/savour-labs/fieryeyes/fe-scrapy/website/whale/etherscan"
	"github.com/savour-labs/fieryeyes/fe-scrapy/website/whale/oklink"
)

type DriverScapyConfig struct {
	DuneCfg     *dune.DuneClientConfig
	EtherConfig *etherscan.EtherClientConfig
	OkConfig    *oklink.OkClientConfig
	DataBase    *db.Database
}

type DriverScapy struct {
	cfg         *DriverScapyConfig
	duneClient  dune.IDuneClient
	etherClient *etherscan.EtherClient
	okClient    oklink.IOkClient
}

func NewDriverScapy(cfg *DriverScapyConfig) (*DriverScapy, error) {
	dClient := dune.NewDuneClient(cfg.DuneCfg)
	eClient := etherscan.NewEtherClient(cfg.EtherConfig)
	oClient := oklink.NewOkClient(cfg.OkConfig)
	return &DriverScapy{
		cfg:         cfg,
		duneClient:  dClient,
		etherClient: eClient,
		okClient:    oClient,
	}, nil
}

func (ds DriverScapy) dealDune() {
	dbLabel := models.Label{}
	whaleList, err := ds.duneClient.GetWhaleAddress()
	if err != nil {
		log.Error("dun client get whale address fail", "err", err)
	}
	var addrLabels []models.Label
	for _, whale := range whaleList {
		label := models.Label{
			ChainName:   whale.ChainName,
			AccountAddr: whale.AccountAddr,
			Holder:      whale.Holder,
			Tag:         whale.Holder,
			AddrType:    whale.AddressType,
			Amount:      whale.Amount,
			TxCount:     whale.TxCount,
		}
		addrLabels = append(addrLabels, label)
	}
	err = dbLabel.SelfInsert(ds.cfg.DataBase.Db)
	if err != nil {
		log.Error("insert db fail", "err", err)
	}
}

func (ds DriverScapy) dealOkLink() {
	dbLabel := models.Label{}
	whaleList, err := ds.okClient.GetWhaleAddress()
	if err != nil {
		log.Error("ok client get whale address fail", "err", err)
	}
	var addrLabels []models.Label
	for _, whale := range whaleList {
		label := models.Label{
			ChainName:   whale.ChainName,
			AccountAddr: whale.AccountAddr,
			Holder:      whale.Holder,
			Tag:         whale.Holder,
			AddrType:    whale.AddressType,
			Amount:      whale.Amount,
			TxCount:     whale.TxCount,
		}
		addrLabels = append(addrLabels, label)
	}
	err = dbLabel.SelfInsert(ds.cfg.DataBase.Db)
	if err != nil {
		log.Error("insert db fail", "err", err)
	}
}

func (ds DriverScapy) dealEtherScan() {
	dbLabel := models.Label{}
	element, err := ds.etherClient.CollEtherLabelAddress()
	if err != nil {
		log.Error("ether scan client get element fail", "err", err)
	}
	whaleList, err := ds.etherClient.DealData(element)
	if err != nil {
		log.Error("ether scan client get whale list fail", "err", err)
	}
	var addrLabels []models.Label
	for _, whale := range whaleList {
		label := models.Label{
			ChainName:   whale.ChainName,
			AccountAddr: whale.AccountAddr,
			Holder:      whale.Holder,
			Tag:         whale.Holder,
			AddrType:    whale.AddressType,
			Amount:      whale.Amount,
			TxCount:     whale.TxCount,
		}
		addrLabels = append(addrLabels, label)
	}
	err = dbLabel.SelfInsert(ds.cfg.DataBase.Db)
	if err != nil {
		log.Error("insert db fail", "err", err)
	}
}

func (ds DriverScapy) Run() {
	// ds.dealDune()
	// ds.dealOkLink()
	ds.dealEtherScan()
}
