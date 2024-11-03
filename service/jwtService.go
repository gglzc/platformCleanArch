package service

import "github.com/golang-jwt/jwt/v5"


type JWTService interface{
	GenerateToken()string
	ValidateToken()(*jwt.Token ,error)
	
}