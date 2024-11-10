package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name              string    `json:"name"`
	TelpNumber        string    `json:"telp_number"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	Role              string    `json:"role"`
	ImageURL          string    `json:"image_url"`
	IsVerified        bool      `json:"is_verified"`
	Balance           int64     `json:"balance"`
	CreateDate        time.Time `json:"createDate"`
	OAuthProvider     string    `json:"oauth_provider"`
	OAuthID           string    `json:"oauth_id"`
	OAuthToken        string    `json:"oauth_token"`
	OAuthRefreshToken string    `json:"oauth_refresh_token"`
	OAuthTokenExpiry  time.Time `json:"oauth_token_expiry"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// var err error
	// u.ID = uuid.New()
	// u.Password, err = helpers.HashPassword(u.Password)
	// if err != nil {
	// 	return err
	// }
	return nil
}
