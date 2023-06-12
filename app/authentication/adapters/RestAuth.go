package adapters

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nx-go-api/app"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func init() {
	err := app.Injector.Provide(NewAuthHandler)
	if err != nil {
		fmt.Println("Error providing AuthHandler Controller:", err)
		panic(err)
	}
}

func (h AuthHandler) Authenticate(c *gin.Context) {
	var request AuthRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, AuthResponse{})
}
