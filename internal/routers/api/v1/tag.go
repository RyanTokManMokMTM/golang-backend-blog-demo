package v1

import (
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/RyanTokManMokMTM/blog-service/internal/service"
	"github.com/RyanTokManMokMTM/blog-service/pkg/app"
	"github.com/RyanTokManMokMTM/blog-service/pkg/convert"
	"github.com/RyanTokManMokMTM/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"log"
)

type (
	Tag struct {
	}
)

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) Get(ctx *gin.Context) {
	//param := service.
}

func (t *Tag) Create(ctx *gin.Context) {
	param := service.CreateTagRequest{}
	res := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errors:%v", errs)
		res.ToErrorResponse(errcode.InvalidParams.WithDetail(errs.Error()))
		return
	}

	svc := service.New(ctx)
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf("Service.CreateTag errors:%v", errs)
		res.ToErrorResponse(errcode.ErrorCreateTagFail.WithDetail(err.Error()))
		return
	}

	res.ToResponse(gin.H{})
	return
}

func (t *Tag) Update(ctx *gin.Context) {
	param := service.UpdateTagRequest{
		ID: convert.StrTo(ctx.Param("id")).MustUInt32(),
	}

	res := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error:%s", errs)
		res.ToErrorResponse(errcode.InvalidParams.WithDetail(errs.Error()))
		return
	}

	svc := service.New(ctx)
	if err := svc.UpdateTag(&param); err != nil {
		global.Logger.Errorf("Service.UpdateTag error:%v", err)
		res.ToErrorResponse(errcode.ErrorUpdateTagFail.WithDetail(err.Error()))
		return
	}

	res.ToResponse(gin.H{})
	return
}

func (t *Tag) Delete(ctx *gin.Context) {
	param := service.DeleteTagRequest{
		ID: convert.StrTo(ctx.Param("id")).MustUInt32(),
	}

	log.Println(param.ID)
	res := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errors:%v", errs)
		res.ToErrorResponse(errcode.InvalidParams.WithDetail(errs.Error()))
		return
	}

	svc := service.New(ctx)
	if err := svc.DeleteTag(&param); err != nil {
		global.Logger.Errorf("Service.DeleteTag error:%v", err)
		res.ToErrorResponse(errcode.ErrorDeleteTagFail.WithDetail(err.Error()))
		return
	}

	res.ToResponse(gin.H{})
	return

}

func (t *Tag) List(ctx *gin.Context) {
	//decode the request
	param := service.TagListRequest{}
	res := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
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
