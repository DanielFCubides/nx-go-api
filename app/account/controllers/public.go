package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nx-go-api/app"
)

type Public struct {
}

func NewPublic() *Public {
	return &Public{}
}

func init() {
	err := app.Injector.Provide(NewPublic)
	if err != nil {
		fmt.Println("Error providing Account Public Controller:", err)
		panic(err)
	}
}

type AccountRequest struct {
}

func (h *Public) CreateAccount(c *gin.Context) {
	var request AccountRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)

}

func (h *Public) EditAccount(c *gin.Context) {

}

func (h *Public) ChangeStatusAccount(c *gin.Context) {

}

func (h *Public) GetAccount(c *gin.Context) {

}
