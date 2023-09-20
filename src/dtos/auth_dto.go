package dtos

import (
	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	jwt.RegisteredClaims

	UserID       string `json:"user_id"`
}
