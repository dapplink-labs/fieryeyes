package models

import (
	"gorm.io/gorm"
)

type NftStat struct {
	Id                    uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	NftId                 uint64 `json:"nft_id"`
	TotalTxn              uint64 `json:"total_txn" gorm:"column:total_txn;default: 0;"`
	TotalHolder           uint64 `json:"total_holder" gorm:"column:total_holder;default: 0;"`
	TotalGiantWhaleHolder uint64 `json:"total_giant_whale_holder" gorm:"column:total_giant_whale_holder;default: 0;"`
	LatestPrice           string `gorm:"type:varchar(256)"  json:"latest_price"`
	DateTime              string `json:"date_time"`
	*gorm.Model
}

func (ns *NftStat) TableName() string {
	return "nft_stat"
}

func (ns *NftStat) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&ns).Error; err != nil {
		return err
	}
	return nil
}

func (ns *NftStat) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&ns).Error; err != nil {
		return err
	}
	return nil
}

func (ns *NftStat) GetDailyANftListById(page, pageSize int, db *gorm.DB) ([]NftStat, error) {
	var nftList []NftStat

	if err := db.Where("nft_id = ?", ns.NftId).Offset((page - 1) * pageSize).Limit(pageSize).Find(&nftList).Error; err != nil {
		return nil, err
	}
	return nftList, nil
}
