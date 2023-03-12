package models

import (
	"gorm.io/gorm"
)

type Chain struct {
	Id     uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name   string `gorm:"type:varchar(64)" json:"name"`
	Icon   string `gorm:"type:varchar(256)" json:"icon"`
	ApiUrl string `gorm:"type:varchar(256)" json:"api_url"`
	*gorm.Model
}

func (ch *Chain) TableName() string {
	return "chain"
}

func (ch *Chain) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&ch).Error; err != nil {
		return err
	}
	return nil
}

func (ch *Chain) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&ch).Error; err != nil {
		return err
	}
	return nil
}

func (ch *Chain) GetChainList(db *gorm.DB) ([]Chain, error) {
	var chainList []Chain
	if err := db.Find(&chainList).Error; err != nil {
		return nil, err
	}
	return chainList, nil
}
