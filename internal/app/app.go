package app

import (
	"github.com/gapidobri/otel-demo/internal/app/api/http"
	"github.com/gapidobri/otel-demo/internal/app/service"
)

func Run() {
	service := service.NewService()

	server := http.NewServer(service)
	server.Run()
}
