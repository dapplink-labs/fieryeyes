package models

import (
	"gorm.io/gorm"
)

type Blocks struct {
	Id                int64  `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	BlockHeight       uint64 `json:"block_height" gorm:"column:block_height;default: 0;"`
	BlockHash         string `json:"block_hash" gorm:"column:block_hash;default: '';"`
	ParentHash        string `json:"parent_hash" gorm:"column:parent_hash;default: '';"`
	LatestBlockHeight uint64 `json:"latest_block_height" gorm:"column:latest_block_height;default: 0;"`
	*gorm.Model
}

func (b *Blocks) TableName() string {
	return "blocks"
}

func (b *Blocks) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&b).Error; err != nil {
		return err
	}
	return nil
}

func (b *Blocks) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&b).Error; err != nil {
		return err
	}
	return nil
}

func (b *Blocks) GetFirstColumn(db *gorm.DB) (*Blocks, error) {
	var block *Blocks
	if err := db.Order("id desc").First(&block).Error; err != nil {
		return nil, err
	}
	return block, nil
}

func (b *Blocks) ExistBlock(db *gorm.DB) bool {
	var count int64
	db.Find(&b).Count(&count)
	if count > 0 {
		return true
	}
	return false
}
