package score

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/savour-labs/fieryeyes/fe-service/db"
	"github.com/savour-labs/fieryeyes/fe-service/services/common"
)

type FeScoreServiceConfig struct {
	LoopInterval time.Duration
	Database     *db.Database
}

type FeScoreService struct {
	Ctx    context.Context
	Cfg    *FeScoreServiceConfig
	Cancel func()
	Wg     sync.WaitGroup
}

func NewFeScoreService(ctx context.Context, cfg *FeScoreServiceConfig) (*FeScoreService, error) {
	_, cancel := context.WithTimeout(ctx, common.DefaultTimeout)
	defer cancel()
	return &FeScoreService{
		Ctx:    ctx,
		Cfg:    cfg,
		Cancel: cancel,
	}, nil
}

func (fsi *FeScoreService) Start() error {
	fsi.Wg.Add(1)
	go fsi.EventLoop()
	return nil
}

func (fsi *FeScoreService) Stop() {
	fsi.Cancel()
	fsi.Wg.Wait()
}

func (fsi *FeScoreService) EventLoop() {
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
