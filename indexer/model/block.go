package model

import (
	"gorm.io/gorm"
	"math/big"
)

type Blocks struct {
	Id                int64    `json:"id" gorm:"primary_key;type:int AUTO_INCREMENT"`
	BlockHeight       *big.Int `json:"block_height" gorm:"column:block_height;type:BIGINT NOT NULL;default: 0;"`
	LatestBlockHeight *big.Int `json:"latest_block_height" gorm:"column:latest_block_height; type:BIGINT NOT NULL;default: 0;"`
	*gorm.Model
}

func (b *Blocks) TableName() string {
	return "blocks"
}

func (b Blocks) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&b).Error; err != nil {
		return err
	}
	return nil
}

func (b Blocks) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&b).Error; err != nil {
		return err
	}
	return nil
}

func (b Blocks) GetFirstColumn(db *gorm.DB) (*Blocks, error) {
	var block *Blocks
	if err := db.First(&block).Error; err != nil {
		return nil, err
	}
	return block, nil
}
