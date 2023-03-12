package models

import (
	"gorm.io/gorm"
)

type NftHolder struct {
	Id        uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	NftId     uint64 `json:"nft_id"`
	AddressId uint64 `json:"address_id"`
	IsCurrent uint8  `json:"is_current" gorm:"column:is_current;default: 0;"` // 0: default, 1: current holder
	*gorm.Model
}

func (na *NftHolder) TableName() string {
	return "nft_holder"
}

func (na *NftHolder) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&na).Error; err != nil {
		return err
	}
	return nil
}

func (na *NftHolder) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&na).Error; err != nil {
		return err
	}
	return nil
}

func (na *NftHolder) GetNftHolderListById(page, pageSize int, db *gorm.DB) ([]NftHolder, error) {
	var NftHolderes []NftHolder
	if err := db.Where("nft_id = ?", na.NftId).Offset((page - 1) * pageSize).Limit(pageSize).Find(&NftHolderes).Error; err != nil {
		return nil, err
	}
	return NftHolderes, nil
}
