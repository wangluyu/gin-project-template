package v1

import (
	"gin-project-template/pkg/e"
	"gin-project-template/pkg/util"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID      string `json:"id" example:"10001"`
	Name    string `json:"name" example:"user name"`
	Address string `json:"address" example:"user address"`
}

// @Summary 获取数据
// @Description
// @Tags V1
// @Produce  json
// @Param id query string true "User ID"
// @Success 200 {object} swagger.GetExample
// @Failure 400 {object} swagger.Failure
// @Router /v1/example/{id} [get]
func GetExample(c *gin.Context) {
	app := util.Gin{C: c}
	app.Success(e.Success, nil)
}

func StoreExample(c *gin.Context) {
	app := util.Gin{C: c}
	app.Success(e.Success, nil)
}

func DelExample(c *gin.Context) {
	app := util.Gin{C: c}
	app.Success(e.Success, nil)
}

func UpdateExample(c *gin.Context) {
	app := util.Gin{C: c}
	app.Success(e.Success, nil)
}
