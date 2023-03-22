package models

import (
	"gorm.io/gorm"
	"time"
)

type Label struct {
	Id          int64      `json:"id" gorm:"primary_key;type:int AUTO_INCREMENT"`
	ChainName   string     `json:"chain_name" gorm:"column:chain_name;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL;default: '';comment:公链名称;uniqueIndex:idx_account_addr_chain_name,priority:2"`
	AccountAddr string     `json:"account_addr" gorm:"column:account_addr;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL;default: '';comment:钱包地址;uniqueIndex:idx_account_addr_chain_name,priority:1"`
	Holder      string     `json:"holder" gorm:"column:holder;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL;default: '';comment:持有者;"`
	Tag         string     `json:"tag" gorm:"column:tag;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL;default: '';comment:标签;"`
	AddrType    string     `json:"addr_type" gorm:"column:addr_type;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL;default: ;comment:地址类型:account,contract;"`
	Amount      string     `json:"amount" gorm:"column:amount;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL;default: '';comment:持有代币的余额;"`
	TxCount     string     `json:"tx_count" gorm:"column:tx_count;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL;default: '';comment:交易数;"`
	CreatedAt   time.Time  `json:"created_at" gorm:"column:created_at;NOT NULL;default:CURRENT_TIMESTAMP;type:TIMESTAMP;index"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"column:updated_at;NOT NULL;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;type:TIMESTAMP"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"column:deleted_at;type:DATETIME"`
}

func (ll *Label) TableName() string {
	return "label"
}

func (ll *Label) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&ll).Error; err != nil {
		return err
	}
	return nil
}

func (ll *Label) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&ll).Error; err != nil {
		return err
	}
	return nil
}
