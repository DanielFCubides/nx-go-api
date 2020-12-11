package infrastructure

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"nx-go-api/app"
	"nx-go-api/app/account/controllers"
	"os"
)

// Server is charge to handle the api operations
type Server struct {
	router *gin.Engine
	//authorizer *middleware.Authorizer
	loggerFile *os.File
	logger     *log.Logger
}

// NewServer retrieves a pointer to Server.
func NewServer() *Server {
	router := gin.Default()

	s := &Server{router: router}
	//s.loadAuthorizer()
	s.setupRoutes()
	//s.initModels()
	//s.setupLogger()
	return s
}

func (s *Server) setupRoutes() {
	setupAccountRoutes(s)
	health(s)
}

//func (s *Server) initModels() {
//	accountInitializer.Migrate()
//}

// Run starts the server
func (s *Server) Run() {
	port := os.Getenv("SERVER_PORT")
	log.Infof("server listening on : %s", port)
	_ = s.router.Run(fmt.Sprintf(":%s", port))
}

func (s *Server) Close() {
	s.Close()
}

func setupAccountRoutes(s *Server) {
	var controller *controllers.Public
	invokeFunc := func(h *controllers.Public) {
		controller = h
	}

	err := app.Injector.Invoke(invokeFunc)
	if err == nil {
		log.Warn("could not register account controller to the server")
	}
	s.router.POST("/v1/accounts/", controller.CreateAccount)
}

func health(s *Server) {
	s.router.GET("/health/", func(c *gin.Context) {
		c.String(http.StatusOK, "api up \xF0\x9F\x92\x9A")
	})
}

//func (s *Server) loadAuthorizer() {
//	var authorizer *middleware.Authorizer
//	invokeFunc := func(a *middleware.Authorizer) {
//		authorizer = a
//	}
//	err := core.Injector.Invoke(invokeFunc)
//
//	if err != nil {
//		panic(err)
//	}
//
//	s.authorizer = authorizer
//}
//
//func (s *Server) setupLogger() {
//	setupLogger(s)
//}
