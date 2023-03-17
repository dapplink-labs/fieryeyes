package models

import (
	"gorm.io/gorm"
)

type Collection struct {
	Id                      uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Rank                    uint64 `json:"rank"`
	Name                    string `gorm:"type:char(128)" json:"name"`
	Creator                 string `gorm:"type:char(128)" json:"creator"`
	ChainName               string `gorm:"type:char(128)" json:"chain_name"`
	CollectionImage         string `gorm:"type:varchar(256)" json:"collection_image"`
	Address                 string `gorm:"type:char(128)" json:"address"`
	Introduce               string `gorm:"type:text" json:"introduce"`
	TotalHolder             uint64 `json:"total_holder" gorm:"column:total_holder;default: 0;"`
	AverageHolder           uint64 `json:"average_holder" gorm:"column:total_holder;default: 0;"`
	TotalGiantWhaleHolder   uint64 `json:"total_giant_whale_holder" gorm:"column:total_giant_whale_holder;default: 0;"`
	AverageGiantWhaleHolder uint64 `json:"average_giant_whale_holder" gorm:"column:average_giant_whale_holder;default: 0;"`
	TotalMint               uint64 `json:"total_mint" gorm:"column:total_mint;default: 0;" json:"total_mint"`
	TotalTxn                uint64 `json:"total_txn" gorm:"column:total_txn;default: 0;"`
	AverageTxn              uint64 `json:"average_txn" gorm:"column:average_txn;default: 0;"`
	FloorPrice              string `gorm:"type:varchar(256)"  json:"floor_price"`
	BestOffer               string `gorm:"type:varchar(256)"  json:"best_offer"`
	AveragePrice            string `gorm:"type:varchar(256)"  json:"average_price"`
	TotalPrice              string `gorm:"type:varchar(256)"  json:"total_price"`
	SuggestLevel            uint8  `json:"suggest_level" gorm:"column:suggest_level;default: 0;"`
	LastMintTime            string `json:"last_mint_time" gorm:"column:last_mint_time;default: 0;"`

	*gorm.Model
}

func (ct *Collection) TableName() string {
	return "collection"
}

func (ct *Collection) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&ct).Error; err != nil {
		return err
	}
	return nil
}

func (ct *Collection) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&ct).Error; err != nil {
		return err
	}
	return nil
}

func (ct *Collection) GetHotCollectionList(db *gorm.DB) ([]Collection, error) {
	var collectionList []Collection
	if err := db.Limit(6).Find(&collectionList).Error; err != nil {
		return nil, err
	}
	return collectionList, nil
}

func (ct *Collection) GetLiveMintList(db *gorm.DB) ([]Collection, error) {
	var collectionList []Collection
	if err := db.Order(ct.LastMintTime).Limit(6).Find(&collectionList).Error; err != nil {
		return nil, err
	}
	return collectionList, nil
}

func (ct *Collection) GetCollectionList(page, pageSize int, orderBy int8, db *gorm.DB) ([]Collection, error) {
	var collectionList []Collection
	if err := db.Order(ct.SuggestLevel).Offset((page - 1) * pageSize).Limit(pageSize).Find(&collectionList).Error; err != nil {
		return nil, err
	}
	return collectionList, nil
}

func (ct *Collection) GetCollectionById(db *gorm.DB) (*Collection, error) {
	var collection *Collection
	if err := db.Where("id = ?", ct.Id).First(&collection).Error; err != nil {
		return nil, err
	}
	return collection, nil
}
