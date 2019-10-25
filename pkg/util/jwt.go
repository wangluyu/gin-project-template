package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret []byte

type Claims struct {
	OpenId     string `json:"open_id"`
	SessionKey string `json:"session_key"`
	jwt.StandardClaims
}

func NewToken(openId string, sessionKey string) (string, error) {
	expireTime := time.Now().Add(3 * time.Hour).Unix()

	claims := Claims{
		openId,
		sessionKey,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "gin-project-template",
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
