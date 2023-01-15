package fe_law

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"
)

func Main(gitVersion string) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		cfg, err := NewConfig(ctx)
		if err != nil {
			return err
		}

		log.Info("Initializing fe law")

		//indexer, err := NewIndexer(cfg)
		//if err != nil {
		//	log.Error("Unable to create indexer", "error", err)
		//	return err
		//}
		//
		//log.Info("Starting indexer")
		//
		//if err := indexer.Start(); err != nil {
		//	return err
		//}
		//defer indexer.Stop()
		//
		//log.Info("Indexer started")
		//
		//<-(chan struct{})(nil)

		return nil
	}
}

type Law struct {
}

func NewLaw() {

}
