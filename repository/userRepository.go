package repository

import (
	"context"
	"errors"

	"github.com/gglzc/mqTest/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		CreateUser(ctx context.Context, tx *gorm.DB, user entity.User) error
		GetUser(ctx context.Context, tx *gorm.DB, userID uuid.UUID) (entity.User, error)
		UpdateBalance(ctx context.Context, tx *gorm.DB, user entity.User, money int64) error
	}
	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// AddBalance implements UserRepository.
func (u *userRepository) UpdateBalance(ctx context.Context, tx *gorm.DB, user entity.User, money int64) error {
	// 開啟交易
	return tx.Transaction(func(tx *gorm.DB) error {
		// // 查詢當前用戶的餘額
		userCurrentBalance := user.Balance
		// var currentBalance int64
		if err := tx.Model(&user).Select("balance").Where("id = ?", user.ID).Scan(&userCurrentBalance).Error; err != nil {
			// 查詢失敗則回滾並返回錯誤
			return err
		}

		// 計算更新後的餘額
		newBalance := userCurrentBalance + money
		// 如果新餘額為負數，返回錯誤
		if newBalance < 0 {
			return errors.New("insufficient balance")
		}

		// 更新餘額
		if err := tx.Model(&user).Where("id = ?", user.ID).Update("Balance", newBalance).Error; err != nil {
			// 更新失敗則回滾並返回錯誤
			return err
		}
		// 成功提交交易
		// 將資料丟入message queue中

		return nil
	})
}

// GetUser implements UserRepository.
func (u *userRepository) GetUser(ctx context.Context, tx *gorm.DB, userID uuid.UUID) (entity.User, error) {
	var user entity.User
	err := tx.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, err
		}
		return entity.User{}, err
	}
	return user, nil
}

func (u *userRepository) CreateUser(ctx context.Context, tx *gorm.DB, user entity.User) error {
	err := tx.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		//success
		return nil
	})

	// 檢查是否有錯誤
	if err != nil {
		// 記錄錯誤日誌等
		return err
	}

	return nil
}
