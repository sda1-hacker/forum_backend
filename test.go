package main

import (
	"fmt"
	"github.com/forum_backend/config"
	"github.com/forum_backend/db"
	"github.com/forum_backend/models"
)

func main() {
	// 初始化配置
	config.InitConfig()

	db.InitMysqlConnection()

	SelectData()
	// InsertData()
}

func SelectData() {
	var comm []models.CommentVo
	//db.MysqlConnection.
	//	Preload("User", func(db *gorm.DB) *gorm.DB {   // 指定Model, select fields from user
	//		return db.Model(&models.User{})
	//	}).
	//	Preload("ReplayUser", func(db *gorm.DB) *gorm.DB {
	//		return db.Model(&models.User{})
	//	}).
	//	Model(&models.Comment{}).
	//	Find(&comm)

	db.MysqlConnection.
		Model(&models.Comment{}).
		Preload("User"). // 为Vo设置了TableName, 这样查询指定的是 select * from user
		Preload("ReplayUser").
		Find(&comm)

	fmt.Printf("resdata ===> %#v", comm)
}

func InsertData() {
	//db.MysqlConnection.
	//	Create(&models.User{
	//		NickName: "赵六",
	//		Email:    "lisi@gmail.com",
	//		Password: "123456",
	//		Gender:   1,
	//		Status:   0,
	//	})

	db.MysqlConnection.
		Create(&models.Comment{
			ArticleID:       2,
			Content:         "第二个内容..",
			UserID:          2,
			ReplayUserID:    3,
			RootCommentID:   0,
			ParentCommentID: 0,
		})
}
