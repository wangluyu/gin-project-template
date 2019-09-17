package api

import (
	"gin-project-template/pkg/common"
	"gin-project-template/pkg/err"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	app := common.Gin{C: c}
	app.Success(err.SUCCESS, nil)
}
