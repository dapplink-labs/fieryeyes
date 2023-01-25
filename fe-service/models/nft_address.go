package models

import (
	"gorm.io/gorm"
)

type NftAddress struct {
	Id        uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	NftId     uint64 `json:"nft_id"`
	AddressId uint64 `json:"address_id"`
	*gorm.Model
}
