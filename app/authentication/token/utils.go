package token

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateToken(email string, status string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email":  email,
		"Status": status,
		"iat":    time.Now().Unix(),
	})

	tokenSigned, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	fmt.Print(tokenSigned)
	return tokenSigned, nil
}
