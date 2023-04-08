package types

type SupportChain struct {
	ChainId   uint64 `json:"chain_id"`
	ChainName string `json:"chain_name"`
	ChainIcon string `json:"chain_icon"`
	ApiUrl    string `json:"api_url"`
}

type HeadDataStat struct {
	TotalNftValue         string    `json:"total_nft_value"`
	TotalNftValueRatio    float64   `json:"total_nft_value_ratio"`
	TotalNftValueStat     []float64 `json:"total_nft_value_stat"`
	TotalCollections      string    `json:"total_collections"`
	TotalCollectionsRatio float64   `json:"total_collections_ratio"`
	TotalCollectionsStat  []float64 `json:"total_collections_stat"`
	TotalWhale            string    `json:"total_whale"`
	TotalWhaleRatio       float64   `json:"total_whale_ratio"`
	TotalWhaleStat        []float64 `json:"total_whale_stat"`
	TotalNft              string    `json:"total_nft"`
	TotalNftRatio         float64   `json:"total_nft_ratio"`
	TotalNftStat          []float64 `json:"total_nft_stat"`
}

type Index struct {
	SupportChains     []SupportChain `json:"support_chain"`
	HeadStat          *HeadDataStat  `json:"head_data"`
	HotCollectionList []Collection   `json:"hot_collection_list"`
	LiveMintList      []LiveMint     `json:"live_mint_list"`
	WhaleHolderList   []WhaleHolder  `json:"whale_holder_list"`
	ShadowScores      *ShadowScore   `json:"shadow_score"`
}
