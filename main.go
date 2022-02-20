package main

import (
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
Server Basic Components
1.Consistent Response Standard(一致性的互動規則)
2.
*/

//used to init process
func init() {
	if err := setUpSetting(); err != nil {
		log.Fatalf("init setting failed :%s", err.Error())
	}

	if err := setUpDBEngine(); err != nil {
		log.Fatalf("init database engine failed %s", err.Error())
	}

	if err := setupLogger(); err != nil {
		log.Fatalf("init logger failed %s", err.Error())
	}
}

func main() {
	//using the config setting variable that have loaded from yaml file
	gin.SetMode(global.ServerSetting.RunMode)
	route := routers.NewRoute()
	server := http.Server{
		Addr:           ":8080",
		Handler:        route,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//Testing Config setting is loaded or not
	//fmt.Println(global.ServerSetting)
	//fmt.Println(global.AppSetting)
	//fmt.Println(global.DatabaseSetting)
	//global logger testing
	global.Logger.Infof("%s :github.com/RyanTokManMokMTM/%s", "Testing", "blog-service")
	log.Fatalln(server.ListenAndServe())
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

	//Server setting had set the request time out(read and write)
	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
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
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = customLog.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,  //in mb
		MaxAge:    10,   //oldest day to retain in log file
		LocalTime: true, //UTC time
	}, "", log.LstdFlags).WithCaller(0)

	return nil
}
