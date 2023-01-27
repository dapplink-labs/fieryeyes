package models

import (
	"gorm.io/gorm"
)

type Collection struct {
	Id                      uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name                    string `gorm:"type:char(128)" json:"name"`
	Address                 string `gorm:"type:char(128)" json:"address"`
	Introduce               string `gorm:"type:text" json:"introduce"`
	TotalHolder             uint64 `json:"total_holder" gorm:"column:total_holder;default: 0;"`
	AverageHolder           uint64 `json:"average_holder" gorm:"column:total_holder;default: 0;"`
	TotalGiantWhaleHolder   uint64 `json:"total_giant_whale_holder" gorm:"column:total_giant_whale_holder;default: 0;"`
	AverageGiantWhaleHolder uint64 `json:"average_giant_whale_holder" gorm:"column:average_giant_whale_holder;default: 0;"`
	TotalTxn                uint64 `json:"total_txn" gorm:"column:total_txn;default: 0;"`
	AverageTxn              uint64 `json:"average_txn" gorm:"column:average_txn;default: 0;"`
	AveragePrice            string `gorm:"type:varchar(256)"  json:"average_price"`
	TotalPrice              string `gorm:"type:varchar(256)"  json:"total_price"`
	SuggestLevel            uint8  `json:"suggest_level" gorm:"column:suggest_level;default: 0;"`
	*gorm.Model
}

func (ct *Collection) TableName() string {
	return "collection"
}

func (ct *Collection) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&ct).Error; err != nil {
		return err
	}
	return nil
}

func (ct *Collection) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&ct).Error; err != nil {
		return err
	}
	return nil
}

func (ct *Collection) GetCollectionById(db *gorm.DB) (*Collection, error) {
	var collection *Collection
	if err := db.Where("address = ?", ct.Address).First(&collection).Error; err != nil {
		return nil, err
	}
	return collection, nil
}
