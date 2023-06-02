package adapters

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"net/http"
	"nx-go-api/app"
	"nx-go-api/app/account"
	"nx-go-api/app/account/usecases"
	"time"
)

type AccountHandler struct {
	UseCase usecases.AccountUseCase
}

func NewPublic(useCase usecases.AccountUseCase) *AccountHandler {
	return &AccountHandler{
		UseCase: useCase,
	}
}

func init() {
	err := app.Injector.Provide(NewPublic)
	if err != nil {
		fmt.Println("Error providing AccountHandler AccountHandler Controller:", err)
		panic(err)
	}
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var request AccountRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	accountDomain := toDomain(request)
	accountResponse := toResponse(h.UseCase.Create(accountDomain))
	c.JSON(http.StatusOK, accountResponse)
}

func (h *AccountHandler) EditAccount(c *gin.Context) {
	email := c.Param("email")
	var request AccountRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	request.Email = email
	accountDomain := toDomain(request)
	accountResponse := toResponse(h.UseCase.Edit(accountDomain))
	c.JSON(http.StatusOK, accountResponse)
}

func (h *AccountHandler) GetAccount(c *gin.Context) {
	email := c.Param("email")
	accountResponse := toResponse(h.UseCase.FindByEmail(email))
	c.JSON(http.StatusOK, accountResponse)
}

func (h *AccountHandler) GetAccounts(c *gin.Context) {
	accounts := funk.Map(h.UseCase.FindAll(), toResponse)
	c.JSON(http.StatusOK, accounts)

}

func toRequest(a account.Account) AccountRequest {
	return AccountRequest{
		Username: a.Username,
		Email:    a.Email,
		Password: a.Password,
	}
}

func toResponse(a account.Account) AccountResponse {
	return AccountResponse{
		Username: a.Username,
		Email:    a.Email,
	}
}

func toDomain(request AccountRequest) account.Account {
	return account.Account{
		Email:        request.Email,
		Password:     request.Password,
		Status:       "active",
		Username:     request.Username,
		CreationDate: time.Now(),
	}
}
