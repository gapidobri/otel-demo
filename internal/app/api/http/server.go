package http

import (
	"github.com/gapidobri/otel-demo/internal/app/service"
	"github.com/gin-gonic/gin"
)

type (
	Server struct {
		service service.Service
	}
)

func NewServer(service service.Service) Server {
	return Server{
		service: service,
	}
}

func (s Server) Run() {
	r := gin.Default()
	registerRoutes(r, s.service)
	r.Run()
}
