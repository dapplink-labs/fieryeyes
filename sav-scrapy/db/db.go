package db

import (
	"github.com/savour-labs/fieryeyes/sav-scrapy/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	db     *gorm.DB
	config string
}

func NewDatabase(config string) (*Database, error) {
	db, err := gorm.Open(mysql.Open(config))
	if err != nil {
		return nil, err
	}
	return &Database{
		db:     db,
		config: config,
	}, nil
}

func (d *Database) Config() string {
	return d.config
}

func (d *Database) MigrateDb() error {
	if err := d.db.AutoMigrate(&models.Chain{}, &models.ContractAccount{}, &models.ChainAccount{}); err != nil {
		log.WithError(err).Fatal("Failed to migrate database")
		return err
	}
	return nil
}
