package utils

import (
	"github.com/forum_backend/db"
	"github.com/forum_backend/models"
)

type GormUtils struct {
}

// 根据models创建对应的数据库
func (utils GormUtils) TableGenerator() {
	db.MysqlConnection.AutoMigrate(&models.User{})
}
