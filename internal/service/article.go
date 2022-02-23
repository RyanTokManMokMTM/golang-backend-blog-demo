package service

//CountArticleRequest count record by name
type CountArticleRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=01"` //oneof for an element inside set
}

//ArticleListRequest list of record by name
type ArticleListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State int    `form:"state,default=1" binding:"oneof=01"`
}

//CreateArticleRequest Create an article
type CreateArticleRequest struct {
	Name     string `form:"name" binding:"required,min=3,max=100"`
	CreateBy string `form:"create_by" binding:"required,min=3,max=100"`
	State    uint8  `form:"state,default=1" binding:"oneof=01"`
}

//UpdateArticleRequest update article by id
type UpdateArticleRequest struct {
	ID         uint32 `form:"id" binding:"required,get=1"`
	Name       string `form:"name" binding:"max=3,max=100"`
	State      uint8  `form:"state" binding:"required，oneof=01"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

//DeleteArticleRequest delete article by id
type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
