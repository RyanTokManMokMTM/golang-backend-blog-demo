package v1

import (
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/RyanTokManMokMTM/blog-service/internal/service"
	"github.com/RyanTokManMokMTM/blog-service/pkg/app"
	"github.com/RyanTokManMokMTM/blog-service/pkg/convert"
	"github.com/RyanTokManMokMTM/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() *Article {
	return &Article{}
}

func (art *Article) Get(ctx *gin.Context) {
	//by id
	param := service.ArticleRequest{ID: convert.StrTo(ctx.Param("id")).MustUInt32()}
	res := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error :%v", errs)
		res.ToErrorResponse(errcode.InvalidParams.WithDetail(errs.Error()))
		return
	}

	serve := service.New(ctx)
	article, err := serve.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf("Service.GetArticle error:%v", err)
		res.ToErrorResponse(errcode.ErrorGetArticleFailed.WithDetail(err.Error()))
		return
	}

	res.ToResponse(article)
	return
}
func (art *Article) List(ctx *gin.Context) {
	var param service.ArticleListRequest
	res := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error :%v", errs)
		res.ToErrorResponse(errcode.InvalidParams.WithDetail(errs.Error()))
		return
	}

	//get the page from uri
	pager := app.Pager{Page: app.GetPage(ctx), PageSize: app.GetPageSize(ctx)}
	serve := service.New(ctx)
	article, totalRow, err := serve.GetArticleList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("Service.GetArticleList error :%v", errs)
		res.ToErrorResponse(errcode.ErrorGetArticleListFailed.WithDetail(errs.Error()))
		return
	}

	res.ToResponseList(article, int(totalRow))
	return
}
func (art *Article) Create(ctx *gin.Context) {
	param := service.CreateArticleRequest{}
	res := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error :%v", errs)
		res.ToErrorResponse(errcode.InvalidParams.WithDetail(errs.Error()))
		return
	}

	serve := service.New(ctx)
	err := serve.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf("Service.CreateArticle error :%v", errs)
		res.ToErrorResponse(errcode.ErrorCreateArticleFailed.WithDetail(errs.Error()))
		return
	}

	res.ToResponse(gin.H{})
	return
}
func (art *Article) Update(ctx *gin.Context) {
	//by id
	param := service.UpdateArticleRequest{ID: convert.StrTo(ctx.Param("id")).MustUInt32()}
	res := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error :%v", errs)
		res.ToErrorResponse(errcode.InvalidParams.WithDetail(errs.Error()))
		return
	}

	serve := service.New(ctx)
	if err := serve.UpdateArticle(&param); err != nil {
		global.Logger.Errorf("Service.UpdateArticle error :%v", err)
		res.ToErrorResponse(errcode.ErrorUpdateArticleFailed.WithDetail(err.Error()))
		return
	}

	res.ToResponse(gin.H{})
	return
}
func (art *Article) Delete(ctx *gin.Context) {
	param := service.DeleteArticleRequest{ID: convert.StrTo(ctx.Param("id")).MustUInt32()}
	res := app.NewResponse(ctx)
	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid error :%v", errs)
		res.ToErrorResponse(errcode.InvalidParams.WithDetail(errs.Error()))
		return
	}

	serve := service.New(ctx)
	if err := serve.DeleteArticle(&param); err != nil {
		global.Logger.Errorf("Service.DeleteArticle error:%v", err)
		res.ToErrorResponse(errcode.ErrorDeleteArticleFailed.WithDetail(err.Error()))
		return
	}
	res.ToResponse(gin.H{})
	return
}
