package model

type (
	ArticleTag struct {
		*Model
		TagID     string `json:"tag_id"`
		ArticleID uint8  `json:"article_id"`
	}
)

func (t ArticleTag) TableName() string {
	return "blog_article_tag"
}
