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

func ValidateToken(tokenString string) bool {
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("s3cr3t"), nil
	})
	if err != nil {
		fmt.Println("Couldn't parse token:", err)
		return false
	}
	if token.Valid {
		fmt.Println("You look nice today")
		return true
	} else {
		fmt.Println("Invalid Token:", err)
		return false
	}
}
