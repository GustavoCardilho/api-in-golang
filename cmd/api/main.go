package main

import (
	"github.com/GustavoCardilho/api-in-golang/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	Gin := gin.Default()
	routes.AppRoutes(Gin)
	Gin.Run()
}
