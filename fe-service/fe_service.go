package fe_service

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/savour-labs/fieryeyes/fe-service/db"
	"github.com/savour-labs/fieryeyes/fe-service/services/indexer"
	"github.com/savour-labs/fieryeyes/fe-service/services/internalrpc"
	"github.com/savour-labs/fieryeyes/fe-service/services/openapi"
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

		log.Info("Initializing fe service")

		feService, err := NewFeService(cfg)
		if err != nil {
			log.Error("Unable to create indexer", "error", err)
			return err
		}

		log.Info("Starting fe service")

		if err := feService.Start(); err != nil {
			return err
		}
		defer feService.Stop()

		log.Info("fe service started")

		<-(chan struct{})(nil)

		return nil
	}
}

type FeService struct {
	ctx                 context.Context
	cfg                 Config
	apiService          *openapi.ApiService
	internalRpcServices *internalrpc.InternalRpcServices
	feServiceIndexer    *indexer.FeServiceIndexer
	db                  *db.Database
	// metrics             *metrics.Metrics
}

func NewFeService(cfg Config) (*FeService, error) {
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

	dbConfig := &db.DatabaseConfig{
		Username: cfg.DbUsername,
		Password: cfg.DbPassword,
		Host:     cfg.DbHost,
		Port:     cfg.DbPort,
		DbName:   cfg.DbName,
	}
	dbSelf, err := db.NewDatabase(ctx, dbConfig)
	if err != nil {
		log.Error("new database fail", "err", err)
		return nil, err
	}
	err = dbSelf.MigrateDb()
	if err != nil {
		log.Error("migrate db fail", "err", err)
		return nil, err
	}
	apiConfig := &openapi.ApiConfig{
		ApiServicePort: int(cfg.ApiServicePort),
		Debug:          cfg.EchoDebug,
		Database:       dbSelf,
	}
	apiService, err := openapi.NewApiService(ctx, apiConfig)
	if err != nil {
		log.Error("new evm chain client fail", "err", err)
		return nil, err
	}

	// m := metrics.NewMetrics()

	//if cfg.MetricsServerEnable {
	//	go func() {
	//		_, err := m.Serve(cfg.MetricsHostname, cfg.MetricsPort)
	//		if err != nil {
	//			log.Error("metrics server failed to start", "err", err)
	//		}
	//	}()
	//	log.Info("metrics server enabled", "host", cfg.MetricsHostname, "port", cfg.MetricsPort)
	//}

	iRpcConfig := &internalrpc.InternalRpcConfig{
		RpcHost:  cfg.RpcHost,
		RpcPort:  strconv.FormatUint(cfg.RpcPort, 10),
		Database: dbSelf,
	}
	iRpcServices, err := internalrpc.NewIndexerRPCServices(ctx, iRpcConfig)
	if err != nil {
		log.Error("new indexer rpc services fail", "err", err)
		return nil, err
	}

	feServiceConfig := &indexer.FeServiceIndexerConfig{
		IndexerRpcSocket: cfg.IndexerSocket,
		LoopInterval:     cfg.LoopInterval,
	}

	feService, err := indexer.NewFeServiceIndexer(ctx, feServiceConfig)
	if err != nil {
		log.Error("new fe services indexer fail", "err", err)
		return nil, err
	}
	return &FeService{
		ctx:                 ctx,
		cfg:                 cfg,
		apiService:          apiService,
		internalRpcServices: iRpcServices,
		feServiceIndexer:    feService,
		db:                  dbSelf,
		// metrics:             m,
	}, nil

}

func (fs *FeService) Start() error {
	go func() {
		err := fs.apiService.Start()
		if err != nil {
			log.Error("api service failed to start", "err", err)
		}
	}()
	go func() {
		err := fs.internalRpcServices.Start()
		if err != nil {
			log.Error("api service failed to start", "err", err)
		}
	}()
	err := fs.feServiceIndexer.Start()
	if err != nil {
		log.Error("fe service indexer failed to start", "err", err)
		return err
	}
	return nil
}

func (fs *FeService) Stop() {
	fs.apiService.Stop()
	fs.internalRpcServices.Stop()
	fs.feServiceIndexer.Stop()
}
