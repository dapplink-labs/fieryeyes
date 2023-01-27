package indexer

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/savour-labs/fieryeyes/fe-service/services/common"
	"sync"
	"time"
)

type FeServiceIndexerConfig struct {
	IndexerRpcSocket string
	LoopInterval     time.Duration
}

type FeServiceIndexer struct {
	Ctx    context.Context
	Cfg    *FeServiceIndexerConfig
	Cancel func()
	Wg     sync.WaitGroup
}

func NewFeServiceIndexer(ctx context.Context, cfg *FeServiceIndexerConfig) (*FeServiceIndexer, error) {
	_, cancel := context.WithTimeout(ctx, common.DefaultTimeout)
	defer cancel()
	return &FeServiceIndexer{
		Ctx:    ctx,
		Cfg:    cfg,
		Cancel: cancel,
	}, nil
}

func (fsi *FeServiceIndexer) Start() error {
	fsi.Wg.Add(1)
	go fsi.EventLoop()
	return nil
}

func (fsi *FeServiceIndexer) Stop() {
	fsi.Cancel()
	fsi.Wg.Wait()
}

func (fsi *FeServiceIndexer) EventLoop() {
	defer fsi.Wg.Done()
	ticker := time.NewTicker(fsi.Cfg.LoopInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			log.Info("sync data from indexer and calc by laws")
		case err := <-fsi.Ctx.Done():
			log.Error("event loop exit, fail reason", "err", err)
			return
		}
	}
}
