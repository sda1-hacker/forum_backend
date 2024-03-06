package main

import (
	"fmt"
	"github.com/forum_backend/config"
	"github.com/forum_backend/dao"
	"github.com/forum_backend/db"
	"github.com/forum_backend/models"
	"github.com/forum_backend/utils"
	"gorm.io/gorm"
)

func main() {
	// 初始化配置
	config.InitConfig()

	db.InitMysqlConnection()

	articles, err := dao.ArticleDao{}.GetArticleListByTagID(1, 10, 0)
	if err == nil {
		for _, article := range articles {
			fmt.Printf("===> %#v \n", article)

		}
	}
}

func AddCategory() {
	dao.CategoryDao{}.AddCategory("闲聊")
	dao.CategoryDao{}.AddCategory("八卦")
	dao.CategoryDao{}.AddCategory("游戏")
	dao.CategoryDao{}.AddCategory("大数据")
	dao.CategoryDao{}.AddCategory("人工智能")
}

func UpdateArticle() {
	dao.ArticleDao{}.UpdateArticle(4, "能否直接修改tags??..", "", 2, "", "", 2, 5, []uint{1, 5})
}

func CreateArticle() {
	dao.ArticleDao{}.CreateArticle("这是一个测试的Title", "", 1, "", "", 1, 3, []uint{3, 4})
}

func Delete(id uint) {
	category := &models.Category{Model: gorm.Model{ID: id}}
	err := db.MysqlClient.Model(&models.Category{}).Delete(&category).Error
	fmt.Println(err)
}

func CreateTable() {
	db.MysqlClient.AutoMigrate(&models.Category{})
}

// 多对多插入
func AddArticle() {
	art := models.Article{
		Title:  "又一天",
		UserID: 1,
		State:  0,
		Tags: []models.Tag{ // 会根据ID自动插入到关联的表中
			{
				Model: gorm.Model{
					ID: 1,
				},
			},
			{
				Model: gorm.Model{
					ID: 2,
				},
			},
		},
	}
	db.MysqlClient.Model(&models.Article{}).Create(&art)
}

func UserCenter() {
	var userCenterVo models.UserCenterVo
	db.MysqlClient.Model(&models.User{}).
		//Preload("Articles", func(db *gorm.DB) *gorm.DB {
		//	return db.Select().Limit(2)
		//}).
		//Preload("Articles.Tags", func(db *gorm.DB) *gorm.DB {
		//	return db.Select("id", "name")
		//}).
		Preload("Articles", utils.Select("id", "user_id", "title")).
		Where("id = ?", 2).
		Find(&userCenterVo)
	fmt.Printf("%#v \n", userCenterVo)
}
