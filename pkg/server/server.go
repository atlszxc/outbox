package server

import (
	"errors"
	"github.com/gin-gonic/gin"
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
	S *gin.Engine
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
		return errors.New("Can not create route")
	}

	return nil
}

func NewServer(mode string) *Server {
	gin.SetMode(mode)

	return &Server{
		S: gin.Default(),
	}
}
