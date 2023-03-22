package models

import (
	"gorm.io/gorm"
)

type Chain struct {
	Name string `gorm:"type:text;description:Name; comment:链名称"   json:"name"`
	*gorm.Model
}

func (ch *Chain) TableName() string {
	return "chain"
}

func (ch *Chain) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&ch).Error; err != nil {
		return err
	}
	return nil
}

func (ch *Chain) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&ch).Error; err != nil {
		return err
	}
	return nil
}
