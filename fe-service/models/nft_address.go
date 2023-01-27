package models

import (
	"gorm.io/gorm"
)

type NftAddress struct {
	Id        uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	NftId     uint64 `json:"nft_id"`
	AddressId uint64 `json:"address_id"`
	IsCurrent uint8  `json:"is_current" gorm:"column:is_current;default: 0;"` // 0: default, 1: current holder
	*gorm.Model
}

func (na *NftAddress) TableName() string {
	return "nft_address"
}

func (na *NftAddress) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&na).Error; err != nil {
		return err
	}
	return nil
}

func (na *NftAddress) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&na).Error; err != nil {
		return err
	}
	return nil
}

func (na *NftAddress) GetNftAddressListById(page, pageSize int, db *gorm.DB) ([]NftAddress, error) {
	var nftAddresses []NftAddress
	if err := db.Where("nft_id = ?", na.NftId).Offset((page - 1) * pageSize).Limit(pageSize).Find(&nftAddresses).Error; err != nil {
		return nil, err
	}
	return nftAddresses, nil
}
