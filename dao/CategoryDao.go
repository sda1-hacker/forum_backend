package dao

import (
	"github.com/forum_backend/db"
	"github.com/forum_backend/models"
	"gorm.io/gorm"
)

type CategoryDao struct {
}

func (dao CategoryDao) AddCategory(name string) error {
	category := &models.Category{
		Name: name,
	}
	err := db.MysqlClient.FirstOrCreate(&category, category).Error
	return err
}

func (dao CategoryDao) ListCategory() ([]models.CategoryVo, error) {
	var categoryList []models.CategoryVo
	err := db.MysqlClient.Model(&models.Category{}).Select("id", "name").Find(&categoryList).Error
	return categoryList, err
}

func (dao CategoryDao) DeleteById(id uint) error {
	category := &models.Category{Model: gorm.Model{ID: id}}
	err := db.MysqlClient.Model(&models.Category{}).Delete(&category).Error
	return err
}
