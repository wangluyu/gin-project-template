package api

import (
	"gin-project-template/pkg/err"
	"gin-project-template/pkg/util"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	app := util.Gin{C: c}
	app.Success(err.SUCCESS, nil)
}
