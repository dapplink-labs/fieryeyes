package types

type CollectionReq struct {
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	OrderBy  int8 `json:"order_by"` //0: shadow score, 1: 24h volume; 2:price; 3:txn
}

type Collection struct {
	Id           uint64 `json:"id"`
	Rank         uint64 `json:"rank"`
	Image        string `json:"image"`
	Name         string `json:"name"`
	Holder       uint64 `json:"holder"`
	WhaleHolder  uint64 `json:"whale_holder"`
	SuggestLevel int8   `json:"suggest_level"`
	Volume       uint64 `json:"volume"`
	FloorPrice   string `json:"floor_price"`
	BestOffer    string `json:"best_offer"`
	ShadowScore  string `json:"shadow_score"`
}

type CollectionDetailReq struct {
	CollectionId uint64 `json:"collection_id"`
	Page         int    `json:"page"`
	PageSize     int    `json:"page_size"`
}

type Trading struct {
	StatTime string `json:"stat_time"`
	Price    string `json:"price"`
}

type Volume struct {
	StatTime string `json:"stat_time"`
	Volume   uint64 `json:"volume"`
}

type List struct {
	StatTime string `json:"stat_time"`
	PriceDis string `json:"price_dis"`
}

type FloorPrice struct {
	StatTime   string `json:"stat_time"`
	FloorPrice string `json:"floor_price"`
	BestOffer  string `json:"best_offer"`
}

type CollectionDetail struct {
	Id             uint64        `json:"id"`
	Name           string        `json:"name"`
	Image          string        `json:"image"`
	Creator        string        `json:"creator"`
	CollectionAddr string        `json:"collection_addr"`
	Holder         uint64        `json:"holder"`
	Chain          string        `json:"chain"`
	Introduce      string        `json:"introduce"`
	ShadowScore    *ShadowScore  `json:"shadow_score"`
	TradingList    []Trading     `json:"trading_list"`
	VolumeList     []Volume      `json:"volume_list"`
	ListList       []List        `json:"list_list"`
	FloorPriceList []FloorPrice  `json:"floor_price_list"`
	WhaleHolder    []WhaleHolder `json:"whale_holder"`
}
