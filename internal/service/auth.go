package service

import (
	"errors"
)

type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

//CheckAuth whether key and secret is existed
func (serve *Service) CheckAuth(param *AuthRequest) error {
	auth, err := serve.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}

	//auth info exist
	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist")
}
