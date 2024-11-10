package controller

import (
	"net/http"

	"github.com/gglzc/mqTest/entity"
	"github.com/gglzc/mqTest/model/request"
	"github.com/gglzc/mqTest/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	UserController interface {
		GetUser(ctx *gin.Context)
		CreateUser(ctx *gin.Context)
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
	userIDParms:= ctx.Param("userID")
	userID, err := uuid.Parse(userIDParms)
    // if err != nil {
    //     ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID Not Found"})
    //     return
    // }

	user ,err := u.userService.GetUser(ctx.Request.Context(),request.CreateUserRequest{
		UserID: userID,
	})

	if err!=nil{
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	// request.CreateUserRequest{

	// })
	ctx.JSON(http.StatusOK, user)
}

// UpdateBalance implements UserController.
func (u *userController) UpdateBalance(ctx *gin.Context) {
	panic("unimplemented")
}

func (u *userController) CreateUser(ctx *gin.Context) {
	var req entity.User

	if err:=ctx.ShouldBindBodyWithJSON(&req); err != nil {
		// 如果解析失敗返回錯誤響應
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}
	resp ,err := u.userService.CreateUser(ctx,req)
	if err != nil {
		// 如果服務層返回錯誤，則回傳 500 錯誤
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
			"error":   err.Error(),
		})
		return
	}
	// 成功時返回 201 狀態碼及響應內容
	ctx.JSON(http.StatusCreated, resp)
}

