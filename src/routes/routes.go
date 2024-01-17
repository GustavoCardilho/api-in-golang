package routes

import (
	"github.com/GustavoCardilho/api-in-golang/src/controller"
	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	userController := controller.NewController()
	v1 := router.Group("/v1")
	{
		v1.GET("/users", userController.GetUsers)
	}

	return v1
}
