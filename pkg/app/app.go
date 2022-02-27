//Package app - this package is all about HTTP Response
package app

import (
	"github.com/RyanTokManMokMTM/blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx *gin.Context //current request context
}

type Pager struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	TotalRow int `json:"total_row"`
}

//NewResponse return a new response instance reference
func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

//ToResponse response a data
func (res *Response) ToResponse(data interface{}) {
	if data == nil {
		//set data to empty
		data = gin.H{}
	}

	res.Ctx.JSON(http.StatusOK, data) //set context json with data
}

//ToResponseList to response a list of data
func (res *Response) ToResponseList(list interface{}, totalRow int) {
	res.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:     GetPage(res.Ctx),     //get the page from context
			PageSize: GetPageSize(res.Ctx), //get total page from context
			TotalRow: totalRow,
		},
	})
}

//ToErrorResponse response a error (custom Error code)
func (res *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
	}

	detail := err.Detail()
	if len(detail) > 0 {
		response["detail"] = detail
	}
	res.Ctx.JSON(err.StatusCode(), response) //return an error code and info
}
