package model

import "github.com/jinzhu/gorm"

type (
	Article struct {
		*Model
		Title         string `json:"title"`
		Desc          string `json:"desc"`
		Content       string `json:"content"`
		CoverImageUrl string `json:"cover_image_url"`
		State         uint8  `json:"state"`
	}
)

func (a Article) TableName() string {
	return "blog_article"
}

//Get -get record by id
func (a Article) Get(db *gorm.DB) (Article, error) {
	var result Article
	db = db.Where("id = ? AND state = ? AND is_del = ?", a.ID, a.State, 0)
	err := db.First(&result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return result, err
	}
	return result, nil
}

//List TODO - Get All Record with not deleted and not state 1
func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		//getting the result by offset and total result is pageSize
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	//filter the name
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}

	//filter the state
	db = db.Where("state = ?", a.State)
	err = db.Where("is_del = ?", 0).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

//Create TODO - Create article
func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

//Update TODO - Update Article
func (a Article) Update(db *gorm.DB, value interface{}) error {
	if err := db.Model(&a).Where("id = ? AND is_del = ?", a.ID, 0).Update(value).Error; err != nil {
		return err
	}
	return nil
}

//Delete TODO - Delete Article ,update is_del to 0 and state = 0
func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", a.ID, a.IsDel).Delete(&a).Error
}

//ListByTagID  need to join the article_tag table
//TODO - Testing
func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	var articles []*ArticleRow
	//change model to row struct
	//article tables fields
	selectedField := []string{
		"ar.id as article_id",
		"ar.title as article_title",
		"ar.desc as article_desc",
		"ar.cover_image_url",
		"ar.content",
	}
	//tag table fields
	selectedField = append(selectedField, []string{"t.id as tag_id", "t.name as tag_name"}...)

	//offset and limit
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	//get row data and return rows
	rows, err := db.Select(selectedField).Table(ArticleTag{}.TableName()+"AS at").
		Joins("LEFT JOIN "+Tag{}.TableName()+"AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN"+Article{}.TableName()+"AS ar ON at.article_id = ar.id").
		Where("at.tag_id = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).Rows()
	if err != nil {
		return nil, err
	}

	/*
		Include:
		(from article table)
		article_id
		article_title
		article_desc
		cover_image_url
		content
		(from tag table)
		tag_id
		tag_name
	*/
	for rows.Next() {
		row := &ArticleRow{}
		if err := rows.Scan(&row.ArticleId, &row.ArticleTitle, &row.ArticleDesc, &row.CoverImageUrl, &row.Content, &row.TagId, &row.TagName); err != nil {
			return nil, err
		} //get data from selected field to the struct

		//adding to result
		articles = append(articles, row)
	}

	return articles, nil
}

//CountByTagID need to join the article_tag table
//TODO - Not Working now
func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (uint64, error) {
	var count uint64
	err := db.Table(ArticleTag{}.TableName()+"AS at").
		Joins("LEFT JOIN"+Tag{}.TableName()+"AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN"+Article{}.TableName()+"AS ar ON at.article_id = ar.id").
		Where("at.tag_id = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

type ArticleRow struct {
	ArticleId     uint32
	ArticleTitle  string
	ArticleDesc   string
	CoverImageUrl string
	Content       string
	TagId         uint32
	TagName       string
}
