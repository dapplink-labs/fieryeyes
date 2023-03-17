package models

import (
	"gorm.io/gorm"
)

type NftTxn struct {
	Id          uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	NftId       uint64 `json:"nft_id"`
	FromAddress string `gorm:"type:char(42)" json:"from_address"`
	ToAddress   string `gorm:"type:char(42)" json:"to_address"`
	TxType      uint8  `json:"tx_type" gorm:"column:tx_type;default: 0;"` // 0: 销售; 1: 报价; 2:列表；3：转移；4:取消
	TradePrice  string `gorm:"type:varchar(256)"  json:"trade_price"`
	DateTime    string `json:"date_time"`
	*gorm.Model
}

func (nt *NftTxn) TableName() string {
	return "nft_txn"
}

func (nt *NftTxn) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&nt).Error; err != nil {
		return err
	}
	return nil
}

func (nt *NftTxn) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&nt).Error; err != nil {
		return err
	}
	return nil
}

func (nt *NftTxn) GetNftTxnList(page, pageSize int, db *gorm.DB) ([]NftTxn, error) {
	var nftTxn []NftTxn
	if err := db.Where("nft_id = ?", nt.NftId).Where("tx_type", nt.TxType).Offset((page - 1) * pageSize).Limit(pageSize).Find(&nftTxn).Error; err != nil {
		return nil, err
	}
	return nftTxn, nil
}
