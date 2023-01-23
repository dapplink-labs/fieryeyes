package models

import (
	"gorm.io/gorm"
)

type Events struct {
	Id          uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Address     string `gorm:"type:char(42)" json:"address"`
	Data        string `gorm:"type:varchar(1024)" json:"data"`
	BlockNumber uint64 `json:"block_number"`
	TxHash      string `gorm:"type:char(66)" json:"tx_hash"`
	TxIndex     uint   `json:"tx_index"`
	BlockHash   string `json:"block_hash"`
	LogIndex    uint   `json:"log_index"`
	Removed     bool   `json:"removed"`
	*gorm.Model
}

func (e *Events) TableName() string {
	return "events"
}

func (e *Events) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&e).Error; err != nil {
		return err
	}
	return nil
}

func (e *Events) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&e).Error; err != nil {
		return err
	}
	return nil
}

func (e *Events) GetEventByTxHash(db *gorm.DB) (*Events, error) {
	var event Events
	if err := db.Where("tx_hash = ?", e.TxHash).First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}
