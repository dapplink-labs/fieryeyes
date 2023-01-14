package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	Id          uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	BlockNumber uint64
	TxHash      string `gorm:"type:char(66)"`
	From        string `gorm:"type:char(42)"`
	to          string `gorm:"type:char(42)"`
	Contract    string `gorm:"type:char(42)"`
	Timestamp   time.Time
	InputData   []byte `gorm:"-"`
	*gorm.Model
}
