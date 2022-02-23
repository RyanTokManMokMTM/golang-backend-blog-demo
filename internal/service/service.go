//Package service - Process the request and using dao object
package service

import (
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/RyanTokManMokMTM/blog-service/internal/dao"
	"github.com/gin-gonic/gin"
)

type Service struct {
	ctx *gin.Context
	dao *dao.Dao
}

//New - Return the service instance with gin.context and dao object
func New(ctx *gin.Context) Service {
	return Service{ctx: ctx, dao: dao.New(global.DBEngine)}
}
