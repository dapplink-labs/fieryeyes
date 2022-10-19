package indexer

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"
)

func Main(gitVersion string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		log.Info("Initializing indexer")
		return nil
	}
}
