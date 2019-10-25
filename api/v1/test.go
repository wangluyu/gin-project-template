package v1

import (
	"gin-project-template/pkg/e"
	"gin-project-template/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetTest(c *gin.Context) {
	app := util.Gin{C: c}
	app.Success(e.Success, nil)
}

func StoreTest(c *gin.Context) {
	app := util.Gin{C: c}
	app.Success(e.Success, nil)
}

func DelTest(c *gin.Context) {
	app := util.Gin{C: c}
	app.Success(e.Success, nil)
}

func UpdateTest(c *gin.Context) {
	app := util.Gin{C: c}
	app.Success(e.Success, nil)
}
