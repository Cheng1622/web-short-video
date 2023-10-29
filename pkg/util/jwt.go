package util

import (
	"errors"
	"time"

	"github.com/Cheng1622/web-short-video/config"
	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	UserId uint
	Email  string
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var jwtSerect = []byte(config.JwtSecret)

// GenToken 生成token
func GenToken(email string, userid uint) (string, error) {
	c := JwtClaims{
		Email:  email,
		UserId: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).UnixNano(), //过期时间
			Issuer:    "blog",                                         //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(jwtSerect)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*JwtClaims, error) {

	var mc = new(JwtClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return jwtSerect, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}

	return nil, errors.New("invalid token")

}
