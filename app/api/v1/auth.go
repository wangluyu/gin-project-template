package v1

import (
	"gin-project-template/pkg/e"
	"gin-project-template/pkg/util"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	ID  string `form:"id" json:"id" binding:"required"`
	Key string `form:"key" json:"key" binding:"required"`
}

func Login(c *gin.Context) {
	app := util.Gin{C: c}
	var auth Auth
	if c.ShouldBind(&auth) == nil {
		app.UnAuth(e.ErrorAuthParams, nil)
		return
	}
	token, err := util.NewToken(auth.ID, auth.Key)
	if err != nil {
		app.Error(e.Error, nil)
		return
	}
	app.Success(e.Success, token)
}
