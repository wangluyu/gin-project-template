package main

import (
	"gin-project-template/router"
)

// @title gin-project-template
// @version 1.0
// @description gin-project-template demo
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8080
// @BasePath /api

func main() {
	router := router.InitRouter()
	router.Run(":8000")
}
