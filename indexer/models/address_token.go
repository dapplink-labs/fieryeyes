package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type AddressToken struct {
	Id        uint64          `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	AddressId uint64          `json:"address_id"`
	TokenId   uint64          `json:"token_id"`
	HoldValue decimal.Decimal `sql:"type:decimal(32,18);default:0;" json:"hold_value"`
	HoldNum   uint64          `json:"hold_num"`
	*gorm.Model
}
