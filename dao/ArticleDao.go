package dao

import (
	"github.com/forum_backend/db"
	"github.com/forum_backend/models"
	"gorm.io/gorm"
)

type ArticleDao struct {
}

// 创建新文章
func (dao ArticleDao) CreateArticle(title string, titleImage string, contentType uint,
	markDownContent string, htmlContent string, userID uint, categoryID uint, tagIDs []uint) (uint, error) {
	tags := []models.Tag{}
	for _, tagID := range tagIDs {
		tags = append(tags, models.Tag{Model: gorm.Model{ID: tagID}})
	}

	article := models.Article{
		Title:           title,
		TitleImage:      titleImage,
		ContentType:     contentType,
		MarkDownContent: markDownContent,
		HtmlContent:     htmlContent,
		UserID:          userID,
		CategoryID:      categoryID,
		Tags:            tags, // 多对多关联的Tag
	}

	err := db.MysqlClient.Create(&article).Error
	return article.ID, err
}

// 修改文章内容
func (dao ArticleDao) UpdateArticle(id uint, title string, titleImage string, contentType uint,
	markDownContent string, htmlContent string, userID uint, categoryID uint, tagIDs []uint) error {

	var tags []models.Tag
	for _, tagID := range tagIDs {
		tags = append(tags, models.Tag{Model: gorm.Model{ID: tagID}})
	}
	article := models.Article{
		Title:           title,
		TitleImage:      titleImage,
		ContentType:     contentType,
		MarkDownContent: markDownContent,
		HtmlContent:     htmlContent,
		UserID:          userID,
		CategoryID:      categoryID,
	}

	article.ID = id

	// 修改与article关联的tags, 这里不能使用Where, 必须要使用model, 并且这个model要给一个id
	// 假如我的id == 4,  tags的id为4 和 5
	// 会执行四条语句
	// 1. INSERT INTO `tag` (`created_at`,`updated_at`,`deleted_at`,`name`,`id`) VALUES ('2024-03-05 21:41:06.46','2024-03-05 21:41:06.46',NULL,'',1),('2024-03-05 21:41:06.46','2024-03-05 21:41:06.46',NULL,'',5) ON DUPLICATE KEY UPDATE `id`=`id`
	// 2. INSERT INTO `article_tags` (`article_id`,`tag_id`) VALUES (4,1),(4,5) ON DUPLICATE KEY UPDATE `article_id`=`article_id`
	// 3. UPDATE `article` SET `updated_at`='2024-03-05 21:41:06.458' WHERE `article`.`deleted_at` IS NULL AND `id` = 4
	// 4. DELETE FROM `article_tags` WHERE `article_tags`.`article_id` = 4 AND `article_tags`.`tag_id` NOT IN (1,5)
	// err := db.MysqlClient.Model(&article).Association("Tags").Replace(&tags) // 和下面的一样
	err := db.MysqlClient.Model(&models.Article{Model: gorm.Model{ID: id}}).Association("Tags").Replace(&tags)
	if err == nil {
		// 修改内容
		err = db.MysqlClient.Model(&models.Article{}).Where("id = ?", id).Updates(&article).Error
	}
	return err
}

// 获取文章详情
func (dao ArticleDao) GetArticleDetails(id uint) (*models.ArticleDetailsVo, error) {
	var detail models.ArticleDetailsVo
	err := db.MysqlClient.Model(&models.ArticleDetailsVo{}).
		Preload("User").
		Preload("Category").
		Preload("Tags").
		Where("id = ?", id).
		Select("id", "created_at", "title", "content_type", "mark_down_content", "html_content", "user_id", "category_id").
		Find(&detail).Error

	return &detail, err
}

func (dao ArticleDao) GetArticleListByUserID(userID uint, limit int, offset int) ([]models.ArticleListItemVo, error) {
	var articleList []models.ArticleListItemVo
	err := db.MysqlClient.
		Preload("User").
		Preload("Tags").
		Where("state = 1 and user_id = ?", userID).
		Limit(limit).
		Offset(offset).
		Find(&articleList).Error
	return articleList, err
}

func (dao ArticleDao) GetArticleListByCategoryID(categoryID uint, limit int, offset int) ([]models.ArticleListItemVo, error) {
	var articleList []models.ArticleListItemVo
	err := db.MysqlClient.
		Preload("User").
		Preload("Tags").
		Where("state = 1 and category_id = ?", categoryID).
		Limit(limit).
		Offset(offset).
		Find(&articleList).Error
	return articleList, err
}

func (dao ArticleDao) GetArticleListByTagID(tagID uint, limit int, offset int) ([]models.ArticleListItemVo, error) {
	var articleIDs []uint
	var articleList []models.ArticleListItemVo
	err := db.MysqlClient.Raw("select article_id from article_tags where tag_id = ?", tagID).Scan(&articleIDs).Error
	if err == nil {
		err = db.MysqlClient.
			Preload("User").
			Preload("Tags").
			Where("state = 1 and id in ?", articleIDs).
			Limit(limit).
			Offset(offset).
			Find(&articleList).Error
	}
	return articleList, err
}
