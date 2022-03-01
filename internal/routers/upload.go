package routers

import (
	"github.com/RyanTokManMokMTM/blog-service/global"
	service "github.com/RyanTokManMokMTM/blog-service/internal/service"
	"github.com/RyanTokManMokMTM/blog-service/pkg/app"
	"github.com/RyanTokManMokMTM/blog-service/pkg/convert"
	"github.com/RyanTokManMokMTM/blog-service/pkg/errcode"
	"github.com/RyanTokManMokMTM/blog-service/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct {
}

func NewUpload() *Upload {
	return &Upload{}
}

func (Up *Upload) UploadFile(ctx *gin.Context) {
	//get file from header
	file, header, err := ctx.Request.FormFile("file")
	//type?
	fileType := convert.StrTo(ctx.Query("type")).MustUInt32() //which type is it
	res := app.NewResponse(ctx)
	if err != nil {
		res.ToErrorResponse(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}

	//check file
	if header == nil && fileType < 0 {
		//file type staring at 1
		res.ToResponse(errcode.InvalidParams)
		return
	}

	serve := service.New(ctx)
	uploadFile, err := serve.UploadFile(upload.FileType(fileType), file, header)

	if err != nil {
		global.Logger.Errorf("Service.UploadFile error: %v", err)
		res.ToErrorResponse(errcode.ErrorUploadFileFailed.WithDetail(err.Error()))
		return
	}

	res.ToResponse(gin.H{
		"file_access_url": uploadFile.AccessURL,
	})
	//response
	return
}
