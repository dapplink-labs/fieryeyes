package flags

import (
	"github.com/urfave/cli"
)

const envVarPrefix = "FE_SERVICES_"

func prefixEnvVar(name string) string {
	return envVarPrefix + name
}

var (
	LoopIntervalFlag = cli.DurationFlag{
		Name:     "loop-interval",
		Usage:    "loop interval for sync data from indexer",
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
	ApiServicePortFlag = cli.Uint64Flag{
		Name:   "api-service-port",
		Usage:  "The port of the RPC server",
		Value:  8080,
		EnvVar: prefixEnvVar("API_SERVICE_PORT"),
	}
	IndexerSocketFlag = cli.StringFlag{
		Name:   "indexer-socket",
		Usage:  "The socket of indexer RPC server",
		Value:  "127.0.0.1:8080",
		EnvVar: prefixEnvVar("INDEXER_SOCKET"),
	}
	LawSocketFlag = cli.StringFlag{
		Name:   "law-socket",
		Usage:  "The socket of law RPC server",
		Value:  "127.0.0.1:8081",
		EnvVar: prefixEnvVar("LAW_SOCKET"),
	}
	ScrapySockerFlag = cli.StringFlag{
		Name:   "scrapy-socket",
		Usage:  "The socket of scrapy RPC server",
		Value:  "127.0.0.1:8082",
		EnvVar: prefixEnvVar("SCRAPY_SOCKET"),
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
	EchoDebugFlag = cli.BoolFlag{
		Name:   "echo-debug",
		Usage:  "Echo log debug",
		EnvVar: prefixEnvVar("ECHO_DEBUG"),
	}
)

var requiredFlags = []cli.Flag{
	LoopIntervalFlag,
	DBUserNameFlag,
	DBHostFlag,
	DBPortFlag,
	DBPasswordFlag,
	DBNameFlag,
	ApiServicePortFlag,
	IndexerSocketFlag,
	LawSocketFlag,
	ScrapySockerFlag,
	RPCHostNameFlag,
	RPCPortFlag,
}

var optionalFlags = []cli.Flag{
	LogLevelFlag,
	LogTerminalFlag,
	MetricsServerEnableFlag,
	MetricsHostnameFlag,
	MetricsPortFlag,
	EchoDebugFlag,
}

var Flags = append(requiredFlags, optionalFlags...)
