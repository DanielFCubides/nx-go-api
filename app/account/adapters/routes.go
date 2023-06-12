package adapters

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"nx-go-api/app"
)

func SetupAccountRoutes(r *gin.Engine) {
	var controller *AccountHandler
	invokeFunc := func(h *AccountHandler) {
		controller = h
	}

	err := app.Injector.Invoke(invokeFunc)
	if err != nil {
		logrus.Warn("could not register account controller to the server")
	}
	r.POST("/v1/accounts/", controller.CreateAccount)
	private := r.Group("")
	//private.Use()
	private.PUT("/v1/accounts/:email", controller.EditAccount)
	private.GET("/v1/accounts/:email", controller.GetAccount)
	private.GET("/v1/accounts/", controller.GetAccounts)
}
