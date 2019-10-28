package v1

import (
	"gin-project-template/pkg/e"
	"gin-project-template/pkg/util"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	OpenId     string `form:"open_id" json:"open_id" binding:"required"`
	SessionKey string `form:"session_key" json:"session_key" binding:"required"`
}

func Login(c *gin.Context) {
	app := util.Gin{C: c}
	var auth Auth
	if c.ShouldBind(&auth) == nil {
		app.UnAuth(e.ErrorAuthParams, nil)
		return
	}
	token, err := util.NewToken(auth.OpenId, auth.SessionKey)
	if err != nil {
		app.Error(e.Error, nil)
		return
	}
	app.Success(e.Success, token)
}
