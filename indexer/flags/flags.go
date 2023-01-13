package flags

import (
	"github.com/urfave/cli"
	"time"
)

const envVarPrefix = "INDEXER_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	BuildEnvFlag = cli.StringFlag{
		Name: "build-env",
		Usage: "Build environment for which the binary is produced, " +
			"e.g. production or development",
		Required: true,
		EnvVar:   prefixEnvVar("BUILD_ENV"),
	}
	EthRpcFlag = cli.StringFlag{
		Name:     "eth-rpc",
		Usage:    "Ethereum network name",
		Required: true,
		EnvVar:   prefixEnvVar("ETH_RPC"),
	}
	ChainIDFlag = cli.StringFlag{
		Name:     "chain-id",
		Usage:    "Ethereum chain ID",
		Required: true,
		EnvVar:   prefixEnvVar("CHAIN_ID"),
	}

	DBHostFlag = cli.StringFlag{
		Name:     "db-host",
		Usage:    "Hostname of the database connection",
		Required: true,
		EnvVar:   prefixEnvVar("DB_HOST"),
	}
	DBPortFlag = cli.Uint64Flag{
		Name:     "db-port",
		Usage:    "Port of the database connection",
		Required: true,
		EnvVar:   prefixEnvVar("DB_PORT"),
	}
	DBUserFlag = cli.StringFlag{
		Name:     "db-user",
		Usage:    "Username of the database connection",
		Required: true,
		EnvVar:   prefixEnvVar("DB_USER"),
	}
	DBPasswordFlag = cli.StringFlag{
		Name:     "db-password",
		Usage:    "Password of the database connection",
		Required: true,
		EnvVar:   prefixEnvVar("DB_PASSWORD"),
	}
	DBNameFlag = cli.StringFlag{
		Name:     "db-name",
		Usage:    "Database name of the database connection",
		Required: true,
		EnvVar:   prefixEnvVar("DB_NAME"),
	}

	/* Optional Flags */
	DisableIndexer = cli.BoolFlag{
		Name:     "disable-indexer",
		Usage:    "Whether or not to enable the indexer on this instance",
		Required: false,
		EnvVar:   prefixEnvVar("DISABLE_INDEXER"),
	}
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
	SentryEnableFlag = cli.BoolFlag{
		Name:   "sentry-enable",
		Usage:  "Whether or not to enable Sentry. If true, sentry-dsn must also be set",
		EnvVar: prefixEnvVar("SENTRY_ENABLE"),
	}
	SentryDsnFlag = cli.StringFlag{
		Name:   "sentry-dsn",
		Usage:  "Sentry data source name",
		EnvVar: prefixEnvVar("SENTRY_DSN"),
	}
	SentryTraceRateFlag = cli.DurationFlag{
		Name:   "sentry-trace-rate",
		Usage:  "Sentry trace rate",
		Value:  50 * time.Millisecond,
		EnvVar: prefixEnvVar("SENTRY_TRACE_RATE"),
	}
	ConfDepthFlag = cli.Uint64Flag{
		Name:   "conf-depth",
		Usage:  "The number of confirmations after which headers are considered confirmed",
		Value:  20,
		EnvVar: prefixEnvVar("CONF_DEPTH"),
	}
	MaxHeaderBatchSizeFlag = cli.Uint64Flag{
		Name:   "max-header-batch-size",
		Usage:  "The maximum number of headers to request as a batch",
		Value:  2000,
		EnvVar: prefixEnvVar("MAX_HEADER_BATCH_SIZE"),
	}
	RPCHostnameFlag = cli.StringFlag{
		Name:   "RPC-hostname",
		Usage:  "The hostname of the RPC server",
		Value:  "127.0.0.1",
		EnvVar: prefixEnvVar("RPC_HOSTNAME"),
	}
	RPCPortFlag = cli.Uint64Flag{
		Name:   "RPC-port",
		Usage:  "The port of the RPC server",
		Value:  8080,
		EnvVar: prefixEnvVar("RPC_PORT"),
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
	HTTP2DisableFlag = cli.BoolFlag{
		Name:   "http2-disable",
		Usage:  "Whether or not to disable HTTP/2 support.",
		EnvVar: prefixEnvVar("HTTP2_DISABLE"),
	}
)

var requiredFlags = []cli.Flag{
	BuildEnvFlag,
	ChainIDFlag,
	EthRpcFlag,
	DBHostFlag,
	DBPortFlag,
	DBUserFlag,
	DBPasswordFlag,
	DBNameFlag,
}

var optionalFlags = []cli.Flag{
	DisableIndexer,
	LogLevelFlag,
	LogTerminalFlag,
	SentryEnableFlag,
	SentryDsnFlag,
	SentryTraceRateFlag,
	ConfDepthFlag,
	MaxHeaderBatchSizeFlag,
	RPCHostnameFlag,
	RPCPortFlag,
	MetricsServerEnableFlag,
	MetricsHostnameFlag,
	MetricsPortFlag,
	HTTP2DisableFlag,
}

// Flags contains the list of configuration options available to the binary.
var Flags = append(requiredFlags, optionalFlags...)
