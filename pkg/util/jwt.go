package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret []byte

type Claims struct {
	ID  string `json:"id"`
	Key string `json:"key"`
	jwt.StandardClaims
}

func NewToken(id string, key string) (string, error) {
	conf, err := FetchAppConf()
	if err != nil {
		return "", err
	}
	jwtConf := conf.Jwt
	expireTime := time.Now().Add(time.Duration(jwtConf.Expire) * time.Hour).Unix()

	claims := Claims{
		id,
		key,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    jwtConf.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
