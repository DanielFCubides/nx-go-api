package adapters

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"nx-go-api/app"
	"nx-go-api/app/authentication"
)

func SetupAccountRoutes(r *gin.Engine) {
	var controller *AccountHandler
	invokeFunc := func(h *AccountHandler) {
		controller = h
	}
	middleware := authentication.New()
	err := app.Injector.Invoke(invokeFunc)
	if err != nil {
		logrus.Warn("could not register account controller to the server")
	}
	r.POST("/v1/accounts/", controller.CreateAccount)

	r.PUT("/v1/accounts/:email/", controller.EditAccount)
	r.GET("/v1/accounts/:email/", controller.GetAccount)
	private := r.Group("")
	private.Use(middleware.Authorize())
	private.GET("/v1/accounts/", controller.GetAccounts)
}
