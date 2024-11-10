package service

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gglzc/mqTest/consts"
	"github.com/gglzc/mqTest/entity"
	"github.com/gglzc/mqTest/model/request"
	"github.com/gglzc/mqTest/model/response"
	"github.com/gglzc/mqTest/repository"
)

type (
	UserService interface {
		GetUser(ctx context.Context, req request.CreateUserRequest) (response.CommonResponse, error)
		CreateUser(ctx context.Context, req entity.User) (response.CommonResponse, error)
		UpdateBalance(ctx context.Context, req request.UpdateBalance) (response.CommonResponse, error)
	}

	userService struct {
		userRepo repository.UserRepository
		// jwtService JWTService
		producer *kafka.Producer
	}
)

func NewUserService(userRepo repository.UserRepository, producer *kafka.Producer) UserService {
	return &userService{
		userRepo: userRepo,
		producer: producer,
	}
}

// GetUser implements UserService.
func (u *userService) GetUser(ctx context.Context, req request.CreateUserRequest) (response.CommonResponse, error) {
	_, err := u.userRepo.GetUser(ctx, nil, req.UserID)
	// fail
	if err != nil {
		return response.CommonResponse{
			Code:    consts.NotFoundCode,
			Message: err.Error(),
		}, err
	}

	// // 寫入 redis 紀錄
	// // 丟入message queue
	// topic:="user_topic"
	// messages := []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"}
	
	// for _, word := range messages {
	// 	// 使用 Produce 傳送訊息
	// 	err := u.producer.Produce(&kafka.Message{
	// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	// 		Value:          []byte(word),
	// 	}, nil)
		
	// 	if err != nil {
	// 		return response.CommonResponse{
	// 			Code:    consts.InternalServerErrorCode,
	// 			Message: err.Error(),
	// 		}, err
	// 	}
	// }

	return response.CommonResponse{
		Code:    consts.SuccessCode,
		Message: "Success",
	}, nil
}

// UpdateBalance implements UserService.
func (u *userService) UpdateBalance(ctx context.Context, req request.UpdateBalance) (response.CommonResponse, error) {
	err := u.userRepo.UpdateBalance(ctx, nil, entity.User{}, req.IntoMoney)
	if err != nil {
		return response.CommonResponse{
			Code:    consts.BalanceNotEnough,
			Message: err.Error(),
		}, err
	}

	// 丟入message queue

	return response.CommonResponse{
		Code:    consts.SuccessCode,
		Message: "Success",
	}, nil
}

func (u *userService) CreateUser(ctx context.Context, req entity.User) (response.CommonResponse, error) {
	err := u.userRepo.CreateUser(ctx, nil, req)

	if err != nil {
		return response.CommonResponse{
			Code:    consts.NoContentCode,
			Message: err.Error(),
		}, err
	}
	return response.CommonResponse{
		Code:    consts.CreatedCode,
		Message: "Success Create User",
	}, nil
}
