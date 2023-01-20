package models

import (
	"gorm.io/gorm"
)

type AddressToken struct {
	Id        uint64 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	AddressId uint64 `json:"address_id"`
	TokenId   uint64 `json:"token_id"`
	HoldValue string `gorm:"type:varchar(256)" json:"hold_value"`
	HoldNum   uint64 `json:"hold_num"`
	*gorm.Model
}

func (at *AddressToken) TableName() string {
	return "address_token"
}

func (at *AddressToken) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&at).Error; err != nil {
		return err
	}
	return nil
}

func (at *AddressToken) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&at).Error; err != nil {
		return err
	}
	return nil
}
