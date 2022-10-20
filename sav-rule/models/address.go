package models

import (
	"gorm.io/gorm"
)

type Address struct {
	*gorm.Model
	ChainId  int64   `gorm:"description:ChainId;comment:ChainId" json:"chain_id"`
	Address  string  `gorm:"type:text;description:Address;comment:Address" json:"address"`
	Holder   string  `gorm:"type:text;description:Holder;comment:Holder" json:"holder"`
	TokenNum int64   `orm:"description:TokenNum;" json:"token_num"`
	TokenUsd float64 `orm:"default(1);digits(22);decimals(8)" json:"token_usd"`
	NftNum   int64   `orm:"description:TokenNum;" json:"nft_num"`
	NftUsd   float64 `orm:"default(1);digits(22);decimals(8)" json:"nft_usd"`
}
