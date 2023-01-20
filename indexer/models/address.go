package models

import (
	"gorm.io/gorm"
)

type Addresses struct {
	Id      uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Address string `gorm:"type:char(42)" json:"address"`
	Label   string `gorm:"type:char(64)" json:"label"`
	Balance string `gorm:"type:varchar(256)"  json:"balance"`
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

func (ad *Addresses) ExistAddress(db *gorm.DB) bool {
	var count int64
	if err := db.Where("address = ?", ad.Address).Count(&count).Error; err != nil {
		return false
	}
	if count > 0 {
		return true
	}
	return false
}
