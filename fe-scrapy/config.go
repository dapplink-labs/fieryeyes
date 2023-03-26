package fe_scrapy

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"

	"github.com/savour-labs/fieryeyes/fe-scrapy/flags"
)

type Config struct {
	DBHost     string
	DBPort     uint64
	DBUser     string
	DBPassword string
	DBName     string

	/* Optional Params */
	OkLink              string
	DisableSavScrapy    bool
	LogLevel            string
	LogTerminal         bool
	ConfDepth           uint64
	RPCHostname         string
	RPCPort             uint64
	DuneClientUrl       string
	DuneResultId        string
	DuneErrorId         string
	OkLinkClient        string
	OkAccessKey         string
	EtherClient         string
	MetricsServerEnable bool
	MetricsHostname     string
	MetricsPort         uint64
}

func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		DBHost:        ctx.GlobalString(flags.DBHostFlag.Name),
		DBPort:        ctx.GlobalUint64(flags.DBPortFlag.Name),
		DBUser:        ctx.GlobalString(flags.DBUserFlag.Name),
		DBPassword:    ctx.GlobalString(flags.DBPasswordFlag.Name),
		DBName:        ctx.GlobalString(flags.DBNameFlag.Name),
		DuneClientUrl: ctx.GlobalString(flags.DuneClientUrlFlag.Name),
		DuneResultId:  ctx.GlobalString(flags.DuneResultIdFlag.Name),
		DuneErrorId:   ctx.GlobalString(flags.DuneErrorIdFlag.Name),
		OkLinkClient:  ctx.GlobalString(flags.OkLinkClientFlag.Name),
		OkAccessKey:   ctx.GlobalString(flags.OkAccessKeyFlag.Name),
		EtherClient:   ctx.GlobalString(flags.EtherClientFlag.Name),
		/* Optional Flags */
		DisableSavScrapy:    ctx.GlobalBool(flags.DisableSavScrapyFlag.Name),
		LogLevel:            ctx.GlobalString(flags.LogLevelFlag.Name),
		LogTerminal:         ctx.GlobalBool(flags.LogTerminalFlag.Name),
		ConfDepth:           ctx.GlobalUint64(flags.ConfDepthFlag.Name),
		RPCHostname:         ctx.GlobalString(flags.RPCHostnameFlag.Name),
		RPCPort:             ctx.GlobalUint64(flags.RPCPortFlag.Name),
		MetricsServerEnable: ctx.GlobalBool(flags.MetricsServerEnableFlag.Name),
		MetricsHostname:     ctx.GlobalString(flags.MetricsHostnameFlag.Name),
		MetricsPort:         ctx.GlobalUint64(flags.MetricsPortFlag.Name),
	}
	err := ValidateConfig(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func ValidateConfig(cfg *Config) error {
	if cfg.LogLevel == "" {
		cfg.LogLevel = "debug"
	}
	_, err := log.LvlFromString(cfg.LogLevel)
	if err != nil {
		return err
	}
	return nil
}
