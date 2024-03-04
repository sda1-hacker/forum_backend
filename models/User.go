package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NickName string `gorm:"type:varchar(32)"`  // 昵称
	Email    string `gorm:"type:varchar(64)"`  // 邮箱, 使用邮箱登陆, 验证
	Password string `gorm:"type:varchar(32)"`  // 密码
	Image    string `gorm:"type:varchar(128)"` // 头像
	Gender   uint   `gorm:"type:tinyint(4)"`   // 性别 0表示女,1表示男
	Status   uint   `gorm:"type:tinyint(4)"`   // 账号状态, 0表示正常,1表示禁止登陆

}

// gorm, 对应数据库中的表名
func (u User) TableName() string {
	return "user"
}

// userVo对象
type UserVo struct {
	ID       uint
	NickName string
}

func (vo UserVo) TableName() string {
	return "user"
}
