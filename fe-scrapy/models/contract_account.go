package models

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
	"github.com/weblazy/easy/utils/db/mysql"
	"gorm.io/gorm"
)

var ContractAccountHandler = &ContractAccount{}

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

func (t *ContractAccount) TableName() string {
	return "contract_account"
}

func (t *ContractAccount) Insert(db *gorm.DB, ctx context.Context, data *ContractAccount) error {
	if db == nil {
		db = GetDB(ctx)
	}
	return db.Create(data).Error
}
func (t *ContractAccount) BulkInsert(db *gorm.DB, ctx context.Context, fields []string, params []map[string]interface{}) error {
	if db == nil {
		db = GetDB(ctx)
	}
	return mysql.BulkInsert(db, t.TableName(), fields, params)
}
func (t *ContractAccount) BulkSave(db *gorm.DB, ctx context.Context, fields []string, params []map[string]interface{}) error {
	if db == nil {
		db = GetDB(ctx)
	}
	return mysql.BulkSave(db, t.TableName(), fields, params)
}
func (t *ContractAccount) Delete(db *gorm.DB, ctx context.Context, where string, args ...interface{}) error {
	if db == nil {
		db = GetDB(ctx)
	}
	return db.Where(where, args...).Delete(&ContractAccount{}).Error
}
func (t *ContractAccount) Updates(db *gorm.DB, ctx context.Context, data map[string]interface{}, where string, args ...interface{}) (int64, error) {
	if db == nil {
		db = GetDB(ctx)
	}
	db = db.Model(&ContractAccount{}).Where(where, args...).Updates(data)
	return db.RowsAffected, db.Error
}
func (t *ContractAccount) GetOne(ctx context.Context, where string, args ...interface{}) (*ContractAccount, error) {
	var obj ContractAccount
	return &obj, GetDB(ctx).Where(where, args...).Take(&obj).Error
}
func (*ContractAccount) GetList(ctx context.Context, where string, args ...interface{}) ([]*ContractAccount, error) {
	var list []*ContractAccount
	return list, GetDB(ctx).Where(where, args...).Find(&list).Error
}
func (t *ContractAccount) GetListWithLimit(ctx context.Context, limit int, where string, args ...interface{}) ([]*ContractAccount, error) {
	var list []*ContractAccount
	return list, GetDB(ctx).Where(where, args...).Limit(limit).Find(&list).Error
}
func (t *ContractAccount) GetListOrderLimit(ctx context.Context, order string, limit int, where string, args ...interface{}) ([]*ContractAccount, error) {
	var list []*ContractAccount
	if limit == 0 || limit > 10000 {
		limit = 10
	}
	return list, GetDB(ctx).Where(where, args...).Order(order).Limit(limit).Find(&list).Error
}
func (t *ContractAccount) GetListPage(ctx context.Context, pageNum, limit int, where string, args ...interface{}) ([]*ContractAccount, error) {
	var list []*ContractAccount
	offset := (pageNum - 1) * limit
	return list, GetDB(ctx).Where(where, args...).Offset(offset).Limit(limit).Find(&list).Error
}
func (t *ContractAccount) GetCount(ctx context.Context, where string, args ...interface{}) (int64, error) {
	var count int64
	return count, GetDB(ctx).Model(&ContractAccount{}).Where(where, args...).Count(&count).Error
}
func (t *ContractAccount) GetSumInt64(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	type sum struct {
		Num int64 `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB(ctx).Raw(sql, args...).Scan(&obj).Error
}
func (t *ContractAccount) GetSumFloat64(ctx context.Context, sql string, args ...interface{}) (float64, error) {
	type sum struct {
		Num float64 `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB(ctx).Raw(sql, args...).Scan(&obj).Error
}
func (t *ContractAccount) GetSumDecimal(ctx context.Context, sql string, args ...interface{}) (decimal.Decimal, error) {
	type sum struct {
		Num decimal.Decimal `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB(ctx).Raw(sql, args...).Scan(&obj).Error
}
