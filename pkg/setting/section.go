package setting

import "time"

type (
	ServerSettings struct {
		RunMode      string
		HttpPort     string
		ReadTimeOut  time.Duration
		WriteTimeOut time.Duration
	}

	AppSettings struct {
		DefaultPageSize int
		MaxPageSize     int
		LogSavePath     string
		LogFileName     string
		LogFileExt      string
	}

	DatabaseSetting struct {
		DBType       string
		User         string
		Password     string
		Host         string
		DBName       string
		TablePrefix  string
		Charset      string
		ParseTime    bool
		MaxIdleConns int
		MaxOpenConns int
	}
)
