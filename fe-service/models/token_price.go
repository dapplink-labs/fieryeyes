package models

import (
	"gorm.io/gorm"
)

type TokenPrice struct {
	Id          uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	MainTokenId uint64 `json:"main_token_id"`
	UsdPrice    string `gorm:"type:varchar(64)" json:"usd_price"`
	CnyPrice    string `gorm:"type:varchar(64)" json:"cny_price"`
	DateTime    string `json:"date_time"`
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

func (tp *TokenPrice) GetTokenPriceByTokenId(db *gorm.DB) (*TokenPrice, error) {
	var mtp *TokenPrice
	if err := db.Where("main_token_id = ?", tp.MainTokenId).First(&mtp).Error; err != nil {
		return nil, err
	}
	return mtp, nil
}
