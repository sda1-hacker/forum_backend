package dao

import (
	"github.com/forum_backend/db"
	"github.com/forum_backend/models"
)

type UserDao struct {
}

func (dao UserDao) AddUser(name string, email string) (uint, error) {
	user := models.User{
		Name:  name,
		Email: email,
	}
	err := db.MysqlConnection.Create(&user).Error
	// insert 主键回填
	return user.ID, err
}

func (dao UserDao) fundUserById(id int) (*models.User, error) {
	user := &models.User{}
	err := db.MysqlConnection.Raw("select * from user where id = ?", id).Scan(user).Error
	return user, err
}
