package service

import (
	"github.com/RyanTokManMokMTM/blog-service/internal/model"
	"github.com/RyanTokManMokMTM/blog-service/pkg/app"
)

type ArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

//CountArticleRequest count record by name
type CountArticleRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=01"` //oneof for an element inside set
}

//ArticleListRequest list of record by name
type ArticleListRequest struct {
	ID    uint32 `form:"tag_id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=01"`
}

//CreateArticleRequest Create an article
type CreateArticleRequest struct {
	TagID         uint32 `form:"tag_id" binding:"require,gte=1"`
	Title         string `form:"title" binding:"required ,min=2,max=100"`
	Desc          string `form:"desc" binding:"required"`
	Content       string `form:"content" binding:"required"`
	CoverImageURL string `form:"cover_image_url" binding:"required"`
	CreateBy      string `form:"create_by" binding:"required,min=3,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=01"`
}

//UpdateArticleRequest update article by id
type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,get=1"`
	TagID         uint32 `form:"tag_id" binding:"require,gte=1"`
	Title         string `form:"title" binding:"required ,min=2,max=100"`
	Desc          string `form:"desc" binding:"required"`
	Content       string `form:"content" binding:"required"`
	CoverImageURL string `form:"cover_image_url" binding:"required"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=3,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=01"`
}

//DeleteArticleRequest delete article by id
type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type Article struct {
	ID            uint32     `json:"id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Content       string     `json:"content"`
	CoverImageURL string     `json:"cover_image_url"`
	State         uint8      `json:"state"`
	Tag           *model.Tag `json:"tag"`
}

//HERE NEED TO USE RELATIONAL DB schema

func (serve *Service) CountArticle(param *CountArticleRequest) (int64, error) {
	return serve.dao.CountTag(param.Name, param.State)
}
func (serve *Service) ListArticle(param *ArticleListRequest, pager *app.Pager) ([]*Article, uint64, error) {
	//total record
	articleCount, err := serve.dao.CountArticleByTagID(param.ID, param.State)
	if err != nil {
		return nil, 0, err
	}

	articles, err := serve.dao.ListArticleByTagID(param.ID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var articlesList []*Article
	for _, v := range articles {
		articlesList = append(articlesList, &Article{
			ID:            v.ArticleId,
			Title:         v.ArticleTitle,
			Desc:          v.ArticleDesc,
			Content:       v.Content,
			CoverImageURL: v.CoverImageUrl,
			Tag: &model.Tag{Model: &model.Model{
				ID: v.TagId,
			}, Name: v.TagName},
		})
	}

	return articlesList, articleCount, nil
}

func (serve *Service) GetArticle(param *ArticleRequest) (*model.Article, error) {
	return serve.dao.GetArticle(param.ID, param.State)
}
func (serve *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return serve.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (serve *Service) CreateArticle(param *CreateArticleRequest) error {
	//get the record info from gorm
	article, err := serve.dao.CreateArticle(param.Title, param.Desc, param.Content, param.CoverImageURL, param.State, param.CreateBy)
	if err != nil {
		return err
	}

	//creat record for ArticleTag Table
	err = serve.dao.CreateArticleTag(article.ID, param.TagID, param.CreateBy)
	if err != nil {
		return err
	}
	return nil
	//create tag record that related to the article
}

func (serve *Service) DeleteArticle(param *DeleteArticleRequest) error {
	//need to delete article_tag table record too )
	err := serve.dao.DeleteArticle(param.ID)
	if err != nil {
		return err
	}
	err = serve.dao.DeleteArticleTag(param.ID)
	if err != nil {
		return err
	}

	return nil
}
