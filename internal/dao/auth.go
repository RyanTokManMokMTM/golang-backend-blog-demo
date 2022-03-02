package dao

import "github.com/RyanTokManMokMTM/blog-service/internal/model"

//GetAuth get auth data
func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
