package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gglzc/mqTest/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MockGormDB struct {
	mock.Mock
	*gorm.DB
}
var mofck sqlmock.Sqlmock

func (m *MockGormDB) Transaction(fc func(tx *gorm.DB) error) error {
	args := m.Called(fc)
	return args.Error(0)
}

func (m *MockGormDB) WithContext(ctx context.Context) *gorm.DB {
	m.Called(ctx)
	return m.DB
}

// Create 模擬 gorm.DB 的 Create 方法
func (m *MockGormDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	db := &gorm.DB{}
	if err, ok := args.Get(0).(error); ok && err != nil {
		db.Error = err
	}
	return db
}

func TestCreateUser(t *testing.T) {
	// 设置 sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("无法打开 sqlmock 数据库: %v", err)
	}
	defer db.Close()

	// 初始化 GORM 数据库
	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("无法初始化 GORM 数据库: %v", err)
	}

	// 创建 userRepository 实例
	repo := &userRepository{db: gormDB}

	ctx := context.TODO()
	user := entity.User{
		ID:                [16]byte{},
		Name:              "Richar",
		TelpNumber:        "094",
		Email:             "qw222qwyz@gmail.com",
		Password:          "dq",
		Role:              "admin",
		ImageURL:          "img:001",
		IsVerified:        false,
		Balance:           0,
		CreateDate:        time.Now(),
		OAuthProvider:     "nil",
		OAuthID:           "0",
		OAuthToken:        "h",
		OAuthRefreshToken: "w",
		OAuthTokenExpiry:  time.Now(),
	}

	// 成功情况
	t.Run("success", func(t *testing.T) {
		// 设置期望的 SQL 语句和结果
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users"`)).
			WithArgs(sqlmock.AnyArg(), // ID
				user.Name,
				user.TelpNumber,
				user.Email,
				user.Password,
				user.Role,
				user.ImageURL,
				user.IsVerified,
				user.Balance,
				sqlmock.AnyArg(), // CreateDate
				user.OAuthProvider,
				user.OAuthID,
				user.OAuthToken,
				user.OAuthRefreshToken,
				sqlmock.AnyArg(), // OAuthTokenExpiry
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// 调用被测试的函数
		err := repo.CreateUser(ctx, gormDB,user)
		assert.NoError(t, err)

		// 确认所有期望的操作都被执行
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("未满足的期望: %s", err)
		}
	})

	// 失败情况：Create 方法返回错误
	t.Run("create error", func(t *testing.T) {
		// 设置期望的 SQL 语句和错误
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users"`)).
			WithArgs(sqlmock.AnyArg(), // ID
				user.Name,
				user.TelpNumber,
				user.Email,
				user.Password,
				user.Role,
				user.ImageURL,
				user.IsVerified,
				user.Balance,
				sqlmock.AnyArg(), // CreateDate
				user.OAuthProvider,
				user.OAuthID,
				user.OAuthToken,
				user.OAuthRefreshToken,
				sqlmock.AnyArg(), // OAuthTokenExpiry
			).
			WillReturnError(errors.New("create error"))
		mock.ExpectRollback()

		// 调用被测试的函数
		err := repo.CreateUser(ctx, gormDB,user)
		assert.Error(t, err)
		assert.EqualError(t, err, "create error")

		// 确认所有期望的操作都被执行
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("未满足的期望: %s", err)
		}
	})
}
