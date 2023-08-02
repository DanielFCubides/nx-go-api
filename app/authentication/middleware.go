package authentication

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nx-go-api/app"
	"nx-go-api/app/authentication/token"
	"strings"
)

type AuthMiddleware struct {
}

func New() *AuthMiddleware {

	return &AuthMiddleware{}
}

func init() {
	err := app.Injector.Provide(New)
	if err != nil {
		fmt.Println("Error providing AuthMiddleware:", err)
		panic(err)
	}
}

func (m AuthMiddleware) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !ValidateToken(c.Request.Header) {
			c.JSON(http.StatusUnauthorized, "")
			c.Abort()
		}
	}
}

func ValidateToken(authHeaders http.Header) bool {
	authHeader, exist := authHeaders["Authorization"]
	if !exist {
		return false
	}
	bearerToken := strings.Split(authHeader[0], " ")
	if len(bearerToken) != 2 {
		return false
	}
	if !token.ValidateToken(bearerToken[1]) {
		return false
	}
	return true
}
