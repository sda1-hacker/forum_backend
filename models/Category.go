package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(32)"`
}

func (c Category) TableName() string {
	return "category"
}

type CategoryVo struct {
	ID   uint
	Name string
}

func (vo CategoryVo) TableName() string {
	return "category"
}
