package models

import (
	"gorm.io/gorm"
)

type TokenPrice struct {
	Id          uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	MainTokenId uint64 `json:"main_token_id"`
	UsdPrice    string `gorm:"type:varchar(64)" json:"usd_price"`
	CnyPrice    string `gorm:"type:varchar(64)" json:"cny_price"`
	StatType    uint8  `json:"stat_type" gorm:"column:stat_type;default: 0;"` // 0: normal; 1: daily
	*gorm.Model
}

func (tp *TokenPrice) TableName() string {
	return "token_price"
}

func (tp *TokenPrice) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&tp).Error; err != nil {
		return err
	}
	return nil
}

func (tp *TokenPrice) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&tp).Error; err != nil {
		return err
	}
	return nil
}
