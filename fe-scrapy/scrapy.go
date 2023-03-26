package fe_scrapy

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/savour-labs/fieryeyes/fe-scrapy/db"
	"github.com/savour-labs/fieryeyes/fe-scrapy/scrapy"
	"github.com/savour-labs/fieryeyes/fe-scrapy/website/whale/dune"
	"github.com/savour-labs/fieryeyes/fe-scrapy/website/whale/etherscan"
	"github.com/savour-labs/fieryeyes/fe-scrapy/website/whale/oklink"
	"github.com/urfave/cli"
	"os"
	"strconv"
)

func Main(gitVersion string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		cfg, err := NewConfig(ctx)
		if err != nil {
			return err
		}
		log.Info("Initializing sav scrapy")
		savScrapy, err := NewSavourScrapy(cfg)
		if err != nil {
			log.Error("Unable to create sav scrapy", "error", err)
			return err
		}
		log.Info("Starting sav scrapy")
		if err := savScrapy.Start(); err != nil {
			return err
		}
		defer savScrapy.Stop()

		log.Info("sav scrapy started")
		<-(chan struct{})(nil)
		return nil
	}
}

type SavourScrapy struct {
	ctx         context.Context
	cfg         Config
	db          *db.Database
	rpcServer   *scrapy.RPCServices
	driverScapy *scrapy.DriverScapy
}

func NewSavourScrapy(cfg Config) (*SavourScrapy, error) {
	ctx := context.Background()
	var logHandler log.Handler
	if cfg.LogTerminal {
		logHandler = log.StreamHandler(os.Stdout, log.TerminalFormat(true))
	} else {
		logHandler = log.StreamHandler(os.Stdout, log.JSONFormat())
	}
	logLevel, err := log.LvlFromString(cfg.LogLevel)
	if err != nil {
		return nil, err
	}
	log.Root().SetHandler(log.LvlFilterHandler(logLevel, logHandler))
	dsnTemplate := "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local"
	dsn := fmt.Sprintf(dsnTemplate, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	dbd, err := db.NewDatabase(dsn)
	if err != nil {
		return nil, err
	}
	err = dbd.MigrateDb()
	if err != nil {
		log.Error("migrate db fail, err:", err)
		return nil, err
	}
	rpcService, err := scrapy.NewRPCServices(dbd, cfg.RPCHostname, strconv.FormatUint(cfg.RPCPort, 10))
	if err != nil {
		log.Error("new rpc server fail, err:", err)
		return nil, err
	}
	duneCfg := &dune.DuneClientConfig{
		ClientUrl: cfg.DuneClientUrl,
		ResultId:  cfg.DuneResultId,
		ErrorId:   cfg.DuneErrorId,
	}
	etherConfig := &etherscan.EtherClientConfig{
		ClientUrl: cfg.EtherClient,
	}
	okConfig := &oklink.OkClientConfig{
		ClientUrl:   cfg.EtherClient,
		OkAccessKey: cfg.OkAccessKey,
	}

	dConfig := &scrapy.DriverScapyConfig{
		DuneCfg:     duneCfg,
		EtherConfig: etherConfig,
		OkConfig:    okConfig,
		DataBase:    dbd,
	}

	ds, err := scrapy.NewDriverScapy(dConfig)
	if err != nil {
		log.Error("new driver scrapy server fail, err:", err)
		return nil, err
	}

	return &SavourScrapy{
		ctx:         ctx,
		cfg:         cfg,
		db:          dbd,
		rpcServer:   rpcService,
		driverScapy: ds,
	}, nil
}

func (ss *SavourScrapy) Start() error {
	if ss.cfg.DisableSavScrapy {
		log.Info("sav scrapy disabled, only serving data")
	}
	err := ss.db.MigrateDb()
	if err != nil {
		log.Error("migrate database fail, err:", err)
		return err
	}
	go ss.driverScapy.Run()
	return ss.rpcServer.Start()
}

func (ss *SavourScrapy) Stop() {
	log.Info("all fe scrapy stop")
}
