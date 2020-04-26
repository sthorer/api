package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(secret string, userID int, expiration time.Time) (string, error) {
	// Create token and set claims.
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	claims["exp"] = expiration.Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
