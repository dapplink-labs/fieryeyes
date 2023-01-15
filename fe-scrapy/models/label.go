package models

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
	"github.com/weblazy/easy/utils/db/mysql"
	"gorm.io/gorm"
)

var LabelHandler = &Label{}

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

func (t *Label) TableName() string {
	return "label"
}

func (t *Label) Insert(db *gorm.DB, ctx context.Context, data *Label) error {
	if db == nil {
		db = GetDB(ctx)
	}
	return db.Create(data).Error
}
func (t *Label) BulkInsert(db *gorm.DB, ctx context.Context, fields []string, params []map[string]interface{}) error {
	if db == nil {
		db = GetDB(ctx)
	}
	return mysql.BulkInsert(db, t.TableName(), fields, params)
}
func (t *Label) BulkSave(db *gorm.DB, ctx context.Context, fields []string, params []map[string]interface{}) error {
	if db == nil {
		db = GetDB(ctx)
	}
	return mysql.BulkSave(db, t.TableName(), fields, params)
}
func (t *Label) Delete(db *gorm.DB, ctx context.Context, where string, args ...interface{}) error {
	if db == nil {
		db = GetDB(ctx)
	}
	return db.Where(where, args...).Delete(&Label{}).Error
}
func (t *Label) Updates(db *gorm.DB, ctx context.Context, data map[string]interface{}, where string, args ...interface{}) (int64, error) {
	if db == nil {
		db = GetDB(ctx)
	}
	db = db.Model(&Label{}).Where(where, args...).Updates(data)
	return db.RowsAffected, db.Error
}
func (t *Label) GetOne(ctx context.Context, where string, args ...interface{}) (*Label, error) {
	var obj Label
	return &obj, GetDB(ctx).Where(where, args...).Take(&obj).Error
}
func (*Label) GetList(ctx context.Context, where string, args ...interface{}) ([]*Label, error) {
	var list []*Label
	return list, GetDB(ctx).Where(where, args...).Find(&list).Error
}
func (t *Label) GetListWithLimit(ctx context.Context, limit int, where string, args ...interface{}) ([]*Label, error) {
	var list []*Label
	return list, GetDB(ctx).Where(where, args...).Limit(limit).Find(&list).Error
}
func (t *Label) GetListOrderLimit(ctx context.Context, order string, limit int, where string, args ...interface{}) ([]*Label, error) {
	var list []*Label
	if limit == 0 || limit > 10000 {
		limit = 10
	}
	return list, GetDB(ctx).Where(where, args...).Order(order).Limit(limit).Find(&list).Error
}
func (t *Label) GetListPage(ctx context.Context, pageNum, limit int, where string, args ...interface{}) ([]*Label, error) {
	var list []*Label
	offset := (pageNum - 1) * limit
	return list, GetDB(ctx).Where(where, args...).Offset(offset).Limit(limit).Find(&list).Error
}
func (t *Label) GetCount(ctx context.Context, where string, args ...interface{}) (int64, error) {
	var count int64
	return count, GetDB(ctx).Model(&Label{}).Where(where, args...).Count(&count).Error
}
func (t *Label) GetSumInt64(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	type sum struct {
		Num int64 `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB(ctx).Raw(sql, args...).Scan(&obj).Error
}
func (t *Label) GetSumFloat64(ctx context.Context, sql string, args ...interface{}) (float64, error) {
	type sum struct {
		Num float64 `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB(ctx).Raw(sql, args...).Scan(&obj).Error
}
func (t *Label) GetSumDecimal(ctx context.Context, sql string, args ...interface{}) (decimal.Decimal, error) {
	type sum struct {
		Num decimal.Decimal `json:"num" gorm:"column:num"`
	}
	var obj sum
	return obj.Num, GetDB(ctx).Raw(sql, args...).Scan(&obj).Error
}
