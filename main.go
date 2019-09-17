package main

import (
	"github.com/wangluyu/gin-project-template/router"
)

func main() {
	router := router.InitRouter()
	router.Run(":8000")
}
