package models

import (
	"gorm.io/gorm"
)

type CollectionStat struct {
	Id                      uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CollectionId            uint64 `json:"collect_id"`
	TotalHolder             uint64 `json:"total_holder" gorm:"column:total_holder;default: 0;"`
	AverageHolder           uint64 `json:"average_holder" gorm:"column:total_holder;default: 0;"`
	TotalGiantWhaleHolder   uint64 `json:"total_giant_whale_holder" gorm:"column:total_giant_whale_holder;default: 0;"`
	AverageGiantWhaleHolder uint64 `json:"average_giant_whale_holder" gorm:"column:average_giant_whale_holder;default: 0;"`
	TotalMint               uint64 `json:"total_mint" gorm:"column:total_mint;default: 0;" json:"total_mint"`
	TotalTxn                uint64 `json:"total_txn" gorm:"column:total_txn;default: 0;"`
	AverageTxn              uint64 `json:"average_txn" gorm:"column:average_txn;default: 0;"`
	FloorPrice              string `gorm:"type:varchar(256)"  json:"floor_price"`
	BestOffer               string `gorm:"type:varchar(256)"  json:"best_offer"`
	AveragePrice            string `gorm:"type:varchar(256)"  json:"average_price"`
	TotalPrice              string `gorm:"type:varchar(256)"  json:"total_price"`
	SuggestLevel            uint8  `json:"suggest_level" gorm:"column:suggest_level;default: 0;"`
	DateTime                string `json:"date_time"`
	*gorm.Model
}

func (cs *CollectionStat) TableName() string {
	return "collection_stat"
}

func (cs *CollectionStat) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&cs).Error; err != nil {
		return err
	}
	return nil
}

func (cs *CollectionStat) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&cs).Error; err != nil {
		return err
	}
	return nil
}

func (cs *CollectionStat) GetDailyCollectionListById(page, pageSize int, db *gorm.DB) ([]CollectionStat, error) {
	var dailyList []CollectionStat
	if err := db.Where("collection_id = ?", cs.CollectionId).Offset((page - 1) * pageSize).Limit(pageSize).Find(&dailyList).Error; err != nil {
		return nil, err
	}
	return dailyList, nil
}
