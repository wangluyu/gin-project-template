package main

import (
	"gin-project-template/router"
)

func main() {
	router := router.InitRouter()
	router.Run(":8000")
}
