package types

type LiveMint struct {
	Id               uint64  `json:"id"`
	Rank             uint64  `json:"rank"`
	Image            string  `json:"image"`
	Name             string  `json:"name"`
	Holder           uint64  `json:"holder"`
	WhaleHolder      uint64  `json:"whale_holder"`
	SuggestLevel     int8    `json:"suggest_level"`
	Mint             uint64  `json:"volume"`
	MintPercent      float64 `json:"mint_percent"`
	TotalMint        uint64  `json:"total_mint"`
	TotalMintPercent float64 `json:"total_mint_percent"`
	LastMintTime     string  `json:"last_mint_time"`
}
