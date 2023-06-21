package adapters

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"nx-go-api/app"
	"nx-go-api/app/account"
	"nx-go-api/app/authentication/usecase"
	"time"
)

type AuthHandler struct {
	Usecase usecase.AuthUseCase
}

func NewAuthHandler(usecase usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		Usecase: usecase,
	}
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
	if !h.Usecase.Authenticate(request.Email, request.Password) {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := h.Usecase.GenerateToken(request.Email, account.StatusActive)
	if err != nil {
		log.Errorf("{\"error\": \"%s\" }", err.Error())
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	unix := time.Now().Add(time.Minute * 10).Unix()
	log.Debugf("{\"token\": \"%s\" , \"exp\": \"%d\"}", token, unix)
	c.JSON(http.StatusOK, AuthResponse{Token: token, Expiration: fmt.Sprintf("%d", unix)})
}
