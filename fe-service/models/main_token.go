package models

import (
	"gorm.io/gorm"
)

type MainToken struct {
	Id   uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"type:varchar(64)" json:"name"`
	*gorm.Model
}

func (mt *MainToken) TableName() string {
	return "main_token"
}

func (mt *MainToken) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&mt).Error; err != nil {
		return err
	}
	return nil
}

func (mt *MainToken) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&mt).Error; err != nil {
		return err
	}
	return nil
}

func (mt *MainToken) GetMainTokenList(db *gorm.DB) ([]MainToken, error) {
	var mainTokenList []MainToken
	if err := db.Find(&mainTokenList).Error; err != nil {
		return nil, err
	}
	return mainTokenList, nil
}
