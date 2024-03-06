package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ArticleID       uint   `gorm:"type:bigint"` // 对应的文章id
	Content         string `gorm:"type:text"`   // 评论内容
	UserID          uint   `gorm:"type:bigint"` // 发表评论的id
	ReplayUserID    uint   `gorm:"type:bigint"` // 被回复的用户
	RootCommentID   uint   `gorm:"type:bigint"` // 根评论的id, 一级评论默认为0, 二级/三级评论为 一级评论的id
	ParentCommentID uint   `gorm:"type:bigint"` // 父评论的id, 一级评论默认为0
}

func (c Comment) TableName() string {
	return "comment"
}
