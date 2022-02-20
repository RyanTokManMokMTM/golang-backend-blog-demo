package v1

import "github.com/gin-gonic/gin"

type (
	Tag struct {
	}
)

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) Get(ctx *gin.Context)    {}
func (t *Tag) List(ctx *gin.Context)   {}
func (t *Tag) Create(ctx *gin.Context) {}
func (t *Tag) Update(ctx *gin.Context) {}
func (t *Tag) Delete(ctx *gin.Context) {}
