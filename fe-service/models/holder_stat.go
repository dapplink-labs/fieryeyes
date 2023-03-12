package models

import (
	"gorm.io/gorm"
)

type HolderStat struct {
	Id         uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	AddressId  uint64 `json:"address_id"`
	DateTime   string `json:"date_time"`
	Balance    string `gorm:"type:varchar(256)"  json:"balance"`
	TokenValue string `gorm:"type:varchar(256)" json:"token_value"`
	NftValue   string `gorm:"type:varchar(256)" json:"nft_value"`
	*gorm.Model
}

func (hs *HolderStat) TableName() string {
	return "holder_stat"
}

func (hs *HolderStat) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&hs).Error; err != nil {
		return err
	}
	return nil
}

func (hs *HolderStat) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&hs).Error; err != nil {
		return err
	}
	return nil
}

func (hs *HolderStat) GetHolderStatListById(page, pageSize int, db *gorm.DB) ([]HolderStat, error) {
	var addrList []HolderStat
	if err := db.Where("address_id = ?", hs.AddressId).Offset((page - 1) * pageSize).Limit(pageSize).Find(&addrList).Error; err != nil {
		return nil, err
	}
	return addrList, nil
}
