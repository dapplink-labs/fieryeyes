package types

type HoldNft struct {
	TotalHold uint64   `json:"total_hold"`
	Images    []string `json:"images"`
}

type HoldCollection struct {
	TotalHold uint64   `json:"total_hold"`
	Images    []string `json:"images"`
}

type WhaleHolder struct {
	Address            string          `json:"address"`
	TotalValue         string          `json:"total_value"`
	HoldNftList        *HoldNft        `json:"hold_nft_list"`
	HoldCollectionList *HoldCollection `json:"hold_collection_list"`
	RealizePnl         string          `json:"realize_pnl"`
	Label              string          `json:"label"`
}
