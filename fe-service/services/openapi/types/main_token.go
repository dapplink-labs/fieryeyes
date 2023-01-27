package types

type TokenPrice struct {
	MainTokenName string `json:"main_token_name"`
	UsdPrice      string `json:"usd_price"`
	CnyPrice      string `json:"cny_price"`
	DateTime      string `json:"date_time"`
}
