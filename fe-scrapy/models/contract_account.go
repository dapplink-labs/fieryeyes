package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type ContractAccount struct {
	Id           int64           `json:"id" gorm:"primary_key;type:int AUTO_INCREMENT"`
	ChainName    string          `json:"chain_name" gorm:"column:chain_name;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL;default: ;comment:公链名称;"`
	CoinName     string          `json:"coin_name" gorm:"column:coin_name;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL;default: ;comment:代币名称;"`
	ContractAddr string          `json:"contract_addr" gorm:"column:contract_addr;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL;default: ;comment:合约地址;uniqueIndex:idx_account_addr_contract_addr,priority:2"`
	AccountAddr  string          `json:"account_addr" gorm:"column:account_addr;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL;default: ;comment:钱包地址;uniqueIndex:idx_account_addr_contract_addr,priority:1"`
	Holder       string          `json:"holder" gorm:"column:holder;type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL;default: ;comment:持有者;"`
	Amount       decimal.Decimal `json:"amount" gorm:"column:amount;type:decimal(40,20)  NOT NULL;default:0.00000000000000000000 ;comment:持有代币的余额;"`
	Price        decimal.Decimal `json:"price" gorm:"column:price;type:decimal(40,20)  NOT NULL;default:0.00000000000000000000 ;comment:代币单价;"`
	UsdAmount    decimal.Decimal `json:"amount_usd" gorm:"column:amount_usd;type:decimal(40,20)  NOT NULL;default:0.00000000000000000000 ;comment:持有代币总价值(usd);"`
	CreatedAt    time.Time       `json:"created_at" gorm:"column:created_at;NOT NULL;default:CURRENT_TIMESTAMP;type:TIMESTAMP;index"`
	UpdatedAt    time.Time       `json:"updated_at" gorm:"column:updated_at;NOT NULL;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;type:TIMESTAMP"`
	DeletedAt    *time.Time      `json:"deleted_at" gorm:"column:deleted_at;type:DATETIME"`
}

func (ca *ContractAccount) TableName() string {
	return "contract_account"
}

func (ca *ContractAccount) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&ca).Error; err != nil {
		return err
	}
	return nil
}

func (ca *ContractAccount) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&ca).Error; err != nil {
		return err
	}
	return nil
}
