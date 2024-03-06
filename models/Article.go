package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	Title           string `gorm:"type:varchar(64)"`  // 标题
	TitleImage      string `gorm:"type:varchar(128)"` // 标题图片
	ContentType     uint   `gorm:"type:tinyint(4)"`   // 内容的类型, 0表示markdown, 1表示html
	MarkDownContent string `gorm:"type:longtext"`     // markdown的内容
	HtmlContent     string `gorm:"type:longtext"`     // html的内容
	State           uint   `gorm:"type:tinyint(4)"`   // 文章状态, 0表示私有,1表示公开
	UserID          uint   `gorm:"type:bigint"`       // 用户创建的用户ID
	Top             uint   `gorm:"type:int"`          // 置顶的权重,数字越大
	CategoryID      uint   `gorm:"type:bigint"`       // 文章分类ID

	User     User     `gorm:"foreignKey:UserID"`
	Category Category `gorm:"foreignKey:CategoryID"`
	Tags     []Tag    `gorm:"many2many:article_tags;"` // many2many
	// 使用many2many的时候会创建一个表, 表的名字为 article_tags, 仅有 article_id, tags_id 两个字段 作为联合主键
}

func (a Article) TableName() string {
	return "article"
}

type ArticleDetailsVo struct {
	ID              uint
	CreatedAt       time.Time
	Title           string
	ContentType     uint
	MarkDownContent string
	HtmlContent     string
	UserID          uint
	CategoryID      uint

	// User     User     `gorm:"foreignKey:UserID"`
	User     SimpleUserVo `gorm:"foreignKey:UserID"`
	Category CategoryVo   `gorm:"foreignKey:CategoryID"`
	// 参考这里吧,多对多,重写join表的外键和关联:  https://blog.csdn.net/kingsill/article/details/134796092
	Tags []Tag `gorm:"many2many:article_tags;joinForeignKey:article_id;joinReferences:tag_id"`
}

func (vo ArticleDetailsVo) TableName() string {
	return "article"
}

type ArticleListItemVo struct {
	ID         uint
	Title      string
	TitleImage string
	UserID     uint

	// User User
	User SimpleUserVo `gorm:"foreignKey:UserID"`
	Tags []TagVo      `gorm:"many2many:article_tags;joinForeignKey:article_id;joinReferences:tag_id"`
	// 使用many2many的时候会创建一个表, 表的名字为 article_tags, 仅有 article_id, tags_id 两个字段 作为联合主键
}

func (vo ArticleListItemVo) TableName() string {
	return "article"
}
