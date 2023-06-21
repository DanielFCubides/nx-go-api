package adapters

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"nx-go-api/app"
)

func SetupAuthRoutes(r *gin.Engine) {
	var controller *AuthHandler
	invokeFunc := func(h *AuthHandler) {
		controller = h
	}

	err := app.Injector.Invoke(invokeFunc)
	if err != nil {
		logrus.Warn("could not register Auth handler to the server")
	}
	r.POST("/v1/auth/", controller.Authenticate)

}
