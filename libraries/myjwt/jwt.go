package myjwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// key 必须是byte类型
var secret = []byte{123}

// 有效时间
var timeHour time.Duration = 3

// 授权信息
type AuthToken struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 创建token
func GenerateToken(param map[string]string) (token string)  {

	notTime := time.Now()
	expireTime := notTime.Add(timeHour * time.Hour)

	token = ""

	loginToken := AuthToken{
		Username: param["username"],
		Password: param["password"],
		StandardClaims : jwt.StandardClaims{
			ExpiresAt : expireTime.Unix(),
			Issuer : "my_gin",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, loginToken)

	token, err := tokenClaims.SignedString(secret)

	if err != nil {
		fmt.Println("获取token失败")
		fmt.Println(err)
	}

	return token

}

// 解密授权信息
func ParseToken(token string) (*AuthToken, error)  {



	authToken, err := jwt.ParseWithClaims(token, &AuthToken{}, func(token *jwt.Token) (i interface{}, err error) {
		return secret, nil
	})

	if authToken != nil {
		if info, ok := authToken.Claims.(*AuthToken); ok && authToken.Valid {
			return info, nil
		}
	}

	return nil, err

}


