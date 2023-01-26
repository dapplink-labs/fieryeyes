package types

type CollectionReq struct {
	TokenAddress  string `json:"token_address"`
	DailyPage     int    `json:"daily_page"`
	DailyPageSize int    `json:"daily_page_size"`
}

type CollectionDailyList struct {
	TotalHolder             uint64 `json:"total_holder"`
	AverageHolder           uint64 `json:"average_holder"`
	TotalGiantWhaleHolder   uint64 `json:"total_giant_whale_holder"`
	AverageGiantWhaleHolder uint64 `json:"average_giant_whale_holder"`
	TotalTxn                uint64 `json:"total_txn"`
	AverageTxn              uint64 `json:"average_txn"`
	AveragePrice            string `json:"average_price"`
	TotalPrice              string `json:"total_price"`
	DateTime                string `json:"date_time"`
}

type CollectionInfo struct {
	Name                    string                `json:"name"`
	Address                 string                `json:"address"`
	Introduce               string                `json:"introduce"`
	TotalHolder             uint64                `json:"total_holder"`
	AverageHolder           uint64                `json:"average_holder"`
	TotalGiantWhaleHolder   uint64                `json:"total_giant_whale_holder"`
	AverageGiantWhaleHolder uint64                `json:"average_giant_whale_holder"`
	TotalTxn                uint64                `json:"total_txn"`
	AverageTxn              uint64                `json:"average_txn"`
	AveragePrice            string                `json:"average_price"`
	TotalPrice              string                `json:"total_price"`
	SuggestLevel            uint8                 `json:"suggest_level"`
	CollectionDaily         []CollectionDailyList `json:"collection_daily"`
}
