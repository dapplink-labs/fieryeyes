package openapi

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/savour-labs/fieryeyes/fe-service/db"
	"github.com/savour-labs/fieryeyes/fe-service/services/common"
	"strconv"
	"sync"
)

type ApiConfig struct {
	ApiServicePort int
	Debug          bool
	Database       *db.Database
}

type ApiService struct {
	Ctx    context.Context
	Cfg    *ApiConfig
	Echo   *echo.Echo
	Cancel func()
	Wg     sync.WaitGroup
}

func NewApiService(ctx context.Context, cfg *ApiConfig) (*ApiService, error) {
	_, cancel := context.WithTimeout(ctx, common.DefaultTimeout)
	defer cancel()
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Debug = true
	e.Use(middleware.Recover())
	server := &ApiService{
		Ctx:    ctx,
		Cfg:    cfg,
		Echo:   e,
		Cancel: cancel,
	}
	server.routes()
	return server, nil
}

func (as *ApiService) routes() {
	as.Echo.GET("api/v1/GetMainToken", as.GetMainToken)
	as.Echo.GET("api/v1/GetIndex", as.GetIndex)
	as.Echo.POST("api/v1/GetHotCollectionList", as.GetHotCollectionList)
	as.Echo.POST("api/v1/GetHotCollectionDetail", as.GetHotCollectionDetail)
	as.Echo.POST("api/v1/GetLiveMintList", as.GetLiveMintList)
	as.Echo.POST("api/v1/GetNftByCollectionId", as.GetNftByCollectionId)
	as.Echo.POST("api/v1/GetNftById", as.GetNftById)

}

func (as *ApiService) Start() error {
	defer as.Wg.Done()
	err := as.Echo.Start(":" + strconv.Itoa(as.Cfg.ApiServicePort))
	if err != nil {
		log.Error("open api sever start fail")
		return err
	}
	log.Info("open api sever start success", "port", as.Cfg.ApiServicePort)
	return nil
}

func (as *ApiService) Stop() {
	as.Cancel()
	as.Wg.Wait()
}
