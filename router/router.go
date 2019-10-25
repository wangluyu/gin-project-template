package router

import (
	"gin-project-template/api/base"
	api "gin-project-template/api/v1"
	"gin-project-template/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("health", base.Health)

	v1 := router.Group("/v1")
	v1.POST("/login", api.Login)
	v1.Use(middleware.Jwt())
	{
		v1.GET("/test", api.GetTest)
		v1.POST("/test", api.StoreTest)
		v1.DELETE("/test", api.DelTest)
		v1.PUT("/test", api.UpdateTest)
	}

	return router
}
