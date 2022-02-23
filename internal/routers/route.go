package routers

import (
	"github.com/RyanTokManMokMTM/blog-service/internal/middleware"
	v1 "github.com/RyanTokManMokMTM/blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	route.Use(middleware.Translation()) //changing the validator to local language(zh/en)
	//swagger
	//route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//article handle obj
	article := v1.NewArticle()
	tag := v1.NewTag()
	//group the path with /api/v1
	//return a new routeGroup with custom relativePath and handlers
	apiV1 := route.Group("/api/v1")
	{
		//tags
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)          //update whole resource
		apiV1.PATCH("/tags/:id/:state", tag.Update) //update some info
		apiV1.GET("/tags", tag.List)

		//article
		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/:state", article.Update)
		apiV1.GET("/articles/:id", article.Get)
		apiV1.GET("/articles", article.Get)

	}

	return route
}
