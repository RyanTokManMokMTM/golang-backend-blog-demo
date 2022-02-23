package service

import (
	"github.com/RyanTokManMokMTM/blog-service/internal/model"
	"github.com/RyanTokManMokMTM/blog-service/pkg/app"
)

//CountTagRequest count tag by name
type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=01"` //oneof for an element inside set
}

//TagListRequest list of tag by name
type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=01"`
}

//CreateTagRequest create a tag
type CreateTagRequest struct {
	Name     string `form:"name" binding:"required,min=3,max=100"`
	CreateBy string `form:"create_by" binding:"required,min=3,max=100"`
	State    uint8  `form:"state,default=1" binding:"oneof=01"`
}

//UpdateTagRequest update a tag by id
type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,get=1"`
	Name       string `form:"name" binding:"max=3,max=100"`
	State      uint8  `form:"state" binding:"required，oneof=01"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

//DeleteTagRequest delete a tag by id
type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

//Tag service logic

func (serve *Service) CountTag(param *CountTagRequest) (int64, error) {
	return serve.dao.CountTag(param.Name, param.State)
}
func (serve *Service) GetList(param *TagListRequest, page *app.Pager) ([]*model.Tag, error) {
	return serve.dao.GetTagLists(param.Name, param.State, page.Page, page.PageSize)
}
func (serve *Service) CreateTag(param *CreateTagRequest) error {
	return serve.dao.CreateTag(param.Name, param.State, param.CreateBy)
}
func (serve Service) UpdateTag(param *UpdateTagRequest) error {
	return serve.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}
func (serve Service) DeleteTag(param *DeleteTagRequest) error {
	return serve.dao.DeleteTag(param.ID)
}
