package middleware

import (
	"gin-project-template/pkg/e"
	"gin-project-template/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.Success
		token := c.Request.Header.Get("Token")
		if token == "" {
			code = e.ErrorAuth
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ErrorAuthTokenExpired
				default:
					code = e.ErrorAuthToken
				}
			}
		}

		if code != e.Success {
			app := util.Gin{C: c}
			app.UnAuth(code, data)
			c.Abort()
			return
		}

		c.Next()
	}
}
