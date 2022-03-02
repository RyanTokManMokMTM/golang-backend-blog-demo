package model

import "github.com/jinzhu/gorm"

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

func (a Auth) Get(db *gorm.DB) (Auth, error) {
	//get the record
	var result Auth
	db = db.Where("app_key=? AND app_secret = ? AND is_del = ?", a.AppKey, a.AppSecret, 0).First(&result)
	err := db.Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return result, err
	}
	return result, nil
}
