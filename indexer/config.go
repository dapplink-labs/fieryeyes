package indexer

import (
	"github.com/savour-labs/fieryeyes/indexer/flags"
	"github.com/urfave/cli"
	"time"
)

type Config struct {
	EthRpc              string
	DisableHTTP2        bool
	SyncBlockHeight     uint64
	LoopInterval        time.Duration
	DbUsername          string
	DbPassword          string
	DbHost              string
	DbPort              uint64
	DbName              string
	RpcHost             string
	RpcPort             uint64
	MetricsServerEnable bool
	MetricsHostname     string
	MetricsPort         uint64
	LogLevel            string
	LogTerminal         bool
}

func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		EthRpc:              ctx.GlobalString(flags.EthRpcFlag.Name),
		DisableHTTP2:        ctx.GlobalBool(flags.HTTP2DisableFlag.Name),
		SyncBlockHeight:     ctx.GlobalUint64(flags.SyncBlockHeightFlag.Name),
		LoopInterval:        ctx.GlobalDuration(flags.LoopIntervalFlag.Name),
		DbUsername:          ctx.GlobalString(flags.DBUserNameFlag.Name),
		DbPassword:          ctx.GlobalString(flags.DBPasswordFlag.Name),
		DbHost:              ctx.GlobalString(flags.DBHostFlag.Name),
		DbPort:              ctx.GlobalUint64(flags.DBPortFlag.Name),
		DbName:              ctx.GlobalString(flags.DBNameFlag.Name),
		RpcHost:             ctx.GlobalString(flags.RPCHostNameFlag.Name),
		RpcPort:             ctx.GlobalUint64(flags.RPCPortFlag.Name),
		MetricsServerEnable: ctx.GlobalBool(flags.MetricsServerEnableFlag.Name),
		MetricsHostname:     ctx.GlobalString(flags.MetricsHostnameFlag.Name),
		MetricsPort:         ctx.GlobalUint64(flags.MetricsPortFlag.Name),
		LogLevel:            ctx.GlobalString(flags.LogLevelFlag.Name),
		LogTerminal:         ctx.GlobalBool(flags.LogTerminalFlag.Name),
	}
	return cfg, nil
}
