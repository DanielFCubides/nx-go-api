package infrastructure

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"nx-go-api/app/account/controllers"
	"nx-go-api/app/account/gateway"
	_ "nx-go-api/app/account/gateway/implementations"
	_ "nx-go-api/app/account/usecases"
	_ "nx-go-api/app/account/usecases/implementations"
	"os"
)

// Server is charge to handle the api operations
type Server struct {
	Router *gin.Engine
	//authorizer *middleware.Authorizer
	loggerFile *os.File
	logger     *log.Logger
}

// NewServer retrieves a pointer to Server.
func NewServer() *Server {
	router := gin.Default()

	s := &Server{Router: router}
	//s.loadAuthorizer()
	s.setupRoutes()
	s.initModels()
	//s.setupLogger()
	return s
}

func (s *Server) setupRoutes() {
	controllers.SetupAccountRoutes(s.Router)
	health(s)
}

func (s *Server) initModels() {
	gateway.Migrate()
}

// Run starts the server
func (s *Server) Run() {
	port := os.Getenv("SERVER_PORT")
	log.Infof("server listening on : %s", port)
	_ = s.Router.Run(fmt.Sprintf(":%s", port))
}

func (s *Server) Close() {
	s.Close()
}

func health(s *Server) {
	s.Router.GET("/health/", func(c *gin.Context) {
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
