package configuration

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	accadapter "nx-go-api/app/account/adapters"
	"nx-go-api/app/account/repositories"
	_ "nx-go-api/app/account/repositories/implementations"
	_ "nx-go-api/app/account/usecases"
	_ "nx-go-api/app/account/usecases/implementations"
	_ "nx-go-api/app/authentication"
	authadapter "nx-go-api/app/authentication/adapters"
	_ "nx-go-api/app/authentication/usecase"
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
	// for testing purpose
	router.Use(CORSMiddleware())

	s := &Server{Router: router}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	//s.loadAuthorizer()
	s.setupRoutes()
	s.initModels()
	//s.setupLogger()
	return s
}

func (s *Server) setupRoutes() {
	accadapter.SetupAccountRoutes(s.Router)
	authadapter.SetupAuthRoutes(s.Router)

	health(s)
}

func (s *Server) initModels() {
	repositories.Migrate()
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
