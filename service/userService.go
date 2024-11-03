package service

import (
	"context"

	"github.com/gglzc/mqTest/consts"
	"github.com/gglzc/mqTest/entity"
	"github.com/gglzc/mqTest/model/request"
	"github.com/gglzc/mqTest/model/response"
	"github.com/gglzc/mqTest/repository"
)

type (
	UserService interface {
		GetUser(ctx context.Context, req request.CreateUserRequest) (response.CommonResponse , error)
		UpdateBalance(ctx context.Context, req request.UpdateBalance) (response.CommonResponse , error)
	}

	userService struct {
		userRepo   repository.UserRepository
		jwtService JWTService
	}
)

func NewUserService(userRepo repository.UserRepository, jwtService JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}


// GetUser implements UserService.
func (u *userService) GetUser(ctx context.Context, req request.CreateUserRequest) (response.CommonResponse , error) {
	_ , err:= u.userRepo.GetUser(ctx,nil,req.UserID)
	// fail 
	if(err!=nil){
		return response.CommonResponse{
			Code: consts.NotFoundCode,
			Message: err.Error(),
		},err
	}
	// 寫入 redis 紀錄
	// 丟入message queue

	return response.CommonResponse{
		Code: consts.SuccessCode,
		Message: "Success",
	},nil
}

// UpdateBalance implements UserService.
func (u *userService) UpdateBalance(ctx context.Context, req request.UpdateBalance)(response.CommonResponse , error) {
	err :=u.userRepo.UpdateBalance(ctx , nil , entity.User{},req.IntoMoney)
	if(err!=nil){
		return response.CommonResponse{
			Code: consts.BalanceNotEnough,
			Message: err.Error(),
		},err
	}

	// 丟入message queue

	return response.CommonResponse{
		Code: consts.SuccessCode,
		Message: "Success",
	},nil
}

