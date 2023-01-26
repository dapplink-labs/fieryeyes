package types

type AddressReq struct {
	AddressId     uint64 `json:"address_id"`
	DailyPage     int    `json:"daily_page"`
	DailyPageSize int    `json:"daily_page_size"`
}

type AddressDaily struct {
	AddressId  uint64 `json:"address_id"`
	Balance    string `json:"balance"`
	TokenValue string `json:"token_value"`
	NftValue   string `json:"nft_value"`
	DateTime   string `json:"date_time"`
}

type AddressInfoRep struct {
	Id               uint64         `json:"id"`
	Address          string         `json:"address"`
	Label            string         `json:"label"`
	IsGiantWhale     uint8          `json:"is_giant_whale"`
	Balance          string         `json:"balance"`
	TokenValue       string         `json:"token_value"`
	NftValue         string         `json:"nft_value"`
	AddressDailyList []AddressDaily `json:"address_daily_list"`
}
