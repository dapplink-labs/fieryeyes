package models

import (
	"gorm.io/gorm"
)

type CollectionDaily struct {
	Id                      uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CollectionId            uint64 `json:"collect_id"`
	TotalHolder             uint64 `json:"total_holder" gorm:"column:total_holder;default: 0;"`
	AverageHolder           uint64 `json:"average_holder" gorm:"column:total_holder;default: 0;"`
	TotalGiantWhaleHolder   uint64 `json:"total_giant_whale_holder" gorm:"column:total_giant_whale_holder;default: 0;"`
	AverageGiantWhaleHolder uint64 `json:"average_giant_whale_holder" gorm:"column:average_giant_whale_holder;default: 0;"`
	TotalTxn                uint64 `json:"total_txn" gorm:"column:total_txn;default: 0;"`
	AverageTxn              uint64 `json:"average_txn" gorm:"column:average_txn;default: 0;"`
	AveragePrice            string `gorm:"type:varchar(256)"  json:"average_price"`
	TotalPrice              string `gorm:"type:varchar(256)"  json:"total_price"`
	SuggestLevel            uint8  `json:"suggest_level" gorm:"column:suggest_level;default: 0;"`
	DateTime                string `json:"date_time"`
	*gorm.Model
}

func (cd *CollectionDaily) TableName() string {
	return "collection_daily"
}

func (cd *CollectionDaily) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&cd).Error; err != nil {
		return err
	}
	return nil
}

func (cd *CollectionDaily) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&cd).Error; err != nil {
		return err
	}
	return nil
}

func (cd *CollectionDaily) GetDailyCollectionListById(page, pageSize int, db *gorm.DB) ([]CollectionDaily, error) {
	var dailyList []CollectionDaily
	if err := db.Where("collection_id = ?", cd.CollectionId).Offset((page - 1) * pageSize).Limit(pageSize).Find(&dailyList).Error; err != nil {
		return nil, err
	}
	return dailyList, nil
}
