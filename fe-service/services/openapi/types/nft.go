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

type NftDetailReq struct {
	NftId    uint64 `json:"nft_id"`
	Type     uint8  `json:"type"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type NftTx struct {
	FromAddr  string `json:"from_addr"`
	ToAddr    string `json:"to_addr"`
	Type      uint8  `json:"type"`
	Price     string `json:"price"`
	TradeTime string `json:"trade_time"`
}

type NftDetail struct {
	Id          uint64  `json:"id"`
	Image       string  `json:"image"`
	Name        string  `json:"name"`
	Chain       string  `json:"chain"`
	Contract    string  `json:"contract"`
	Creator     string  `json:"creator"`
	TokenUrl    string  `json:"token_url"`
	TokeId      string  `json:"toke_id"`
	Describe    string  `json:"describe"`
	MintHash    string  `json:"mint_hash"`
	MintTime    string  `json:"mint_time"`
	Holder      uint64  `json:"holder"`
	WhaleHolder uint64  `json:"whale_holder"`
	Price       string  `json:"price"`
	UsdPrice    string  `json:"usd_price"`
	TotalTxn    uint64  `json:"total_txn"`
	NftTxn      []NftTx `json:"nft_txn"`
}
