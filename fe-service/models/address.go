package models

import (
	"gorm.io/gorm"
)

type Addresses struct {
	Id           uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Address      string `gorm:"type:char(42)"`
	Label        string `gorm:"type:char(64)" json:"label"`
	IsGiantWhale uint8  `json:"is_giant_whale" gorm:"column:is_giant_whale;default: 0;"` // 0: normal; 1: giant whale wallet
	Balance      string `gorm:"type:varchar(256)"  json:"balance"`
	TokenValue   string `gorm:"type:varchar(256)" json:"token_value"`
	NftValue     string `gorm:"type:varchar(256)" json:"nft_value"`
	*gorm.Model
}

func (ad *Addresses) TableName() string {
	return "addresses"
}

func (ad *Addresses) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&ad).Error; err != nil {
		return err
	}
	return nil
}

func (ad *Addresses) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&ad).Error; err != nil {
		return err
	}
	return nil
}

func (ad *Addresses) GetAddressById(db *gorm.DB) (*Addresses, error) {
	var addr *Addresses
	if err := db.Where("id = ?", ad.Id).First(&addr).Error; err != nil {
		return nil, err
	}
	return addr, nil
}
