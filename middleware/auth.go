package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWT Struct
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// verify user token

func AuthToken() gin.HandlerFunc {
	// 从环境变量中获取 JWT 密钥
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET is not found")
	}

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "需要 Authorization "})
			return
		}

		// 檢查是否為Bear token
		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization 需要使用 Bearer Token"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, jwt.ErrInvalidKey
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的 Token"})
			return
		}
		ctx.Set("token", authHeader)
		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}
