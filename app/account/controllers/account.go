package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nx-go-api/app"
)

type Account struct {
}

func NewPublic() *Account {
	return &Account{}
}

func init() {
	err := app.Injector.Provide(NewPublic)
	if err != nil {
		fmt.Println("Error providing Account Account Controller:", err)
		panic(err)
	}
}

type AccountRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Account) CreateAccount(c *gin.Context) {
	var request AccountRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	//c.Writer.WriteHeader(http.StatusOK)
	c.JSON(http.StatusOK, request)
}

func (h *Account) EditAccount(c *gin.Context) {
	id := c.Param("id")
	var request AccountRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	request.Email = id
	c.JSON(http.StatusOK, request)
}

func (h *Account) GetAccount(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, id)
}

func (h *Account) GetAccounts(c *gin.Context) {
	c.String(http.StatusOK, "looking all ")
}
