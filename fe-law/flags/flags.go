package flags

import (
	"github.com/urfave/cli"
)

const envVarPrefix = "FE_LAW_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	RPCHostNameFlag = cli.StringFlag{
		Name:   "RPC-hostname",
		Usage:  "The hostname of the RPC server",
		Value:  "127.0.0.1",
		EnvVar: prefixEnvVar("RPC_HOST_NAME"),
	}
	RPCPortFlag = cli.Uint64Flag{
		Name:   "RPC-port",
		Usage:  "The port of the RPC server",
		Value:  8080,
		EnvVar: prefixEnvVar("RPC_PORT"),
	}

	TokenValueFlag = cli.StringFlag{
		Name:     "token-value",
		Usage:    "address hold total token value by usd",
		Required: true,
		EnvVar:   prefixEnvVar("TOKEN_VALUE"),
	}
	NftValueFlag = cli.StringFlag{
		Name:     "nft-value",
		Usage:    "address hold total nft value by usd",
		Required: true,
		EnvVar:   prefixEnvVar("NFT_VALUE"),
	}
	TokenActivityFlag = cli.Uint64Flag{
		Name:     "token-activity",
		Usage:    "address token activity",
		Required: true,
		EnvVar:   prefixEnvVar("TOKEN_ACTIVITY"),
	}
	NftActivityFlag = cli.Uint64Flag{
		Name:     "nft-activity",
		Usage:    "address nft activity",
		Required: true,
		EnvVar:   prefixEnvVar("NFT_ACTIVITY"),
	}
	TotalTokenFlag = cli.Uint64Flag{
		Name:     "total-token",
		Usage:    "address total token",
		Required: true,
		EnvVar:   prefixEnvVar("TOTAL_TOKEN"),
	}
	TotalNftFlag = cli.Uint64Flag{
		Name:     "total-nft",
		Usage:    "address total nft",
		Required: true,
		EnvVar:   prefixEnvVar("TOTAL_NFT"),
	}
	AverageValueFlag = cli.StringFlag{
		Name:     "collection-average-value",
		Usage:    "nft collection average value",
		Required: true,
		EnvVar:   prefixEnvVar("COLLECTION_AVERAGE_VALUE"),
	}
	AverageTransactionsFlag = cli.Uint64Flag{
		Name:     "collection-average-transactions",
		Usage:    "nft collection average transactions",
		Required: true,
		EnvVar:   prefixEnvVar("COLLECTION_AVERAGE_TRANSACTIONS"),
	}
	DailyTransactionsFlag = cli.Uint64Flag{
		Name:     "collection-daily-transactions",
		Usage:    "nft collection daily transactions",
		Required: true,
		EnvVar:   prefixEnvVar("COLLECTION_DAILY_TRANSACTIONS"),
	}
	HolderAddressFlag = cli.Uint64Flag{
		Name:     "collection_holder_address",
		Usage:    "nft collection holder addresses",
		Required: true,
		EnvVar:   prefixEnvVar("COLLECTION_HOLD_ADDRESS"),
	}
	AverageTransactionPriceFlag = cli.StringFlag{
		Name:     "collection-average-transactions-price",
		Usage:    "nft collection average transactions price",
		Required: true,
		EnvVar:   prefixEnvVar("COLLECTION_AVERAGE_TRANSACTIONS_PRICE"),
	}
	DailyTransactionPriceFlag = cli.StringFlag{
		Name:     "collection-daily-transactions-price",
		Usage:    "nft collection daily transactions price",
		Required: true,
		EnvVar:   prefixEnvVar("COLLECTION_DAILY_TRANSACTIONS_PRICE"),
	}
	STotalTransactionsFlag = cli.Uint64Flag{
		Name:     "single_nft_total_transactions",
		Usage:    "single nft total transactions",
		Required: true,
		EnvVar:   prefixEnvVar("NFT_TOTAL_TRANSACTIONS"),
	}
	SDailyTransactionsFlag = cli.Uint64Flag{
		Name:     "single_nft_daily_transactions",
		Usage:    "single nft daily transactions",
		Required: true,
		EnvVar:   prefixEnvVar("NFT_DAILY_TRANSACTIONS"),
	}
	SLatestPriceFlag = cli.StringFlag{
		Name:     "nft-latest-price",
		Usage:    "nft latest price",
		Required: true,
		EnvVar:   prefixEnvVar("NFT_LATEST_PRICE"),
	}
	SAverageTransactionPriceFlag = cli.StringFlag{
		Name:     "nft-average-transactions-price",
		Usage:    "nft average transactions price",
		Required: true,
		EnvVar:   prefixEnvVar("NFT_AVERAGE_TRANSACTIONS_PRICE"),
	}
	SDailyTransactionPriceFlag = cli.StringFlag{
		Name:     "nft-daily-transactions-price",
		Usage:    "nft daily transactions price",
		Required: true,
		EnvVar:   prefixEnvVar("NFT_DAILY_TRANSACTIONS_PRICE"),
	}

	/* Optional Flags */
	LogLevelFlag = cli.StringFlag{
		Name:   "log-level",
		Usage:  "The lowest log level that will be output",
		Value:  "info",
		EnvVar: prefixEnvVar("LOG_LEVEL"),
	}
	LogTerminalFlag = cli.BoolFlag{
		Name: "log-terminal",
		Usage: "If true, outputs logs in terminal format, otherwise prints " +
			"in JSON format. If SENTRY_ENABLE is set to true, this flag is " +
			"ignored and logs are printed using JSON",
		EnvVar: prefixEnvVar("LOG_TERMINAL"),
	}
	MetricsServerEnableFlag = cli.BoolFlag{
		Name:   "metrics-server-enable",
		Usage:  "Whether or not to run the embedded metrics server",
		EnvVar: prefixEnvVar("METRICS_SERVER_ENABLE"),
	}
	MetricsHostnameFlag = cli.StringFlag{
		Name:   "metrics-hostname",
		Usage:  "The hostname of the metrics server",
		Value:  "127.0.0.1",
		EnvVar: prefixEnvVar("METRICS_HOSTNAME"),
	}
	MetricsPortFlag = cli.Uint64Flag{
		Name:   "metrics-port",
		Usage:  "The port of the metrics server",
		Value:  7300,
		EnvVar: prefixEnvVar("METRICS_PORT"),
	}
)

var requiredFlags = []cli.Flag{
	TokenValueFlag,
	NftValueFlag,
	TokenActivityFlag,
	NftActivityFlag,
	TotalTokenFlag,
	TotalNftFlag,

	AverageValueFlag,
	AverageTransactionsFlag,
	DailyTransactionsFlag,
	HolderAddressFlag,
	AverageTransactionPriceFlag,
	DailyTransactionPriceFlag,

	STotalTransactionsFlag,
	SDailyTransactionsFlag,
	SLatestPriceFlag,
	SAverageTransactionPriceFlag,
	SDailyTransactionPriceFlag,

	RPCHostNameFlag,
	RPCPortFlag,
}

var optionalFlags = []cli.Flag{
	LogLevelFlag,
	LogTerminalFlag,
	MetricsServerEnableFlag,
	MetricsHostnameFlag,
	MetricsPortFlag,
}

var Flags = append(requiredFlags, optionalFlags...)
