package fe_law

import (
	"github.com/savour-labs/fieryeyes/fe-law/flags"
	"github.com/savour-labs/fieryeyes/fe-law/law"
	"github.com/urfave/cli"
)

type Config struct {
	GwwAddressConfig     *law.GiantWhaleWalletAddress
	NftCollectionsConfig *law.NftCollections
	SingleNftConfig      *law.SingleNft
	RpcHost              string
	RpcPort              uint64
	MetricsServerEnable  bool
	MetricsHostname      string
	MetricsPort          uint64
	LogLevel             string
	LogTerminal          bool
}

func NewConfig(ctx *cli.Context) (Config, error) {
	cfg := Config{
		GwwAddressConfig: &law.GiantWhaleWalletAddress{
			TokenValue:    ctx.GlobalString(flags.TokenValueFlag.Name),
			NftValue:      ctx.GlobalString(flags.NftValueFlag.Name),
			TokenActivity: ctx.GlobalUint64(flags.TokenActivityFlag.Name),
			NftActivity:   ctx.GlobalUint64(flags.NftActivityFlag.Name),
			TotalToken:    ctx.GlobalUint64(flags.TotalTokenFlag.Name),
			TotalNft:      ctx.GlobalUint64(flags.TotalNftFlag.Name),
		},
		NftCollectionsConfig: &law.NftCollections{
			AverageValue:            ctx.GlobalString(flags.AverageValueFlag.Name),
			AverageTransactions:     ctx.GlobalUint64(flags.AverageTransactionsFlag.Name),
			DailyTransactions:       ctx.GlobalUint64(flags.DailyTransactionsFlag.Name),
			HolderAddress:           ctx.GlobalUint64(flags.HolderAddressFlag.Name),
			AverageTransactionPrice: ctx.GlobalString(flags.AverageTransactionPriceFlag.Name),
			DailyTransactionPrice:   ctx.GlobalString(flags.DailyTransactionPriceFlag.Name),
		},
		SingleNftConfig: &law.SingleNft{
			TotalTransactions:       ctx.GlobalUint64(flags.STotalTransactionsFlag.Name),
			DailyTransactions:       ctx.GlobalUint64(flags.SDailyTransactionsFlag.Name),
			LatestPrice:             ctx.GlobalString(flags.SLatestPriceFlag.Name),
			AverageTransactionPrice: ctx.GlobalString(flags.SAverageTransactionPriceFlag.Name),
			DailyTransactionPrice:   ctx.GlobalString(flags.SDailyTransactionPriceFlag.Name),
		},
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
