package whale

type AddressLabel struct {
	ChainName    string `json:"chain_name"`
	CoinName     string `json:"coin_name"`
	ContractAddr string `json:"contract_addr"`
	AddressType  string `json:"address_type"`
	AccountAddr  string `json:"account_addr"`
	Holder       string `json:"holder"`
	Price        string `json:"price"`
	Amount       string `json:"amount"`
	AmountUsd    string `json:"amount_usd"`
	TxCount      string `json:"tx_count"`
}
