package router

import (
	"gin-project-template/api"
	apiV1 "gin-project-template/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("health", api.Health)

	v1 := router.Group("/v1")
	v1.Use(jwt.JWT())
	{
		v1.GET("/test", apiV1.GetTest)
		v1.POST("/test", apiV1.StoreTest)
		v1.DELETE("/test", apiV1.DelTest)
		v1.PUT("/test", apiV1.UpdateTest)
	}

	return router
}
