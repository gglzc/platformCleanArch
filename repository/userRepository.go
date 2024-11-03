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
		GetUser(ctx context.Context, tx *gorm.DB, userID uuid.UUID) (entity.User, error)
		UpdateBalance(ctx context.Context, tx *gorm.DB, user entity.User , money int64) error
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
	panic("unimplemented")
}




