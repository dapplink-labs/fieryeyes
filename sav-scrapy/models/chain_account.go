package models

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
	"github.com/weblazy/easy/utils/db/mysql"
	"gorm.io/gorm"
)

var ChainAccountHandler = &ChainAccount{}

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

func (t *ChainAccount) TableName() string {
	return "chain_account"
}

func (t *ChainAccount) Insert(db *gorm.DB, ctx context.Context, data *ChainAccount) error {
	if db == nil {
		db = GetDB(ctx)
	}
	return db.Create(data).Error
}
func (t *ChainAccount) BulkInsert(db *gorm.DB, ctx context.Context, fields []string, params []map[string]interface{}) error {
	if db == nil {
		db = GetDB(ctx)
	}
	return mysql.BulkInsert(db, t.TableName(), fields, params)
}
func (t *ChainAccount) BulkSave(db *gorm.DB, ctx context.Context, fields []string, params []map[string]interface{}) error {
	if db == nil {
		db = GetDB(ctx)
	}
	return mysql.BulkSave(db, t.TableName(), fields, params)
}
func (t *ChainAccount) Delete(db *gorm.DB, ctx context.Context, where string, args ...interface{}) error {
	if db == nil {
		db = GetDB(ctx)
	}
	return db.Where(where, args...).Delete(&ChainAccount{}).Error
}
func (t *ChainAccount) Updates(db *gorm.DB, ctx context.Context, data map[string]interface{}, where string, args ...interface{}) (int64, error) {
	if db == nil {
		db = GetDB(ctx)
	}
	db = db.Model(&ChainAccount{}).Where(where, args...).Updates(data)
	return db.RowsAffected, db.Error
}
func (t *ChainAccount) GetOne(ctx context.Context, where string, args ...interface{}) (*ChainAccount, error) {
	var obj ChainAccount
	return &obj, GetDB(ctx).Where(where, args...).Take(&obj).Error
}
func (*ChainAccount) GetList(ctx context.Context, where string, args ...interface{}) ([]*ChainAccount, error) {
	var list []*ChainAccount
	return list, GetDB(ctx).Where(where, args...).Find(&list).Error
}
func (t *ChainAccount) GetListWithLimit(ctx context.Context, limit int, where string, args ...interface{}) ([]*ChainAccount, error) {
	var list []*ChainAccount
	return list, GetDB(ctx).Where(where, args...).Limit(limit).Find(&list).Error
}
func (t *ChainAccount) GetListOrderLimit(ctx context.Context, order string, limit int, where string, args ...interface{}) ([]*ChainAccount, error) {
	var list []*ChainAccount
	if limit == 0 || limit > 10000 {
		limit = 10
	}
	return list, GetDB(ctx).Where(where, args...).Order(order).Limit(limit).Find(&list).Error
}
func (t *ChainAccount) GetListPage(ctx context.Context, pageNum, limit int, where string, args ...interface{}) ([]*ChainAccount, error) {
	var list []*ChainAccount
	offset := (pageNum - 1) * limit
	return list, GetDB(ctx).Where(where, args...).Offset(offset).Limit(limit).Find(&list).Error
}
func (t *ChainAccount) GetCount(ctx context.Context, where string, args ...interface{}) (int64, error) {
	var count int64
	return count, GetDB(ctx).Model(&ChainAccount{}).Where(where, args...).Count(&count).Error
}
func (t *ChainAccount) GetSumInt64(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	type sum struct {
		Num int64 `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB(ctx).Raw(sql, args...).Scan(&obj).Error
}
func (t *ChainAccount) GetSumFloat64(ctx context.Context, sql string, args ...interface{}) (float64, error) {
	type sum struct {
		Num float64 `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB(ctx).Raw(sql, args...).Scan(&obj).Error
}
func (t *ChainAccount) GetSumDecimal(ctx context.Context, sql string, args ...interface{}) (decimal.Decimal, error) {
	type sum struct {
		Num decimal.Decimal `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB(ctx).Raw(sql, args...).Scan(&obj).Error
}
