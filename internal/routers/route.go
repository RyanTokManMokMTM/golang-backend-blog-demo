package routers

import (
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/RyanTokManMokMTM/blog-service/internal/middleware"
	"github.com/RyanTokManMokMTM/blog-service/internal/routers/api"
	v1 "github.com/RyanTokManMokMTM/blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	JSON WEB TOKEN Authorization theory
	Structure: Using hash to generate a token with format xxx.yyy.zzz(Header.payload.signature)
	Header
	{ //Information include
		"alg":"HS256", //using algorithm ,HMACSHA256(HS256) by default
		"typ:"JWT" //token Type
	}
	//Using base64URLEncode to this object and generate JWT header

	Payload
	{ //Storing information of jwt
      //example:
		sub :"",//subject
		aud :"",//audience
		jti :"",//jwt id
		iat :"",//Issued at
		iss :"",//Issuer - who is the publisher of the jwt
		nbf :"",//Not Before - JWT is not available before the time jwt set.
	}
	//Using base64URLEncode to this object and generate JWT Payload
	//base64URLEncode can be revered ,do not put any secure information into payload

	Signature
	{ //used to check whether header and payload was modified and using private key to sign the token
		//when generating the JWT ,it uses a specified  key(secret) and a specified algorithm(default: SHA256) to produce the signature message/info
		//MACSHA256(base64URLEncode(header).base64URLEncode(payload).secret) => JWT xxx.yyy.zzz
	}

	NOTE:base64URLEncode is base64 modified version -> Why need to be modified?
	ANS:JWT is stored inside Header or used as query parameter.In URL some character is meaningful,so base64URLEncode will use another no meaning character to instead
*/

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
	upload := NewUpload()

	//static route
	/* NOTE
		How gin serve a Static File System:???
		According to the source code of StaticFS -> be able to access $WORKDIR/relativePath
		1.StaticFS - not allow relative path with . or *
		2.Create a static Handler ：createStaticHandler -> return a handler
			1. http.StripPrefix -> return prefix path and return a handler
		3.path.Join ->join relative path with /*filepath -> /static + /*filePath => /static/*filePath
			/*filePath -> match any pattern that consisted with
			/src/*filePath
	 		/src/a/*filePath
			/src/a/b*filePath

		4.register GET and HEAD with handler that created from createStaticHandler
	*/

	route.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath)) //serving with File system

	route.GET("/auth", api.GetAuth) //testing to token service
	route.POST("/upload/file", upload.UploadFile)
	apiV1 := route.Group("/api/v1").Use(middleware.JWT())
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
		//apiV1.PATCH("/articles/:id/:state", article.Update)
		apiV1.GET("/articles/:id", article.Get)
		apiV1.GET("/articles", article.List)

	}

	return route
}
