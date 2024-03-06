package dao

import (
	"github.com/forum_backend/db"
	"github.com/forum_backend/models"
)

type CommentDao struct {
}

// 添加评论
func (dao CommentDao) AddComment(articleID uint, content string, userID uint,
	replayUserID uint, rootCommentID uint, parentCommentID uint) error {

	comment := &models.Comment{
		ArticleID:       articleID,
		Content:         content,
		UserID:          userID,
		ReplayUserID:    replayUserID,
		RootCommentID:   rootCommentID,
		ParentCommentID: parentCommentID,
	}
	err := db.MysqlClient.Create(comment).Error
	return err
}
