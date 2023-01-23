package models

import (
	"gorm.io/gorm"
)

type Topics struct {
	Id      uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	EventId uint64 `json:"event_id"`
	Topic   string `gorm:"type:longtext" json:"topic"`
	*gorm.Model
}

func (tc *Topics) TableName() string {
	return "topics"
}

func (tc *Topics) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&tc).Error; err != nil {
		return err
	}
	return nil
}

func (tc *Topics) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&tc).Error; err != nil {
		return err
	}
	return nil
}
