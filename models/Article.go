package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title           string `gorm:"type:varchar(64)"`  // 标题
	TitleImage      string `gorm:"type:varchar(128)"` // 标题图片
	ContentType     uint   `gorm:"type:tinyint(4)"`   // 内容的类型, 0表示markdown, 1表示html
	MarkDownContent string `gorm:"type:longtext"`     // markdown的内容
	HtmlContent     string `gorm:"type:longtext"`     // html的内容
	State           uint   `gorm:"type:tinyint(4)"`   // 文章状态
	UserID          uint   `gorm:"type:bigint"`       // 用户创建的用户ID
	Top             uint   `gorm:"type:int"`          // 置顶的权重,数字越大
	ArticleTypeID   uint   `gorm:"type:bigint"`       // 文章分类ID
}
