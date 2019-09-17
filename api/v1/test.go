package v1

import (
	"gin-project-template/pkg/common"
	"gin-project-template/pkg/err"
	"github.com/gin-gonic/gin"
)

func GetTest(c *gin.Context) {
	app := common.Gin{C: c}
	app.Success(err.SUCCESS, nil)
}

func StoreTest(c *gin.Context) {
	app := common.Gin{C: c}
	app.Success(err.SUCCESS, nil)
}

func DelTest(c *gin.Context) {
	app := common.Gin{C: c}
	app.Success(err.SUCCESS, nil)
}

func UpdateTest(c *gin.Context) {
	app := common.Gin{C: c}
	app.Success(err.SUCCESS, nil)
}
