package models

import (
	"gorm.io/gorm"
)

type Token struct {
	Id                uint64 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Address           string `gorm:"type:char(42)" json:"address"`
	Name              string `gorm:"type:char(64)" json:"name"`
	Symbol            string `gorm:"type:char(64)" json:"symbol"`
	Type              string `gorm:"type:varchar(30);" json:"type"` // erc20:erc20; nft:nft(erc721/erc1155); other:other
	TotalSupply       string `gorm:"type:varchar(256)" json:"total_supply"`
	Collections       uint64 `json:"collections"`
	TotalNft          uint64 `json:"total_nft"`
	TotalTransactions uint64 `json:"total_transactions"`
	TotalHolders      uint64 `json:"total_holders"`
	*gorm.Model
}

func (t *Token) TableName() string {
	return "token"
}

func (t *Token) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&t).Error; err != nil {
		return err
	}
	return nil
}

func (t *Token) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&t).Error; err != nil {
		return err
	}
	return nil
}

func (t *Token) ExistToken(db *gorm.DB) bool {
	var count int64
	if err := db.Where("address = ?", t.Address).Count(&count).Error; err != nil {
		return false
	}
	if count > 0 {
		return true
	}
	return false
}
