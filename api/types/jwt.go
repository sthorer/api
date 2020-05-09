package types

import "github.com/dgrijalva/jwt-go"

type JWTCustomClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}
