package db

import (
	"context"
	"fmt"
	"github.com/savour-labs/fieryeyes/fe-service/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     uint64
	DbName   string
}

type Database struct {
	Ctx context.Context
	Db  *gorm.DB
	Cfg *DatabaseConfig
}

func NewDatabase(ctx context.Context, cfg *DatabaseConfig) (*Database, error) {
	dsnTemplate := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local"
	dsn := fmt.Sprintf(dsnTemplate, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	return &Database{
		Ctx: ctx,
		Db:  db,
		Cfg: cfg,
	}, nil
}

func (d *Database) MigrateDb() error {
	if err := d.Db.AutoMigrate(
		&models.Holders{},
		&models.Chain{},
		&models.Collection{},
		&models.CollectionStat{},
		&models.CollectionWhale{},
		&models.HolderStat{},
		&models.MainToken{},
		&models.Nft{},
		&models.NftHolder{},
		&models.NftStat{},
		&models.NftTxn{},
		&models.ShadowScore{},
		&models.TokenPrice{},
	); err != nil {
		log.WithError(err).Fatal("Failed to migrate database")
		return err
	}
	return nil
}
