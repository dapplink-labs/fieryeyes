package models

import (
	"gorm.io/gorm"
)

type NftDaily struct {
	Id                    uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	NftId                 uint64 `json:"nft_id"`
	TotalTxn              uint64 `json:"total_txn" gorm:"column:total_txn;default: 0;"`
	TotalHolder           uint64 `json:"total_holder" gorm:"column:total_holder;default: 0;"`
	TotalGiantWhaleHolder uint64 `json:"total_giant_whale_holder" gorm:"column:total_giant_whale_holder;default: 0;"`
	LatestPrice           string `gorm:"type:varchar(256)"  json:"latest_price"`
	*gorm.Model
}
