package v1

import (
	"gin-project-template/pkg/err"
	"gin-project-template/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetTest(c *gin.Context) {
	app := util.Gin{C: c}
	app.Success(err.SUCCESS, nil)
}

func StoreTest(c *gin.Context) {
	app := util.Gin{C: c}
	app.Success(err.SUCCESS, nil)
}

func DelTest(c *gin.Context) {
	app := util.Gin{C: c}
	app.Success(err.SUCCESS, nil)
}

func UpdateTest(c *gin.Context) {
	app := util.Gin{C: c}
	app.Success(err.SUCCESS, nil)
}
