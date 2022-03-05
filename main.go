package main

import (
	"fmt"
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/RyanTokManMokMTM/blog-service/internal/model"
	"github.com/RyanTokManMokMTM/blog-service/internal/routers"
	customLog "github.com/RyanTokManMokMTM/blog-service/pkg/logger"
	"github.com/RyanTokManMokMTM/blog-service/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"log"
	"net/http"
	"time"
)

/*
TOOL to use:
Viper - Load setting file
ErrCode - Custom Error - consistent
Logger -  Logger with different info and store at log file
*/

// @title music api server
// @version 1.0
// @description  IOS Music Web Service

// @contact.name jackson.tmm
// @contact.url https://github.com/RyanTokManMokMTM
// @contact.email RyanTokManMokMTM@hotmaiol.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http

func main() {
	//using the config setting variable that have loaded from yaml file
	gin.SetMode(global.ServerSetting.RunMode)
	route := routers.NewRoute()
	//ignore API First
	//if gin.Mode() == gin.DebugMode {
	//	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFile.Handler))
	//}
	server := http.Server{
		Addr:           fmt.Sprintf("127.0.0.1:%s", global.ServerSetting.HttpPort),
		Handler:        route,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatalln(server.ListenAndServe())
}

//used to init process
func init() {
	if err := setUpSetting(); err != nil {
		log.Fatalf("init setting failed :%s", err.Error())
	}

	if err := setUpDBEngine(); err != nil {
		log.Fatalf("init database engine failed %s", err.Error())
	}

	if err := setUpLogger(); err != nil {
		log.Fatalf("init logger failed %s", err.Error())
	}
}

//setUpSetting Set up the global setting variable
func setUpSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {

		return err
	}

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	//Server setting had set the request time out(read and write)
	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	global.AppSetting.ContextTimeOut *= time.Second
	//JWT Expired time
	global.JWTSetting.Expire *= time.Second
	return nil
}

//setUpDBEngine Set up the global DBEngine(gorm)
func setUpDBEngine() error {
	var err error
	//Warning:using := will cause global.DBEngine variable still be nil/nothing
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)

	return err
}

//setupLogger Set up the global DBEngine(gorm)
func setUpLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = customLog.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,  //in mb
		MaxAge:    10,   //oldest day to retain in log file
		LocalTime: true, //UTC time
	}, "", log.LstdFlags).WithCaller(0)

	return nil
}
