package openapi

import (
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

func (as *ApiService) GetMainToken(c echo.Context) error {
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

func (as *ApiService) GetIndex(c echo.Context) error {
	chain := models.Chain{}
	collection := models.Collection{}
	holder := models.Holders{}
	chainList, err := chain.GetChainList(as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "get support chain list fail")
		return c.JSON(http.StatusOK, retValue)
	}
	var supportChainList []types.SupportChain
	for _, value := range chainList {
		sChain := types.SupportChain{
			ChainId:   value.Id,
			ChainName: value.Name,
			ChainIcon: value.Icon,
		}
		supportChainList = append(supportChainList, sChain)
	}
	// todo head stat
	headStat := &types.HeadDataStat{
		TotalNftValue:         "10000",
		TotalNftValueRatio:    0.95,
		TotalCollections:      "10000",
		TotalCollectionsRatio: 0.95,
		TotalWhale:            "10000",
		TotalWhaleRatio:       0.95,
		TotalNft:              "10000",
		TotalNftRatio:         0.95,
	}
	hotCollections, err := collection.GetHotCollectionList(as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "get hot collection fail")
		return c.JSON(http.StatusOK, retValue)
	}
	var hotCollectionList []types.HotCollection
	for _, value := range hotCollections {
		hotC := types.HotCollection{
			Id:           value.Id,
			Rank:         1,
			Image:        "",
			Name:         value.Name,
			Holder:       value.TotalHolder,
			WhaleHolder:  value.TotalGiantWhaleHolder,
			SuggestLevel: int8(value.SuggestLevel),
			Volume:       value.TotalTxn,
			FloorPrice:   value.FloorPrice,
			BestOffer:    value.BestOffer,
			ShadowScore:  "10",
		}
		hotCollectionList = append(hotCollectionList, hotC)
	}
	liveMints, err := collection.GetLiveMintList(as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "get live mint fail")
		return c.JSON(http.StatusOK, retValue)
	}
	var LiveMintList []types.LiveMint
	for _, value := range liveMints {
		lMint := types.LiveMint{
			Id:               value.Id,
			Rank:             1,
			Image:            "",
			Name:             value.Name,
			Holder:           value.TotalHolder,
			WhaleHolder:      value.TotalGiantWhaleHolder,
			SuggestLevel:     int8(value.SuggestLevel),
			Mint:             value.TotalMint,
			MintPercent:      0.95,
			TotalMint:        value.TotalMint,
			TotalMintPercent: 0.95,
			LastMintTime:     value.LastMintTime,
		}
		LiveMintList = append(LiveMintList, lMint)
	}
	whaleHolders, err := holder.GetWhaleHolderList(as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "get live mint fail")
		return c.JSON(http.StatusOK, retValue)
	}
	var WhaleHolderList []types.WhaleHolder
	for _, value := range whaleHolders {
		wHolder := types.WhaleHolder{
			Address:            value.Address,
			TotalValue:         value.TokenValue + value.NftValue,
			HoldNftList:        nil,
			HoldCollectionList: nil,
			RealizePnl:         "10",
			Label:              value.Label,
		}
		WhaleHolderList = append(WhaleHolderList, wHolder)
	}
	// todo: shadow score
	shadowScore := &types.ShadowScore{
		BlueChip:        "95",
		Fluidity:        "80",
		Reliability:     "60",
		CommunityActive: "70",
		Heat:            "50",
		PotentialIncome: "80",
	}
	index := &types.Index{
		SupportChains:     supportChainList,
		HeadStat:          headStat,
		HotCollectionList: hotCollectionList,
		LiveMintList:      LiveMintList,
		WhaleHolderList:   WhaleHolderList,
		ShadowScores:      shadowScore,
	}
	return c.JSON(http.StatusOK, index)
}

func (as *ApiService) GetHotCollectionList(c echo.Context) error {
	return c.JSON(http.StatusOK, "retValue")
}

func (as *ApiService) GetHotCollectionDetail(c echo.Context) error {
	return c.JSON(http.StatusOK, "retValue")
}

func (as *ApiService) GetLiveMintList(c echo.Context) error {
	return c.JSON(http.StatusOK, "retValue")
}
