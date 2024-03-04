package models

import "gorm.io/gorm"

type ArticleType struct {
	gorm.Model
	Name string `gorm:"type:varchar(32)"`
}
