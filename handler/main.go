package handler

import (
	"architecture-test/registry"
	"architecture-test/usecase"
	"github.com/gin-gonic/gin"
)

func (con *Controller) GetUsers(c *gin.Context, registry registry.Registry) {

	userRepo := registry.NewUserRepository()
	users := usecase.GetUsers(userRepo)
	c.JSON(200, UsersDTO(users))
}
