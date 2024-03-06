package dao

import (
	"github.com/forum_backend/db"
	"github.com/forum_backend/models"
	"gorm.io/gorm"
)

type TagDao struct {
}

func (dao TagDao) AddTag(name string) error {
	tag := &models.Tag{
		Name: name,
	}
	err := db.MysqlClient.FirstOrCreate(&tag, tag).Error
	return err
}

func (dao TagDao) ListTag() ([]models.TagVo, error) {
	var tagList []models.TagVo
	err := db.MysqlClient.Model(&models.Tag{}).Select("id", "name").Find(&tagList).Error
	return tagList, err
}

func (dao TagDao) DeleteById(id uint) error {
	tag := &models.Tag{Model: gorm.Model{ID: id}}
	err := db.MysqlClient.Model(&models.Tag{}).Delete(&tag).Error
	return err
}
