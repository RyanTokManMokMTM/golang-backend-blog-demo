package global

import (
	"github.com/RyanTokManMokMTM/blog-service/pkg/logger"
	"github.com/RyanTokManMokMTM/blog-service/pkg/setting"
)

//read config from yaml with viper framework
var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSetting
	Logger          *logger.Logger
)
