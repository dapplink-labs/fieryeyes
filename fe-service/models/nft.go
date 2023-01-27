package models

import (
	"gorm.io/gorm"
)

type Nft struct {
	Id                    uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Address               string `gorm:"type:char(128)"`
	TokenId               string `gorm:"type:varchar(128)" json:"token_id"`
	TokenUrl              string `gorm:"type:varchar(128)" json:"token_url"`
	TotalTxn              uint64 `json:"total_txn" gorm:"column:total_txn;default: 0;"`
	TotalHolder           uint64 `json:"total_holder" gorm:"column:total_holder;default: 0;"`
	TotalGiantWhaleHolder uint64 `json:"total_giant_whale_holder" gorm:"column:total_giant_whale_holder;default: 0;"`
	LatestPrice           string `gorm:"type:varchar(256)"  json:"latest_price"`
	SuggestLevel          uint8  `json:"suggest_level" gorm:"column:suggest_level;default: 0;"`
	*gorm.Model
}

func (nft *Nft) TableName() string {
	return "nft"
}

func (nft *Nft) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&nft).Error; err != nil {
		return err
	}
	return nil
}

func (nft *Nft) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&nft).Error; err != nil {
		return err
	}
	return nil
}

func (nft *Nft) GetNftById(db *gorm.DB) (*Nft, error) {
	var newNft *Nft
	if err := db.Where("id = ?", nft.Id).First(&newNft).Error; err != nil {
		return nil, err
	}
	return newNft, nil
}
