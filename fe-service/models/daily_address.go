package models

import (
	"gorm.io/gorm"
)

type DailyAddress struct {
	Id         uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	AddressId  uint64 `json:"address_id"`
	DateTime   string `json:"date_time"`
	Balance    string `gorm:"type:varchar(256)"  json:"balance"`
	TokenValue string `gorm:"type:varchar(256)" json:"token_value"`
	NftValue   string `gorm:"type:varchar(256)" json:"nft_value"`
	*gorm.Model
}

func (da *DailyAddress) TableName() string {
	return "daily_address"
}

func (da *DailyAddress) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&da).Error; err != nil {
		return err
	}
	return nil
}

func (da *DailyAddress) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&da).Error; err != nil {
		return err
	}
	return nil
}

func (da *DailyAddress) GetDailyAddressListById(page, pageSize int, db *gorm.DB) ([]DailyAddress, error) {
	var addrList []DailyAddress
	if err := db.Where("address_id = ?", da.AddressId).Offset((page - 1) * pageSize).Limit(pageSize).Find(&addrList).Error; err != nil {
		return nil, err
	}
	return addrList, nil
}
