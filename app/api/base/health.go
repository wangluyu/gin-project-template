package base

import (
	"gin-project-template/pkg/e"
	"gin-project-template/pkg/log"
	"gin-project-template/pkg/util"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	app := util.Gin{C: c}
	log.Info["test"] = 1
	log.Info["hello"] = "world"
	log.Info["aa"] = "bb"
	_ = log.WriteInfo()
	_ = log.Error("Error")
	_ = log.Warn("warning")
	app.Success(e.Success, nil)
}
