package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"nx-go-api/app"
	"nx-go-api/app/authentication/adapters"
)

func SetupAuthRoutes(r *gin.Engine) {
	var controller *adapters.AuthHandler
	invokeFunc := func(h *adapters.AuthHandler) {
		controller = h
	}

	err := app.Injector.Invoke(invokeFunc)
	if err != nil {
		logrus.Warn("could not register account controller to the server")
	}
	r.POST("/v1/auth/", controller.Authenticate)

}
