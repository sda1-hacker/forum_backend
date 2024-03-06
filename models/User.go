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

	Articles []Article `gorm:"foreignKey:UserID"`
}

func (u User) TableName() string {
	return "user"
}

// 用户详情信息
type UserDetailsInfoVo struct {
	ID       uint
	NickName string
	Email    string
	Password string
	Image    string
	Gender   string
}

func (vo UserDetailsInfoVo) TableName() string {
	return "user"
}

// 个人主页
type UserCenterVo struct {
	ID       uint
	NickName string
	Image    string
	Gender   uint
	Articles []ArticleListItemVo `gorm:"foreignKey:UserID"`
}

func (vo UserCenterVo) TableName() string {
	return "user"
}

type SimpleUserVo struct {
	ID       uint
	NickName string
	Image    string
	Gender   uint
}

func (vo SimpleUserVo) TableName() string {
	return "user"
}
