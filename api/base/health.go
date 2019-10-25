package base

import (
	"gin-project-template/pkg/e"
	"gin-project-template/pkg/util"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	app := util.Gin{C: c}
	app.Success(e.Success, nil)
}
