package controller

import (
	"net/http"

	"github.com/GustavoCardilho/api-in-golang/src/entities"
	"github.com/gin-gonic/gin"
)

type userController struct {
	users []entities.User
}

func NewController() *userController {
	return &userController{}
}

func (u *userController) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, u.users)
}
