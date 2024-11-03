package controller

import (
	"github.com/gglzc/mqTest/service"
	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		GetUser(ctx *gin.Context)
		UpdateBalance(ctx *gin.Context)
	}

	userController struct {
		userService service.UserService
	}
)


func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

// GetUser implements UserController.
func (u *userController) GetUser(ctx *gin.Context) {
	panic("unimplemented")
}

// UpdateBalance implements UserController.
func (u *userController) UpdateBalance(ctx *gin.Context) {
	panic("unimplemented")
}

