package app

import (
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/RyanTokManMokMTM/blog-service/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	AppKey    string
	AppSecret string
	jwt.StandardClaims
}

//GetJWTSecret return our secret key
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
	now := time.Now()
	expiredTime := now.Add(global.JWTSetting.Expire)
	claim := Claims{
		AppKey:    util.EnCodeMD5(appKey),
		AppSecret: util.EnCodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tokenWithClaims.SignedString(GetJWTSecret())
	return token, err
}

//ParseToken Parsing the token and get the info
func ParseToken(token string) (*Claims, error) {
	claimInfo, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		//return key of token
		return GetJWTSecret(), nil
	})

	if claimInfo != nil { //parsed
		//get token
		claims, ok := claimInfo.Claims.(Claims)
		if ok && claimInfo.Valid { //cast to custom claims object and check token is valid depend on time. if claims not set any info,it remains valid
			return &claims, nil //return the claim object
		}
	}

	return nil, err
}
