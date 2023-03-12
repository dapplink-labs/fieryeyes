package models

import (
	"gorm.io/gorm"
)

type ShadowScore struct {
	Id              uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CollectionId    uint64 `json:"collect_id"`
	BlueChip        string `gorm:"type:varchar(256)"  json:"blue_chip"`
	Fluidity        string `gorm:"type:varchar(256)"  json:"fluidity"`
	Reliability     string `gorm:"type:varchar(256)"  json:"reliability"`
	CommunityActive string `gorm:"type:varchar(256)"  json:"communityActive"`
	Heat            string `gorm:"type:varchar(256)"  json:"heat"`
	PotentialIncome string `gorm:"type:varchar(256)"  json:"potentialIncome"`
	*gorm.Model
}

func (ss *ShadowScore) TableName() string {
	return "shadow_score"
}

func (ss *ShadowScore) SelfInsert(db *gorm.DB) error {
	if err := db.Create(&ss).Error; err != nil {
		return err
	}
	return nil
}

func (ss *ShadowScore) SelfUpdate(db *gorm.DB) error {
	if err := db.Updates(&ss).Error; err != nil {
		return err
	}
	return nil
}
