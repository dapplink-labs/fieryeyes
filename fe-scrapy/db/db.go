package db

import (
	"github.com/savour-labs/fieryeyes/fe-scrapy/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Db     *gorm.DB
	Config string
}

func NewDatabase(config string) (*Database, error) {
	db, err := gorm.Open(mysql.Open(config))
	if err != nil {
		return nil, err
	}
	return &Database{
		Db:     db,
		Config: config,
	}, nil
}

func (d *Database) ConfigInfo() string {
	return d.Config
}

func (d *Database) MigrateDb() error {
	if err := d.Db.AutoMigrate(&models.Chain{}, &models.ContractAccount{}, &models.ChainAccount{}, &models.Label{}); err != nil {
		log.WithError(err).Fatal("Failed to migrate database")
		return err
	}
	return nil
}
