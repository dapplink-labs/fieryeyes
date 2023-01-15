package models

import (
	"gorm.io/gorm"
)

type Chain struct {
	Name string `gorm:"type:text;description:Name; comment:链名称"   json:"name"`
	*gorm.Model
}
