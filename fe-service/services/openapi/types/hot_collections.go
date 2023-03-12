package types

type HotCollection struct {
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
