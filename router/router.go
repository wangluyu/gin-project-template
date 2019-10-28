package router

import (
	"gin-project-template/app/api/base"
	app "gin-project-template/app/api/v1"
	_ "gin-project-template/docs"
	"gin-project-template/middleware"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("health", base.Health)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/v1")
	v1.POST("/login", app.Login)
	v1.Use(middleware.Jwt())
	{
		v1.GET("/example", app.GetExample)
		v1.POST("/example", app.StoreExample)
		v1.DELETE("/example", app.DelExample)
		v1.PUT("/example", app.UpdateExample)
	}

	return router
}
