package models

import (
	"gorm.io/gorm"
)

type Holders struct {
	Id           uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Address      string `gorm:"type:char(42)"`
	Label        string `gorm:"type:char(64)" json:"label"`
	IsGiantWhale uint8  `json:"is_giant_whale" gorm:"column:is_giant_whale;default: 0;"` // 0: normal; 1: giant whale wallet
	Balance      string `gorm:"type:varchar(256)"  json:"balance"`
	Owned        string `gorm:"type:varchar(256)"  json:"owned"`
	TokenValue   string `gorm:"type:varchar(256)" json:"token_value"`
	NftValue     string `gorm:"type:varchar(256)" json:"nft_value"`
	*gorm.Model
}

func (hl *Holders) TableName() string {
	return "holders"
}

func (hl *Holders) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&hl).Error; err != nil {
		return err
	}
	return nil
}

func (hl *Holders) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&hl).Error; err != nil {
		return err
	}
	return nil
}

func (hl *Holders) GetWhaleHolderList(db *gorm.DB) ([]Holders, error) {
	var addrList []Holders
	if err := db.Limit(6).Order(hl.NftValue).Find(&addrList).Error; err != nil {
		return nil, err
	}
	return addrList, nil
}

func (hl *Holders) GetAddressById(db *gorm.DB) (*Holders, error) {
	var addr *Holders
	if err := db.Where("id = ?", hl.Id).First(&addr).Error; err != nil {
		return nil, err
	}
	return addr, nil
}
