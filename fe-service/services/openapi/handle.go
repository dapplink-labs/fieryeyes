package openapi

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/labstack/echo/v4"
	"github.com/savour-labs/fieryeyes/fe-service/models"
	"github.com/savour-labs/fieryeyes/fe-service/services/common"
	"github.com/savour-labs/fieryeyes/fe-service/services/openapi/types"
	"net/http"
)

const (
	SelfServiceOK     = 2000
	SelfServiceError  = 4000
	SelfInvalidParams = 4001
)

func (as *ApiService) GetMainTokenPrice(c echo.Context) error {
	mainToken := models.MainToken{}
	mainTokenList, err := mainToken.GetMainTokenList(as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "main token is not config")
		return c.JSON(http.StatusOK, retValue)
	}
	var mainTokenPriceList []types.TokenPrice
	for _, dbMt := range mainTokenList {
		tp := models.TokenPrice{MainTokenId: dbMt.Id}
		mtp, _ := tp.GetTokenPriceByTokenId(as.Cfg.Database.Db)
		mainTokenPriceList = append(
			mainTokenPriceList,
			types.TokenPrice{
				MainTokenName: dbMt.Name,
				UsdPrice:      mtp.UsdPrice,
				CnyPrice:      mtp.CnyPrice,
				DateTime:      mtp.DateTime,
			},
		)
	}
	retValue := common.BaseResource(true, SelfServiceOK, mainTokenPriceList, "get main token price success")
	return c.JSON(http.StatusOK, retValue)
}

func (as *ApiService) GetAddressInfo(c echo.Context) error {
	var addrReq types.AddressReq
	if err := c.Bind(&addrReq); err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "invalid request params")
		return c.JSON(http.StatusOK, retValue)
	}
	address := &models.Addresses{
		Id: addrReq.AddressId,
	}
	dbAddress, err := address.GetAddressById(as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "no this address in system")
		return c.JSON(http.StatusOK, retValue)
	}
	addrDaily := &models.DailyAddress{
		AddressId: addrReq.AddressId,
	}
	addressDailyList, _ := addrDaily.GetDailyAddressListById(addrReq.DailyPage, addrReq.DailyPageSize, as.Cfg.Database.Db)
	var addressDailyArray []types.AddressDaily
	for _, key := range addressDailyList {
		addressDailyArray = append(
			addressDailyArray,
			types.AddressDaily{
				AddressId:  key.AddressId,
				Balance:    key.Balance,
				TokenValue: key.TokenValue,
				NftValue:   key.NftValue,
				DateTime:   key.DateTime,
			},
		)
	}
	resultRet := &types.AddressInfoRep{
		Id:               dbAddress.Id,
		Address:          dbAddress.Address,
		Label:            dbAddress.Label,
		IsGiantWhale:     dbAddress.IsGiantWhale,
		Balance:          dbAddress.Balance,
		TokenValue:       dbAddress.TokenValue,
		NftValue:         dbAddress.NftValue,
		AddressDailyList: addressDailyArray,
	}
	retValue := common.BaseResource(true, SelfServiceOK, resultRet, "get address info success")
	return c.JSON(http.StatusOK, retValue)
}

