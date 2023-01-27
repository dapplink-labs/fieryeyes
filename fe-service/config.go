package fe_service

import (
	"github.com/savour-labs/fieryeyes/fe-service/flags"
	"github.com/urfave/cli"
	"time"
)

type Config struct {
	LoopInterval        time.Duration
	DbUsername          string
	DbPassword          string
	DbHost              string
	DbPort              uint64
	DbName              string
	ApiServicePort      uint64
	RpcHost             string
	RpcPort             uint64
	IndexerSocket       string
	LawSocket           string
	ScrapySocket        string
	EchoDebug           bool
	MetricsServerEnable bool
	MetricsHostname     string
	MetricsPort         uint64
	LogLevel            string
	LogTerminal         bool
}

func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		LoopInterval:        ctx.GlobalDuration(flags.LoopIntervalFlag.Name),
		DbUsername:          ctx.GlobalString(flags.DBUserNameFlag.Name),
		DbPassword:          ctx.GlobalString(flags.DBPasswordFlag.Name),
		DbHost:              ctx.GlobalString(flags.DBHostFlag.Name),
		DbPort:              ctx.GlobalUint64(flags.DBPortFlag.Name),
		DbName:              ctx.GlobalString(flags.DBNameFlag.Name),
		ApiServicePort:      ctx.GlobalUint64(flags.ApiServicePortFlag.Name),
		IndexerSocket:       ctx.GlobalString(flags.IndexerSocketFlag.Name),
		LawSocket:           ctx.GlobalString(flags.LawSocketFlag.Name),
		ScrapySocket:        ctx.GlobalString(flags.ScrapySockerFlag.Name),
		RpcHost:             ctx.GlobalString(flags.RPCHostNameFlag.Name),
		RpcPort:             ctx.GlobalUint64(flags.RPCPortFlag.Name),
		MetricsServerEnable: ctx.GlobalBool(flags.MetricsServerEnableFlag.Name),
		MetricsHostname:     ctx.GlobalString(flags.MetricsHostnameFlag.Name),
		MetricsPort:         ctx.GlobalUint64(flags.MetricsPortFlag.Name),
		LogLevel:            ctx.GlobalString(flags.LogLevelFlag.Name),
		LogTerminal:         ctx.GlobalBool(flags.LogTerminalFlag.Name),
		EchoDebug:           ctx.GlobalBool(flags.EchoDebugFlag.Name),
	}
	return cfg, nil
}
