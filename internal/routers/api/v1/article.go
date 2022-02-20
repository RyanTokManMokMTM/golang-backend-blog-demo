package v1

import (
	"github.com/RyanTokManMokMTM/blog-service/pkg/app"
	"github.com/RyanTokManMokMTM/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() *Article {
	return &Article{}
}

func (art *Article) Get(ctx *gin.Context) {
	app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
	return
}
func (art *Article) List(ctx *gin.Context)   {}
func (art *Article) Create(ctx *gin.Context) {}
func (art *Article) Update(ctx *gin.Context) {}
func (art *Article) Delete(ctx *gin.Context) {}
