package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	fe_scrapy "github.com/savour-labs/fieryeyes/fe-scrapy"
	"github.com/urfave/cli"
	"os"

	"github.com/savour-labs/fieryeyes/fe-scrapy/flags"
)

var (
	GitVersion = ""
	GitCommit  = ""
	GitDate    = ""
)

func main() {
	log.Root().SetHandler(
		log.LvlFilterHandler(
			log.LvlInfo,
			log.StreamHandler(os.Stdout, log.TerminalFormat(true)),
		),
	)

	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = fmt.Sprintf("%s-%s", GitVersion, params.VersionWithCommit(GitCommit, GitDate))
	app.Name = "scrapy"
	app.Usage = "Scrapy Service"
	app.Description = "Service for Scrapy address and nft data"
	app.Action = fe_scrapy.Main(GitVersion)
	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
