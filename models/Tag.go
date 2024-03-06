package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"type:varchar(32)"`
}

func (t Tag) TableName() string {
	return "tag"
}

type TagVo struct {
	ID   uint
	Name string
}

func (vo TagVo) TableName() string {
	return "tag"
}
