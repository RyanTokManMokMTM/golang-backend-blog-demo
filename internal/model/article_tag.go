package model

import (
	"github.com/jinzhu/gorm"
)

type (
	ArticleTag struct {
		*Model
		TagID     uint32 `json:"tag_id"`
		ArticleID uint32 `json:"article_id"`
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
	if err := db.Create(&a).Error; err != nil {
		return err
	}
	return nil
}

//GetByAID by article ID
func (a ArticleTag) GetByAID(db *gorm.DB) (ArticleTag, error) {
	var result ArticleTag
	if err := db.Where("article_id = ? AND is_del = ?", a.ArticleID, 0).First(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

//GetByTID by Tag ID
func (a ArticleTag) GetByTID(db *gorm.DB) (ArticleTag, error) {
	var result ArticleTag
	if err := db.Where("tag_id = ? AND is_del = ?", a.TagID, 0).First(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

//GetListByAIDs - by a group of article
func (a ArticleTag) GetListByAIDs(db *gorm.DB, articles []uint32) ([]*ArticleTag, error) {
	var results []*ArticleTag
	err := db.Where("article_id IN (?)  AND is_del = ?", articles, 0).Find(&results).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return results, nil
}

//GetListByTID by tag id
func (a ArticleTag) GetListByTID(db *gorm.DB) ([]*ArticleTag, error) {
	var results []*ArticleTag
	if err := db.Where("tag_id = ? AND is_del = ?", a.TagID, 0).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

//UpdateOne TODO - Update by article id and only one record
func (a ArticleTag) UpdateOne(db *gorm.DB, value interface{}) error {
	if err := db.Model(&a).Where("article_id = ? AND is_del = ?", a.ArticleID, 0).Limit(1).Update(value).Error; err != nil {
		return err
	}
	return nil
	//return db.Model(&a).Where("article_id = ? AND is_del = ?", a.ArticleID, 0).Update(value).Error
}

//Delete TODO - Delete by ID
func (a ArticleTag) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

//DeleteOne TODO - Delete by article_id and only one record
func (a ArticleTag) DeleteOne(db *gorm.DB) error {
	return db.Where("article_id = ? AND is_del = ?", a.ArticleID, 0).Delete(&a).Limit(1).Error
}
