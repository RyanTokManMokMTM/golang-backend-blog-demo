//Package app - Working with server app
package app

import (
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/RyanTokManMokMTM/blog-service/pkg/convert"
	"github.com/gin-gonic/gin"
)

//GetPage get page from uri query
func GetPage(ctx *gin.Context) int {
	//get the page from url query and using custom StringConvertor
	page := convert.StrTo(ctx.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

//GetPageSize get page size from uri query
func GetPageSize(ctx *gin.Context) int {
	pageSize := convert.StrTo(ctx.Query("page_size")).MustInt()
	if pageSize <= 0 {
		//return the default page size
		return global.AppSetting.DefaultPageSize
	} else if pageSize > global.AppSetting.MaxPageSize {
		//return the default page size
		return global.AppSetting.DefaultPageSize
	}
	return pageSize
}

//GetPageOffset return the page offset
func GetPageOffset(page, pageSize int) (result int) {
	result = 0
	if page > 0 {
		//page 1 - 1 * 5(size) = 0->return page 0
		result = (page - 1) * pageSize
	}
	return
}
