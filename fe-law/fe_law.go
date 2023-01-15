package fe_law

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/savour-labs/fieryeyes/fe-law/law"
	"github.com/savour-labs/fieryeyes/fe-law/metrics"
	"github.com/savour-labs/fieryeyes/fe-law/service"
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
		log.Info("Initializing fe law", "curret version", gitVersion)
		feLaw, err := NewLaw(cfg)
		if err != nil {
			log.Error("Unable to create fe law", "error", err)
			return err
		}

		log.Info("Starting feLaw")

		if err := feLaw.Start(); err != nil {
			return err
		}
		defer feLaw.Stop()

		log.Info("Indexer started")

		<-(chan struct{})(nil)

		return nil
	}
}

type Law struct {
	ctx           context.Context
	cfg           Config
	lawRpcService *services.LawRPCServices
	metrics       *metrics.Metrics
}

func NewLaw(cfg Config) (*Law, error) {
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

	m := metrics.NewMetrics(nil)

	if cfg.MetricsServerEnable {
		go func() {
			_, err := m.Serve(cfg.MetricsHostname, cfg.MetricsPort)
			if err != nil {
				log.Error("metrics server failed to start", "err", err)
			}
		}()
		log.Info("metrics server enabled", "host", cfg.MetricsHostname, "port", cfg.MetricsPort)
	}

	feLawCfg := &law.FeLawConfig{
		GiantWhaleWalletAddress: cfg.GwwAddressConfig,
		NftCollections:          cfg.NftCollectionsConfig,
		SingleNft:               cfg.SingleNftConfig,
	}
	flaw, err := law.NewFeLaw(ctx, feLawCfg)
	if err != nil {
		log.Error("new fe law fail", "err", err)
		return nil, err
	}

	iRpcConfig := &services.LawRPCConfig{
		RpcHost: cfg.RpcHost,
		RpcPort: strconv.FormatUint(cfg.RpcPort, 10),
		FeLaw:   flaw,
	}
	iRpcServices, err := services.NewLawRPCServices(ctx, iRpcConfig)
	if err != nil {
		log.Error("new law rpc services fail", "err", err)
		return nil, err
	}
	return &Law{
		ctx:           ctx,
		cfg:           cfg,
		lawRpcService: iRpcServices,
		metrics:       m,
	}, nil
}

func (l Law) Start() error {
	err := l.lawRpcService.Start()
	if err != nil {
		return err
	}
	return nil
}

func (l Law) Stop() {
	l.lawRpcService.Stop()
}
