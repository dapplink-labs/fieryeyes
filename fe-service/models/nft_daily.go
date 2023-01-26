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
	DateTime              string `json:"date_time"`
	*gorm.Model
}

func (nd *NftDaily) TableName() string {
	return "nft_daily"
}

func (nd *NftDaily) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&nd).Error; err != nil {
		return err
	}
	return nil
}

func (nd *NftDaily) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&nd).Error; err != nil {
		return err
	}
	return nil
}
