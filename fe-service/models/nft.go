package models

import (
	"gorm.io/gorm"
)

type Nft struct {
	Id                    uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CollectionId          uint64 `json:"collect_id"`
	Name                  string `gorm:"type:varchar(256)" json:"name"`
	Image                 string `gorm:"type:varchar(256)" json:"image"`
	Introduce             string `gorm:"type:text" json:"introduce"`
	CurrentHolderId       uint64 `json:"current_holder_id"`
	Creator               string `gorm:"type:char(128)" json:"creator"`
	Address               string `gorm:"type:char(128)"`
	MintTxHash            string `gorm:"type:char(66)" json:"mint_tx_hash"`
	MintTime              string `gorm:"type:varchar(128)" json:"mint_time"`
	TokenId               string `gorm:"type:varchar(128)" json:"token_id"`
	TokenUrl              string `gorm:"type:varchar(128)" json:"token_url"`
	TotalTxn              uint64 `json:"total_txn" gorm:"column:total_txn;default: 0;"`
	TotalHolder           uint64 `json:"total_holder" gorm:"column:total_holder;default: 0;"`
	TotalGiantWhaleHolder uint64 `json:"total_giant_whale_holder" gorm:"column:total_giant_whale_holder;default: 0;"`
	LatestPrice           string `gorm:"type:varchar(256)"  json:"latest_price"`
	SalesUnit             string `gorm:"type:varchar(256)"  json:"sales_unit"`
	PriceToUsd            string `gorm:"type:varchar(256)"  json:"price_to_usd"`
	SuggestLevel          uint8  `json:"suggest_level" gorm:"column:suggest_level;default: 0;"`
	*gorm.Model
}

func (nft *Nft) TableName() string {
	return "nft"
}

func (nft *Nft) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&nft).Error; err != nil {
		return err
	}
	return nil
}

func (nft *Nft) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&nft).Error; err != nil {
		return err
	}
	return nil
}

func (nft *Nft) GetNftById(db *gorm.DB) (*Nft, error) {
	var newNft *Nft
	if err := db.Where("id = ?", nft.Id).First(&newNft).Error; err != nil {
		return nil, err
	}
	return newNft, nil
}

func (nft *Nft) GetNftListByCollectionId(page, pageSize int, db *gorm.DB) ([]Nft, error) {
	var nftList []Nft
	if err := db.Where("collection_id = ?", nft.CollectionId).Offset((page - 1) * pageSize).Limit(pageSize).Find(&nftList).Error; err != nil {
		return nil, err
	}
	return nftList, nil
}
