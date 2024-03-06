package models

import (
	"gorm.io/gorm"
	"time"
)

/*
仿照B设计的评论功能, 评论按照两级显示
从二级评论开始,都显示  (xxx 回复 yyy: 内容)
如果说删除一级评论,那么就会删除这个评论下的所有评论.
如果说某个三级评论回复某个二级评论, 那么删除这个二级评论不会删除对应的三级评论.
*/
type Comment struct {
	gorm.Model
	ArticleID     uint   `gorm:"type:bigint"`           // 对应的文章id
	Content       string `gorm:"type:text"`             // 评论内容
	UserID        uint   `gorm:"type:bigint"`           // 发表评论的用户id
	ReplayUserID  uint   `gorm:"type:bigint;default:0"` // 被回复评论用户的id, 如果是一级评论则为0
	RootCommentID uint   `gorm:"type:bigint;default:0"` // 根评论的id, 如果是一级评论则是0, 如果是二级以及一下评论则是一级评论的id

	User       User `gorm:"foreignKey:UserID"`
	ReplayUser User `gorm:"foreignKey:ReplayUserID"`
}

func (c Comment) TableName() string {
	return "comment"
}

type CommentListItemVo struct {
	ID            uint
	ArticleID     uint
	Content       string
	UserID        uint
	ReplayUserID  uint
	RootCommentID uint
	CreatedAt     time.Time

	ChildCount int64 // 子评论数量

	User       SimpleUserVo `gorm:"foreignKey:UserID"`
	ReplayUser SimpleUserVo `gorm:"foreignKey:ReplayUserID"`
}

func (c CommentListItemVo) TableName() string {
	return "comment"
}
