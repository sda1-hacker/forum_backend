package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Email string
}

// gorm, 对应数据库中的表名
func (User) TableName() string {
	return "user"
}
