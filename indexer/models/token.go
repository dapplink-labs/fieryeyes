package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Token struct {
	Id                uint64          `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Address           string          `gorm:"type:char(42)" json:"address"`
	Name              string          `gorm:"type:char(64)" json:"name"`
	Symbol            string          `gorm:"type:char(64)" json:"symbol"`
	Type              string          `gorm:"type:varchar(30);" json:"type"` // erc20:erc20; nft:nft(erc721/erc1155); other:other
	TotalSupply       decimal.Decimal `sql:"type:decimal(32,18);default:0;" json:"total_supply"`
	Collections       uint64          `json:"collections"`
	TotalNft          uint64          `json:"total_nft"`
	TotalTransactions uint64          `json:"total_transactions"`
	TotalHolders      uint64          `json:"total_holders"`
	*gorm.Model
}