func (as *ApiService) GetNftCollectionsInfo(c echo.Context) error {
	var collectionReq types.CollectionReq
	if err := c.Bind(&collectionReq); err != nil {
		retValue := common.BaseResource(false, SelfInvalidParams, nil, "params format error")
		return c.JSON(http.StatusOK, retValue)
	}
	collection := &models.Collection{
		Address: collectionReq.TokenAddress,
	}
	dbCollection, err := collection.GetCollectionById(as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "no this collection in system")
		return c.JSON(http.StatusOK, retValue)
	}
	collectionDaily := &models.CollectionDaily{
		CollectionId: dbCollection.Id,
	}
	cltDailyList, _ := collectionDaily.GetDailyCollectionListById(collectionReq.DailyPage, collectionReq.DailyPageSize, as.Cfg.Database.Db)
	var cltDailyArray []types.CollectionDailyList
	for _, key := range cltDailyList {
		cltDailyArray = append(
			cltDailyArray,
			types.CollectionDailyList{
				TotalHolder:             key.TotalHolder,
				AverageHolder:           key.AverageHolder,
				TotalGiantWhaleHolder:   key.TotalGiantWhaleHolder,
				AverageGiantWhaleHolder: key.AverageGiantWhaleHolder,
				TotalTxn:                key.TotalTxn,
				AverageTxn:              key.AverageTxn,
				AveragePrice:            key.AveragePrice,
				TotalPrice:              key.TotalPrice,
				DateTime:                key.DateTime,
			},
		)
	}
	resultRet := &types.CollectionInfo{
		Name:                    dbCollection.Name,
		Address:                 dbCollection.Address,
		Introduce:               dbCollection.Introduce,
		TotalHolder:             dbCollection.TotalHolder,
		AverageHolder:           dbCollection.AverageHolder,
		TotalGiantWhaleHolder:   dbCollection.TotalGiantWhaleHolder,
		AverageGiantWhaleHolder: dbCollection.AverageGiantWhaleHolder,
		TotalTxn:                dbCollection.TotalTxn,
		AverageTxn:              dbCollection.AverageTxn,
		AveragePrice:            dbCollection.AveragePrice,
		TotalPrice:              dbCollection.TotalPrice,
		SuggestLevel:            dbCollection.SuggestLevel,
		CollectionDaily:         cltDailyArray,
	}
	retValue := common.BaseResource(true, SelfServiceOK, resultRet, "get address info success")
	return c.JSON(http.StatusOK, retValue)
}

func (as *ApiService) GetNftInfo(c echo.Context) error {
	var nftReq types.NftReq
	if err := c.Bind(&nftReq); err != nil {
		retValue := common.BaseResource(false, SelfInvalidParams, nil, "params format error")
		return c.JSON(http.StatusOK, retValue)
	}
	nft := &models.Nft{
		TokenId: nftReq.TokenId,
	}
	dbNft, err := nft.GetNftById(as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "no this nft in system")
		return c.JSON(http.StatusOK, retValue)
	}
	nftInfo := &models.NftDaily{NftId: dbNft.Id}
	nftList, _ := nftInfo.GetDailyANftListById(nftReq.Page, nftReq.PageSize, as.Cfg.Database.Db)
	var nftArr []types.NftDailyStat
	for _, key := range nftList {
		nds := types.NftDailyStat{
			NftId:                 key.NftId,
			TotalTxn:              key.TotalTxn,
			TotalHolder:           key.TotalHolder,
			TotalGiantWhaleHolder: key.TotalGiantWhaleHolder,
			LatestPrice:           key.LatestPrice,
			DateTime:              key.DateTime,
		}
		nftArr = append(nftArr, nds)
	}
	nftAddress := &models.NftAddress{NftId: dbNft.Id}
	nftAddrList, _ := nftAddress.GetNftAddressListById(nftReq.Page, nftReq.PageSize, as.Cfg.Database.Db)
	var holderList []types.CurrentHolder
	var holderHistoryList []types.HistoricalHolderList
	for _, key := range nftAddrList {
		addr := models.Addresses{
			Id: key.AddressId,
		}
		dbAddress, _ := addr.GetAddressById(as.Cfg.Database.Db)
		log.Info("key.IsCurrent", "key.IsCurrent", key.IsCurrent)
		if key.IsCurrent != 0 {
			holderList = append(
				holderList,
				types.CurrentHolder{
					AddressId: dbAddress.Id,
					Address:   dbAddress.Address,
					Label:     dbAddress.Label,
				},
			)
		} else {
			holderHistoryList = append(
				holderHistoryList,
				types.HistoricalHolderList{
					AddressId: dbAddress.Id,
					Address:   dbAddress.Address,
					Label:     dbAddress.Label,
				},
			)
		}
	}
	nftInfos := &types.NftInfo{
		Id:                    dbNft.Id,
		Address:               dbNft.Address,
		TokenId:               dbNft.TokenId,
		TokenUrl:              dbNft.TokenUrl,
		TotalTxn:              dbNft.TotalTxn,
		TotalHolder:           dbNft.TotalHolder,
		TotalGiantWhaleHolder: dbNft.TotalGiantWhaleHolder,
		LatestPrice:           dbNft.LatestPrice,
		SuggestLevel:          dbNft.SuggestLevel,
		NftDaily:              nftArr,
		CurrentHolder:         holderList,
		HistoricalHolder:      holderHistoryList,
	}
	retValue := common.BaseResource(true, SelfServiceOK, nftInfos, "get nft info success")
	return c.JSON(http.StatusOK, retValue)
}
