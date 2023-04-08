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
			ApiUrl:    value.ApiUrl,
		}
		supportChainList = append(supportChainList, sChain)
	}
	// todo head stat
	headStat := &types.HeadDataStat{
		TotalNftValue:         "2020022",
		TotalNftValueRatio:    0.15,
		TotalNftValueStat:     []float64{2020022, 2010022, 2010022, 2030022, 2000022},
		TotalCollections:      "40120",
		TotalCollectionsRatio: 0.25,
		TotalCollectionsStat:  []float64{10001, 20001, 10001, 20001, 10000},
		TotalWhale:            "20001",
		TotalWhaleRatio:       0.35,
		TotalWhaleStat:        []float64{29001, 21001, 20801, 25001, 22001},
		TotalNft:              "1033332211",
		TotalNftRatio:         0.55,
		TotalNftStat:          []float64{1013332211, 1023332211, 1033332211, 1043332211, 1053332211},
	}
	hotCollections, err := collection.GetHotCollectionList(as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "get hot collection fail")
		return c.JSON(http.StatusOK, retValue)
	}
	var hotCollectionList []types.Collection
	for _, value := range hotCollections {
		hotC := types.Collection{
			Id:           value.Id,
			Rank:         value.Rank,
			Image:        value.CollectionImage,
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
			Rank:             value.Rank,
			Image:            value.CollectionImage,
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
	imgS := []string{"https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png", "https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png"}
	var WhaleHolderList []types.WhaleHolder
	HoldNftList := &types.HoldNft{
		TotalHold: 10,
		Images:    imgS,
	}
	HoldCollectionList := &types.HoldCollection{
		TotalHold: 10,
		Images:    imgS,
	}
	for _, value := range whaleHolders {
		wHolder := types.WhaleHolder{
			Address:            value.Address,
			TotalValue:         value.TokenValue + value.NftValue,
			HoldNftList:        HoldNftList,
			HoldCollectionList: HoldCollectionList,
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
	retValue := common.BaseResource(true, SelfServiceOK, index, "success")
	return c.JSON(http.StatusOK, retValue)
}

func (as *ApiService) GetHotCollectionList(c echo.Context) error {
	var collectionReq types.CollectionReq
	if err := c.Bind(&collectionReq); err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "invalid request params")
		return c.JSON(http.StatusOK, retValue)
	}
	collection := models.Collection{}
	collectionArray, err := collection.GetCollectionList(collectionReq.Page, collectionReq.PageSize, collectionReq.OrderBy, as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "get collection list fail")
		return c.JSON(http.StatusOK, retValue)
	}
	var collectionList []types.Collection
	for _, value := range collectionArray {
		hotC := types.Collection{
			Id:           value.Id,
			Rank:         value.Rank,
			Image:        value.CollectionImage,
			Name:         value.Name,
			Holder:       value.TotalHolder,
			WhaleHolder:  value.TotalGiantWhaleHolder,
			SuggestLevel: int8(value.SuggestLevel),
			Volume:       value.TotalTxn,
			FloorPrice:   value.FloorPrice,
			BestOffer:    value.BestOffer,
			ShadowScore:  "10",
		}
		collectionList = append(collectionList, hotC)
	}
	retValue := common.BaseResource(true, SelfServiceOK, collectionList, "success")
	return c.JSON(http.StatusOK, retValue)
}

func (as *ApiService) GetHotCollectionDetail(c echo.Context) error {
	collect := models.Collection{}
	collectStat := models.CollectionStat{}
	holder := models.Holders{}
	var cldIdReq types.CollectionDetailReq
	if err := c.Bind(&cldIdReq); err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "invalid request params")
		return c.JSON(http.StatusBadRequest, retValue)
	}
	collect.Id = cldIdReq.CollectionId
	collectStat.CollectionId = cldIdReq.CollectionId
	clDatail, err := collect.GetCollectionById(as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "get collection detail fail")
		return c.JSON(http.StatusBadRequest, retValue)
	}
	clStatList, err := collectStat.GetDailyCollectionListById(cldIdReq.Page, cldIdReq.PageSize, as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "get collection stat fail")
		return c.JSON(http.StatusBadRequest, retValue)
	}
	var tradeList []types.Trading
	var volumeList []types.Volume
	var listList []types.List
	var floorPriceList []types.FloorPrice
	for _, value := range clStatList {
		trade := types.Trading{
			StatTime: value.DateTime,
			Price:    value.TotalPrice,
		}
		tradeList = append(tradeList, trade)
		volume := types.Volume{
			StatTime: value.DateTime,
			Volume:   value.TotalTxn,
		}
		volumeList = append(volumeList, volume)
		list := types.List{
			StatTime: value.DateTime,
			PriceDis: value.AveragePrice,
		}
		listList = append(listList, list)
		floorPrice := types.FloorPrice{
			StatTime:   value.DateTime,
			FloorPrice: value.FloorPrice,
			BestOffer:  value.BestOffer,
		}
		floorPriceList = append(floorPriceList, floorPrice)
	}
	shadowScore := &types.ShadowScore{
		BlueChip:        "95",
		Fluidity:        "80",
		Reliability:     "60",
		CommunityActive: "70",
		Heat:            "50",
		PotentialIncome: "80",
	}
	whaleHolders, err := holder.GetWhaleHolderList(as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "get live mint fail")
		return c.JSON(http.StatusOK, retValue)
	}
	var WhaleHolderList []types.WhaleHolder
	imgS := []string{"https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png", "https://logo.nftscan.com/logo/0x34eebee6942d8def3c125458d1a86e0a897fd6f9.png"}
	HoldNftList := &types.HoldNft{
		TotalHold: 10,
		Images:    imgS,
	}
	HoldCollectionList := &types.HoldCollection{
		TotalHold: 10,
		Images:    imgS,
	}
	for _, value := range whaleHolders {
		wHolder := types.WhaleHolder{
			Address:            value.Address,
			TotalValue:         value.TokenValue + value.NftValue,
			HoldNftList:        HoldNftList,
			HoldCollectionList: HoldCollectionList,
			RealizePnl:         "10",
			Label:              value.Label,
		}
		WhaleHolderList = append(WhaleHolderList, wHolder)
	}
	collectDtl := types.CollectionDetail{
		Id:             clDatail.Id,
		Name:           clDatail.Name,
		Image:          clDatail.CollectionImage,
		Creator:        clDatail.Creator,
		CollectionAddr: clDatail.Address,
		Holder:         clDatail.TotalHolder,
		Chain:          "Ethereum",
		Introduce:      clDatail.Introduce,
		ShadowScore:    shadowScore,
		TradingList:    tradeList,
		VolumeList:     volumeList,
		ListList:       listList,
		FloorPriceList: floorPriceList,
		WhaleHolder:    WhaleHolderList,
	}
	retValue := common.BaseResource(true, SelfServiceOK, collectDtl, "success")
	return c.JSON(http.StatusOK, retValue)
}

