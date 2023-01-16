package openapi

import (
	"github.com/labstack/echo/v4"
	"github.com/savour-labs/fieryeyes/fe-service/services/common"
	"net/http"
)

const (
	SelfServiceOK     = 2000
	SelfServiceError  = 4000
	SelfInvalidParams = 4001
)

type NftCollectsRequest struct {
	TokenAddress string `json:"token_address"`
}

type NftRequest struct {
	TokenId int64 `json:"token_id"`
}

func (as *ApiService) GetAddressInfo(c echo.Context) error {
	retValue := common.BaseResource(true, SelfServiceOK, "address info", "get address info success")
	return c.JSON(http.StatusOK, retValue)
}

func (as *ApiService) GetNftCollectionsInfo(c echo.Context) error {
	var rsReq NftCollectsRequest
	if err := c.Bind(&rsReq); err != nil {
		retValue := common.BaseResource(false, SelfInvalidParams, nil, "params format error")
		return c.JSON(http.StatusOK, retValue)
	}
	retValue := common.BaseResource(true, SelfServiceOK, "Nft collections", "get nft collections success")
	return c.JSON(http.StatusOK, retValue)
}

func (as *ApiService) GetNftInfo(c echo.Context) error {
	var txReq NftRequest
	if err := c.Bind(&txReq); err != nil {
		retValue := common.BaseResource(false, SelfInvalidParams, nil, "Params format error")
		return c.JSON(http.StatusOK, retValue)
	}
	retValue := common.BaseResource(true, SelfServiceOK, "nft info", "get nft info success")
	return c.JSON(http.StatusOK, retValue)
}
