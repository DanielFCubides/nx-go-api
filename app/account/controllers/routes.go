package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"nx-go-api/app"
)

func SetupAccountRoutes(r *gin.Engine) {
	var controller *Account
	invokeFunc := func(h *Account) {
		controller = h
	}

	err := app.Injector.Invoke(invokeFunc)
	if err != nil {
		logrus.Warn("could not register account controller to the server")
	}
	r.POST("/v1/accounts/", controller.CreateAccount)
	r.PUT("/v1/accounts/:id", controller.EditAccount)
	r.GET("/v1/accounts/:id", controller.GetAccount)
	r.GET("/v1/accounts/", controller.GetAccounts)
}