func (as *ApiService) GetLiveMintList(c echo.Context) error {
	collect := models.Collection{}
	liveMints, err := collect.GetHotCollectionList(as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "get live mint fail")
		return c.JSON(http.StatusOK, retValue)
	}
	var liveMintArray []types.LiveMint
	for _, value := range liveMints {
		liveMint := types.LiveMint{
			Id:               value.Id,
			Rank:             value.Rank,
			Image:            value.CollectionImage,
			Name:             value.Name,
			Holder:           value.TotalHolder,
			WhaleHolder:      value.TotalGiantWhaleHolder,
			SuggestLevel:     int8(value.SuggestLevel),
			Mint:             value.TotalMint,
			MintPercent:      0.98,
			TotalMint:        value.TotalMint,
			TotalMintPercent: 0.98,
			LastMintTime:     value.LastMintTime,
		}
		liveMintArray = append(liveMintArray, liveMint)
	}
	retValue := common.BaseResource(true, SelfServiceOK, liveMintArray, "get live mint success")
	return c.JSON(http.StatusOK, retValue)
}

func (as *ApiService) GetNftByCollectionId(c echo.Context) error {
	nft := models.Nft{}
	var cNftReq types.CollectionNftReq
	if err := c.Bind(&cNftReq); err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "invalid request params")
		return c.JSON(http.StatusBadRequest, retValue)
	}
	nft.CollectionId = cNftReq.CollectId
	nfts, err := nft.GetNftListByCollectionId(cNftReq.Page, cNftReq.PageSize, as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "get nft list fail")
		return c.JSON(http.StatusOK, retValue)
	}
	var nftList []types.Nft
	for _, value := range nfts {
		nft := types.Nft{
			Id:        value.Id,
			Image:     value.Image,
			Name:      value.Name,
			Chain:     "Ethereum",
			Holder:    value.TotalHolder,
			HoldLabel: "cz",
			Price:     value.LatestPrice,
			UsdPrice:  value.PriceToUsd,
		}
		nftList = append(nftList, nft)
	}
	retValue := common.BaseResource(true, SelfServiceOK, nftList, "get nft list success")
	return c.JSON(http.StatusOK, retValue)
}

func (as *ApiService) GetNftById(c echo.Context) error {
	nft := models.Nft{}
	nftTx := models.NftTxn{}
	var nftReq types.NftDetailReq
	if err := c.Bind(&nftReq); err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "invalid request params")
		return c.JSON(http.StatusBadRequest, retValue)
	}
	nft.Id = nftReq.NftId
	nftTx.NftId = nftReq.NftId
	nftTx.TxType = nftReq.Type
	nftDtl, err := nft.GetNftById(as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "get nft detail fail")
		return c.JSON(http.StatusOK, retValue)
	}
	nftTxn, err := nftTx.GetNftTxnList(nftReq.Page, nftReq.PageSize, as.Cfg.Database.Db)
	if err != nil {
		retValue := common.BaseResource(true, SelfServiceError, nil, "get nft tx list fail")
		return c.JSON(http.StatusOK, retValue)
	}
	var nftTxnList []types.NftTx
	for _, value := range nftTxn {
		nftTx := types.NftTx{
			FromAddr:  value.FromAddress,
			ToAddr:    value.ToAddress,
			Type:      value.TxType,
			Price:     value.TradePrice,
			TradeTime: value.DateTime,
		}
		nftTxnList = append(nftTxnList, nftTx)
	}
	nftDetail := types.NftDetail{
		Id:          nftDtl.Id,
		Image:       nftDtl.Image,
		Name:        nftDtl.Name,
		Chain:       "Ethereum",
		Contract:    nftDtl.Address,
		Creator:     nftDtl.Creator,
		TokenUrl:    nftDtl.TokenUrl,
		TokeId:      nftDtl.TokenId,
		Describe:    nftDtl.Introduce,
		MintHash:    nftDtl.MintTxHash,
		MintTime:    nftDtl.MintTime,
		Holder:      nftDtl.TotalHolder,
		WhaleHolder: nftDtl.TotalGiantWhaleHolder,
		Price:       nftDtl.LatestPrice,
		UsdPrice:    nftDtl.PriceToUsd,
		TotalTxn:    nftDtl.TotalTxn,
		NftTxn:      nftTxnList,
	}
	retValue := common.BaseResource(true, SelfServiceOK, nftDetail, "get nft detail success")
	return c.JSON(http.StatusOK, retValue)
}
