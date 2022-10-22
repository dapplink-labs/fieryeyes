package models

import (
	"context"

	"github.com/weblazy/easy/utils/db/mysql"
	"github.com/weblazy/easy/utils/db/mysql/mysql_config"
	"github.com/weblazy/easy/utils/elog"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var mysqlClent *mysql.MysqlClient

func NewMysqlClient() {
	cfg := mysql_config.DefaultConfig()
	cfg.DSN = "root:123456@tcp(localhost:13306)/test?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local&timeout=1s&readTimeout=3s&writeTimeout=3s"
	client, err := mysql.NewMysqlClient(cfg)
	if err != nil {
		panic(err)
	}
	mysqlClent = client
	MigrateDb(GetDB(context.Background()))
}

func MigrateDb(db *gorm.DB) error {
	if err := db.AutoMigrate(&Chain{}, &ChainAccount{}, &ContractAccount{}, &Label{}); err != nil {
		elog.ErrorCtx(context.Background(), "Failed to migrate database", zap.Error(err))
		return err
	}
	return nil
}

func GetDB(ctx context.Context) *gorm.DB {
	return mysqlClent.WithContext(ctx).DB
}
