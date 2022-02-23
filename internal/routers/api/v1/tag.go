package v1

import (
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/RyanTokManMokMTM/blog-service/internal/service"
	"github.com/RyanTokManMokMTM/blog-service/pkg/app"
	"github.com/RyanTokManMokMTM/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type (
	Tag struct {
	}
)

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) Get(ctx *gin.Context) {}

func (t *Tag) Create(ctx *gin.Context) {}

func (t *Tag) Update(ctx *gin.Context) {}

func (t *Tag) Delete(ctx *gin.Context) {}

func (t *Tag) List(ctx *gin.Context) {
	//decode the request
	param := service.TagListRequest{}
	res := app.NewResponse(ctx)
	vaild, errs := app.BindAndValid(ctx, &param)
	if !vaild {
		//failed
		global.Logger.Errorf("App.BindAndValid Error:%v", errs)
		res.ToErrorResponse(errcode.InvalidParams.WithDetail(errs.Error()))
		return
	}
	//handling the request and services
	svc := service.New(ctx)
	pager := &app.Pager{Page: app.GetPage(ctx), PageSize: app.GetPageSize(ctx)}
	totalRow, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf("service.CountTagRequest error: %v", err)
		res.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tagList, err := svc.GetList(&param, pager)
	if err != nil {
		global.Logger.Errorf("service.GetList error: %v", err)
		res.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	res.ToResponseList(tagList, int(totalRow))
	return
}
