package flags

import (
	"github.com/urfave/cli"
)

const envVarPrefix = "FE_SCRAPY_"

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
	DuneClientUrlFlag = cli.StringFlag{
		Name:     "dune-client-url",
		Usage:    "scrapy oklink website data",
		Required: true,
		EnvVar:   prefixEnvVar("DUNE_CLIENT_URL"),
	}
	DuneResultIdFlag = cli.StringFlag{
		Name:     "dune-result-id",
		Usage:    "scrapy oklink website data",
		Required: true,
		EnvVar:   prefixEnvVar("DUNE_RESULT_ID"),
	}
	DuneErrorIdFlag = cli.StringFlag{
		Name:     "dune-error-id",
		Usage:    "scrapy oklink website data",
		Required: true,
		EnvVar:   prefixEnvVar("DUNE_ERROR_ID"),
	}
	OkLinkClientFlag = cli.StringFlag{
		Name:     "ok_link-client",
		Usage:    "scrapy oklink website data",
		Required: true,
		EnvVar:   prefixEnvVar("OK_LINK_CLIENT"),
	}
	OkAccessKeyFlag = cli.StringFlag{
		Name:     "ok_access_key",
		Usage:    "ok access key",
		Required: true,
		EnvVar:   prefixEnvVar("OK_ACCESS_KEY"),
	}
	EtherClientFlag = cli.StringFlag{
		Name:     "ether-client",
		Usage:    "ether client website data",
		Required: true,
		EnvVar:   prefixEnvVar("ETHER_CLIENT"),
	}

	/* Optional Flags */
	DisableSavScrapyFlag = cli.BoolFlag{
		Name:     "disable-sav-scrapy",
		Usage:    "Whether or not to enable the sav-scrapy on this instance",
		Required: false,
		EnvVar:   prefixEnvVar("DISABLE_SAV_SCRAPY"),
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
	ConfDepthFlag = cli.Uint64Flag{
		Name:   "conf-depth",
		Usage:  "The number of confirmations after which headers are considered confirmed",
		Value:  20,
		EnvVar: prefixEnvVar("CONF_DEPTH"),
	}
	RPCHostnameFlag = cli.StringFlag{
		Name:   "rpc-hostname",
		Usage:  "The hostname of the RPC server",
		Value:  "127.0.0.1",
		EnvVar: prefixEnvVar("RPC_HOSTNAME"),
	}
	RPCPortFlag = cli.Uint64Flag{
		Name:   "rpc-port",
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
)

var requiredFlags = []cli.Flag{
	BuildEnvFlag,
	DBHostFlag,
	DBPortFlag,
	DBUserFlag,
	DBPasswordFlag,
	DBNameFlag,
	DuneClientUrlFlag,
	DuneResultIdFlag,
	DuneErrorIdFlag,
	OkLinkClientFlag,
	OkAccessKeyFlag,
	EtherClientFlag,
}

var optionalFlags = []cli.Flag{
	DisableSavScrapyFlag,
	LogLevelFlag,
	LogTerminalFlag,
	ConfDepthFlag,
	RPCHostnameFlag,
	RPCPortFlag,
	MetricsServerEnableFlag,
	MetricsHostnameFlag,
	MetricsPortFlag,
}

// Flags contains the list of configuration options available to the binary.
var Flags = append(requiredFlags, optionalFlags...)
