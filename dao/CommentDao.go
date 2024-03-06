package dao

import (
	"github.com/forum_backend/db"
	"github.com/forum_backend/models"
)

type CommentDao struct {
}

// 添加评论
func (dao CommentDao) AddComment(articleID uint, content string, userID uint,
	replayUserID uint, rootCommentID uint) error {
	comment := &models.Comment{
		ArticleID:     articleID,
		Content:       content,
		UserID:        userID,
		ReplayUserID:  replayUserID,
		RootCommentID: rootCommentID,
	}
	err := db.MysqlClient.Create(comment).Error
	return err
}

// 一级评论
func (dao CommentDao) GetCommentsByArticleID(id uint, limit int, offset int) ([]models.CommentListItemVo, error) {
	var commentList []models.CommentListItemVo
	err := db.MysqlClient.
		Preload("User").
		Where("root_comment_id = 0 and article_id = ?", id).
		Limit(limit).
		Offset(offset).
		Find(&commentList).
		Error
	return commentList, err
}

// 子评论
func (dao CommentDao) GetChildComments(id uint) ([]models.CommentListItemVo, error) {
	var commentList []models.CommentListItemVo
	err := db.MysqlClient.
		Preload("User").
		Preload("ReplayUser").
		Where("root_comment_id = ?", id).
		Find(&commentList).
		Error
	return commentList, err
}

// 删除评论
func (dao CommentDao) DeleteCommentByID(id uint) error {
	err := db.MysqlClient.
		Delete(&models.Comment{}, id).
		Error
	return err
}
