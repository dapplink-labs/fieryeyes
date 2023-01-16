package flags

import (
	"github.com/urfave/cli"
)

const envVarPrefix = "INDEXER_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	EthRpcFlag = cli.StringFlag{
		Name:     "eth-rpc",
		Usage:    "Ethereum network name",
		Required: true,
		EnvVar:   prefixEnvVar("ETH_RPC"),
	}
	SyncBlockHeightFlag = cli.Uint64Flag{
		Name:     "sync-block-height",
		Usage:    "sync current block",
		Required: true,
		EnvVar:   prefixEnvVar("SYNC_BLOCK_HEIGHT"),
	}
	LoopIntervalFlag = cli.DurationFlag{
		Name:     "loop-interval",
		Usage:    "loop interval for sync block",
		Required: true,
		EnvVar:   prefixEnvVar("LOOP_INTERVAL"),
	}
	DBUserNameFlag = cli.StringFlag{
		Name:     "db-user",
		Usage:    "Username of the database connection",
		Required: true,
		EnvVar:   prefixEnvVar("DB_USER_NAME"),
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
	HTTP2DisableFlag = cli.BoolFlag{
		Name:   "http2-disable",
		Usage:  "Whether or not to disable HTTP/2 support.",
		EnvVar: prefixEnvVar("HTTP2_DISABLE"),
	}
)

var requiredFlags = []cli.Flag{
	LoopIntervalFlag,
	DBUserNameFlag,
	DBHostFlag,
	DBPortFlag,
	DBPasswordFlag,
	DBNameFlag,
	RPCHostNameFlag,
	RPCPortFlag,
}

var optionalFlags = []cli.Flag{
	MetricsServerEnableFlag,
	MetricsHostnameFlag,
	MetricsPortFlag,
	LogLevelFlag,
	LogTerminalFlag,
	HTTP2DisableFlag,
}

// Flags contains the list of configuration options available to the binary.
var Flags = append(requiredFlags, optionalFlags...)
