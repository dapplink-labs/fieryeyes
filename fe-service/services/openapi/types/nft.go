package types

type NftReq struct {
	TokenId  string `json:"token_id"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type NftDailyStat struct {
	NftId                 uint64 `json:"nft_id"`
	TotalTxn              uint64 `json:"total_txn"`
	TotalHolder           uint64 `json:"total_holder"`
	TotalGiantWhaleHolder uint64 `json:"total_giant_whale_holder"`
	LatestPrice           string `json:"latest_price"`
}

type CurrentHolder struct {
	Address string `json:"address"`
	Label   string `json:"label"`
}

type HistoricalHolderList struct {
	Address string `json:"address"`
	Label   string `json:"label"`
}

type NftInfo struct {
	Id                    uint64                 `json:"id"`
	Address               string                 `json:"address"`
	TokenId               string                 `json:"token_id"`
	TokenUrl              string                 `json:"token_url"`
	TotalTxn              uint64                 `json:"total_txn"`
	TotalHolder           uint64                 `json:"total_holder"`
	TotalGiantWhaleHolder uint64                 `json:"total_giant_whale_holder"`
	LatestPrice           string                 `json:"latest_price"`
	SuggestLevel          uint8                  `json:"suggest_level"`
	NftDaily              []NftDailyStat         `json:"nft_daily"`
	Holder                *CurrentHolder         `json:"holder"`
	HistoricalHolder      []HistoricalHolderList `json:"historical_holder"`
}
