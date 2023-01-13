package indexer

import (
	"github.com/savour-labs/fieryeyes/indexer/flags"
	"github.com/urfave/cli"
)

type Config struct {
	EthRpc              string
	DisableHTTP2        bool
	LogLevel            string
	LogTerminal         bool
	DbUsername          string
	DbPassword          string
	DbHost              string
	DbPort              int
	DbName              string
	RpcHost             string
	RpcPort             string
	MetricsServerEnable bool
	MetricsHostname     string
	MetricsPort         uint64
}

func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		EthRpc:       ctx.GlobalString(flags.EthRpcFlag.Name),
		DisableHTTP2: ctx.GlobalBool(flags.HTTP2DisableFlag.Name),
		LogLevel:     ctx.GlobalString(flags.LogLevelFlag.Name),
		LogTerminal:  ctx.GlobalBool(flags.LogTerminalFlag.Name),
	}
	return cfg, nil
}
