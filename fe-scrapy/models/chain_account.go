package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type ChainAccount struct {
	Id          int64           `json:"id" gorm:"primary_key;type:int AUTO_INCREMENT"`
	ChainName   string          `json:"chain_name" gorm:"column:chain_name;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL;default: ;comment:公链名称;uniqueIndex:idx_account_addr_chain_name,priority:2"`
	AccountAddr string          `json:"account_addr" gorm:"column:account_addr;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL;default: ;comment:钱包地址;uniqueIndex:idx_account_addr_chain_name,priority:1"`
	Holder      string          `json:"holder" gorm:"column:holder;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL;default: '';comment:持有者;"`
	Amount      decimal.Decimal `json:"amount" gorm:"column:amount;type:decimal(40,20)  NOT NULL;default:0.00000000000000000000 ;comment:持仓数量;"`
	TokenNum    decimal.Decimal `json:"token_num" gorm:"column:token_num;type:decimal(40,20)  NOT NULL;default:0.00000000000000000000 ;comment:代币的种类数;"`
	TokenUsd    decimal.Decimal `json:"token_usd" gorm:"column:token_usd;type:decimal(40,20)  NOT NULL;default:0.00000000000000000000 ;comment:代币总价值(usd);"`
	NftNum      decimal.Decimal `json:"nft_num" gorm:"column:nft_num;type:decimal(40,20)  NOT NULL;default:0.00000000000000000000 ;comment:nft种类;"`
	NftUsd      decimal.Decimal `json:"nft_usd" gorm:"column:nft_usd;type:decimal(40,20)  NOT NULL;default:0.00000000000000000000 ;comment:nft总价值(usd);"`
	CreatedAt   time.Time       `json:"created_at" gorm:"column:created_at;NOT NULL;default:CURRENT_TIMESTAMP;type:TIMESTAMP;index"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"column:updated_at;NOT NULL;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;type:TIMESTAMP"`
	DeletedAt   *time.Time      `json:"deleted_at" gorm:"column:deleted_at;type:DATETIME"`
}

func (ch *ChainAccount) TableName() string {
	return "chain_account"
}

func (ch *ChainAccount) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&ch).Error; err != nil {
		return err
	}
	return nil
}

func (ch *ChainAccount) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&ch).Error; err != nil {
		return err
	}
	return nil
}
