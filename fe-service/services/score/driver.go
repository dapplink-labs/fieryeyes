package score

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/savour-labs/fieryeyes/fe-service/services/common"
)

type FeServiceScoreConfig struct {
	LoopInterval time.Duration
}

type FeServiceScore struct {
	Ctx    context.Context
	Cfg    *FeServiceScoreConfig
	Cancel func()
	Wg     sync.WaitGroup
}

func NewFeServiceScore(ctx context.Context, cfg *FeServiceScoreConfig) (*FeServiceScore, error) {
	_, cancel := context.WithTimeout(ctx, common.DefaultTimeout)
	defer cancel()
	return &FeServiceScore{
		Ctx:    ctx,
		Cfg:    cfg,
		Cancel: cancel,
	}, nil
}

func (fsi *FeServiceScore) Start() error {
	fsi.Wg.Add(1)
	go fsi.EventLoop()
	return nil
}

func (fsi *FeServiceScore) Stop() {
	fsi.Cancel()
	fsi.Wg.Wait()
}

func (fsi *FeServiceScore) EventLoop() {
	defer fsi.Wg.Done()
	ticker := time.NewTicker(fsi.Cfg.LoopInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			log.Info("calc collection score")
			go fsi.CalcScores()
		case err := <-fsi.Ctx.Done():
			log.Error("event loop exit, fail reason", "err", err)
			return
		}
	}
}
