package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	Id          uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	BlockNumber uint64    `json:"block_number"`
	TxHash      string    `gorm:"type:char(66)" json:"tx_hash"`
	From        string    `gorm:"type:char(42)" json:"from"`
	To          string    `gorm:"type:char(42)" json:"to"`
	Value       string    `gorm:"type:varchar(256)" json:"value"`
	Contract    string    `gorm:"type:char(42)" json:"contract"`
	Status      uint64    `json:"status"`
	Timestamp   time.Time `json:"timestamp"`
	InputData   []byte    `gorm:"-" json:"input_data"`
	*gorm.Model
}

func (tx *Transaction) TableName() string {
	return "transaction"
}

func (tx *Transaction) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&tx).Error; err != nil {
		return err
	}
	return nil
}

func (tx *Transaction) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&tx).Error; err != nil {
		return err
	}
	return nil
}

func (tx *Transaction) GetFirstColumn(db *gorm.DB) (*Transaction, error) {
	var transaction *Transaction
	if err := db.First(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (tx *Transaction) GetTransactionByAddress(db *gorm.DB) (*Transaction, error) {
	var transaction *Transaction
	if err := db.Where("from = ? or to = ?", tx.From, tx.To).First(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}
