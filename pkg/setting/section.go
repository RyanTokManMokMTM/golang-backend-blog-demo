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
		DefaultPageSize      int
		MaxPageSize          int
		LogSavePath          string
		LogFileName          string
		LogFileExt           string
		ContextTimeOut       time.Duration
		UploadSavePath       string
		UploadSavePathURL    string
		UploadImageMaxSize   int
		UploadImageAllowExts []string
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

	JWTSetting struct {
		Secret string
		Issuer string
		Expire time.Duration
	}

	EmailSetting struct {
		Host     string
		Port     int
		Email    string
		Password string
		IsSSL    bool
		From     string
		To       []string
	}
)
