package model

import "github.com/jinzhu/gorm"

type (
	ArticleTag struct {
		*Model
		TagID     string `json:"tag_id"`
		ArticleID uint8  `json:"article_id"`
	}
)

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

//This model
/* TODO -
1.Create Tag and Article id
2.Update Tag id or Article id
3.Delete Tag and Article
4.Delete Tag or Article
*/

//Create TODO - Create record with tag id and article id
func (a ArticleTag) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

//UpdateArticleID TODO - Update by article id
func (a ArticleTag) UpdateArticleID(db *gorm.DB, value interface{}) error {
	return db.Model(&a).Where("article_id = ? AND is_del = ?", a.ArticleID, 0).Update(value).Error
}

//DeleteArticleID TODO - Delete by Article ID
func (a ArticleTag) DeleteArticleID(db *gorm.DB) error {
	return db.Where("article_id = ? AND is_del = ?", a.ArticleID, 0).Delete(&a).Error
}

//Delete TODO - Delete by ID
func (a ArticleTag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", a.ID, 0).Delete(&a).Error
}
