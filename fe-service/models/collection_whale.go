package models

import (
	"gorm.io/gorm"
)

type CollectionWhale struct {
	Id           uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CollectionId uint64 `json:"collect_id"`
	AddressId    uint64 `json:"address_id"`
	*gorm.Model
}

func (cw *CollectionWhale) TableName() string {
	return "collection_whale"
}

func (cw *CollectionWhale) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&cw).Error; err != nil {
		return err
	}
	return nil
}

func (cw *CollectionWhale) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&cw).Error; err != nil {
		return err
	}
	return nil
}
