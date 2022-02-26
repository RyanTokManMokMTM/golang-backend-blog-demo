package dao

import "github.com/RyanTokManMokMTM/blog-service/internal/model"

//GetArticleTagByAID -  by an article ID
func (d *Dao) GetArticleTagByAID(articleID uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleID: articleID}
	return articleTag.GetByAID(d.engine)
}

//GetArticleTagListByTID - by a Tag ID
func (d *Dao) GetArticleTagListByTID(TagID uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{TagID: TagID}
	return articleTag.GetByTID(d.engine)
}

//GetArticleTagListByAIDs - by an article id
func (d *Dao) GetArticleTagListByAIDs(articleID []uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{}
	return articleTag.GetListByAIDs(d.engine, articleID)
}

//CreateArticleTag - create ArticleTag record
func (d *Dao) CreateArticleTag(articleID, tagID uint32, createBY string) error {
	articleTag := model.ArticleTag{ArticleID: articleID, TagID: tagID, Model: &model.Model{CreatedBy: createBY}}
	return articleTag.Create(d.engine)
}

//UpdateArticleTag - update ArticleTag record
func (d *Dao) UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{ArticleID: articleID}
	value := map[string]interface{}{
		"article_id":  articleID,
		"tag_id":      tagID,
		"modified_by": modifiedBy,
	}
	return articleTag.UpdateOne(d.engine, value)
}

//DeleteArticleTag - delete ArticleTag record
func (d *Dao) DeleteArticleTag(articleID uint32) error {
	articleTag := model.ArticleTag{ArticleID: articleID}
	return articleTag.DeleteOne(d.engine)
}
