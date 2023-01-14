package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Addresses struct {
	Id      uint64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Address string          `gorm:"type:char(42)"`
	Label   string          `gorm:"type:char(64)" json:"label"`
	Balance decimal.Decimal `sql:"type:decimal(32,18);default:0;"  json:"balance"`
	*gorm.Model
}
