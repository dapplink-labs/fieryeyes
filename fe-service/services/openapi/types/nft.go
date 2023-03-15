package types

type CollectionNftReq struct {
	CollectId uint64 `json:"collect_id"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
}

type Nft struct {
	Id        uint64 `json:"id"`
	Image     string `json:"image"`
	Name      string `json:"name"`
	Chain     string `json:"chain"`
	Holder    uint64 `json:"holder"`
	HoldLabel string `json:"hold_label"`
	Price     string `json:"price"`
	UsdPrice  string `json:"usd_price"`
}
