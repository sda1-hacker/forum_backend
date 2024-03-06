package dao

import (
	"errors"
	"github.com/forum_backend/db"
	"github.com/forum_backend/models"
	"github.com/forum_backend/utils"
)

type UserDao struct {
}

// 新增用户
func (dao UserDao) CreateUser(nickName string, email string, password string, imageUrl string, gender uint) (uint, error) {
	user := &models.User{
		NickName: nickName,
		Email:    email,
		Password: password,
		Image:    imageUrl,
		Gender:   gender,
		Status:   0,
	}

	var ctn int64
	err := db.MysqlClient.Raw("select count(*) from user where email = ?", email).Scan(&ctn).Error
	if err == nil {
		if ctn == 0 {
			err = db.MysqlClient.Create(user).Error
		}
	}
	return user.ID, err
}

// 根据id查找用户
func (dao UserDao) GetUserById(id uint) (*models.UserDetailsInfoVo, error) {
	user := &models.UserDetailsInfoVo{}
	err := db.MysqlClient.
		Raw("select * from user where status = 0 and deleted_at is null and id = ?", id).
		Scan(user).
		Error
	return user, err
}

// 根据邮箱查找用户
func (dao UserDao) GetUserByEmail(email string) (*models.UserDetailsInfoVo, error) {
	user := &models.UserDetailsInfoVo{}
	err := db.MysqlClient.Where("status = 0 and email = ?", email).Find(user).Error
	return user, err
}

// 修改用户状态
func (dao UserDao) updateStatus(userId uint, status uint) error {
	err := db.MysqlClient.Model(&models.User{}).
		Where("id = ?", userId).
		Updates(models.User{Status: status}).
		Error
	return err
}

// 修改用户密码
func (dao UserDao) UpdatePassword(id uint, currentPassword string, newPassword string) error {
	var passwd string
	err := db.MysqlClient.Model(&models.User{}).
		Where("id = ?", id).
		Select("password").
		Find(&passwd).
		Error
	if err == nil {
		if currentPassword == passwd {
			err = db.MysqlClient.Model(&models.User{}).
				Where("id = ?", id).
				Update("password", newPassword).
				Error
			return err
		} else {
			return errors.New("原密码不正确··")
		}
	}
	return err
}

// 查找用户个人中心 -- utils.Select() 默认查找10条
func (dao UserDao) GetUserCenter(id uint) (*models.UserCenterVo, error) {
	var userCenter models.UserCenterVo
	err := db.MysqlClient.Model(&models.UserCenterVo{}).
		Preload("Articles", utils.LimitAndOffset(10, 0)).
		Preload("Articles.Tags").
		Where("id = ?", id).
		Find(&userCenter).
		Error
	return &userCenter, err
}
