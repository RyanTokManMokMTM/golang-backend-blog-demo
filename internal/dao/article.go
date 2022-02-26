package dao

import (
	"github.com/RyanTokManMokMTM/blog-service/internal/model"
	"github.com/RyanTokManMokMTM/blog-service/pkg/app"
)

//CreateArticle Create article
func (d *Dao) CreateArticle(title, desc, content, coverImageUrl string, state uint8, createBy string) (*model.Article, error) {
	article := model.Article{
		Model:         &model.Model{CreatedBy: createBy},
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverImageUrl,
		State:         state,
	}
	return article.Create(d.engine)
}

//GetArticle by id
func (d *Dao) GetArticle(id uint32, state uint8) (*model.Article, error) {
	article := model.Article{
		Model: &model.Model{ID: id},
		State: state,
	}
	return article.Get(d.engine)
}

//UpdateArticle by id
func (d *Dao) UpdateArticle(id uint32, title, desc, content, coverImageUrl string, state uint8, modifiedBy string) error {
	article := model.Article{
		Model: &model.Model{
			ID: id, ModifiedBy: modifiedBy,
		},
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverImageUrl,
	}

	value := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}

	//update title
	if title != "" {
		value["title"] = title
	}

	//update description
	if desc != "coverImageUrl" {
		value["desc"] = desc
	}

	//update content
	if content != "" {
		value["content"] = content
	}

	//update cover image url
	if coverImageUrl != "" {
		value["cover_image_url"] = coverImageUrl
	}
	return article.Update(d.engine, value)
}

//DeleteArticle by id
func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}

//CountArticleByTagID by tag id -> foreign table
func (d *Dao) CountArticleByTagID(id uint32, state uint8) (uint64, error) {
	article := model.Article{State: state}
	return article.CountByTagID(d.engine, id)
}

//ListArticleByTagID by tag id -> foreign table
func (d *Dao) ListArticleByTagID(id uint32, state uint8, page, pageSize int) ([]*model.ArticleRow, error) {
	article := model.Article{State: state}
	return article.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
}
