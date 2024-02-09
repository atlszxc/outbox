package server

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"ot/pkg/logger"
)

const (
	DEBUG   = "debug"
	RELEASE = "release"
)

const (
	GET    = "GET"
	POST   = "POST"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

type Server struct {
	S      *gin.Engine
	Port   int
	logger *logger.Logger
}

func (s *Server) Start() {
	runStr := fmt.Sprintf(":%d", s.Port)
	err := s.S.Run(runStr)
	if err != nil {
		s.logger.Log(err.Error(), logger.Fatal)
		os.Exit(1)
	}
}

func (s *Server) AddRoute(method string, path string, handler gin.HandlerFunc) error {
	switch method {
	case GET:
		s.S.GET(path, handler)
	case POST:
		s.S.POST(path, handler)
	case PATCH:
		s.S.PATCH(path, handler)
	case DELETE:
		s.S.DELETE(path, handler)
	default:
		{
			s.logger.Log("Can not create route", logger.Fatal)
			return errors.New("Can not create route")
		}
	}

	return nil
}

func NewServer(mode string, port int) *Server {
	gin.SetMode(mode)

	return &Server{
		S:    gin.Default(),
		Port: port,
		logger: &logger.Logger{
			Tag: "Server",
		},
	}
}
