package types

type ShadowScore struct {
	BlueChip        string `json:"blue_chip"`
	Fluidity        string `json:"fluidity"`
	Reliability     string `json:"reliability"`
	CommunityActive string `json:"community_active"`
	Heat            string `json:"heat"`
	PotentialIncome string `json:"potential_income"`
}
