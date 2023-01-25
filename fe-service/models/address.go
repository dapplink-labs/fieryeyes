package models

import (
	"gorm.io/gorm"
)

type Addresses struct {
	Id           uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Address      string `gorm:"type:char(42)"`
	Label        string `gorm:"type:char(64)" json:"label"`
	IsGiantWhale uint8  `json:"is_giant_whale" gorm:"column:is_giant_whale;default: 0;"` // 0: normal; 1: giant whale wallet
	Balance      string `gorm:"type:varchar(256)"  json:"balance"`
	*gorm.Model
}
