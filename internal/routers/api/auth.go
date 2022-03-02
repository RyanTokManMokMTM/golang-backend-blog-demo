package api

import (
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/RyanTokManMokMTM/blog-service/internal/service"
	"github.com/RyanTokManMokMTM/blog-service/pkg/app"
	"github.com/RyanTokManMokMTM/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(ctx *gin.Context) {
	res := app.NewResponse(ctx)
	auth := service.AuthRequest{}
	valid, errs := app.BindAndValid(ctx, &auth)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errors:%v", errs)
		res.ToErrorResponse(errcode.InvalidParams.WithDetail(errs.Error()))
		return
	}

	serve := service.New(ctx)

	//get the key from db
	err := serve.CheckAuth(&auth)
	if err != nil {
		global.Logger.Errorf("Service.CheckAuth error: %v", err)
		res.ToErrorResponse(errcode.UnauthorizedAuthNotExist.WithDetail())
		return
	}

	//using key to generate the token
	token, err := app.GenerateToken(auth.AppKey, auth.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken error: %v", err)
		res.ToErrorResponse(errcode.UnauthorizedTokenGenerateError.WithDetail())
		return
	}

	res.ToResponse(gin.H{
		"token": token,
	})
	return
}
