package app

import (
	"github.com/Pedrommb91/openapi-example/internal/api"
	"github.com/Pedrommb91/openapi-example/pkg/logger"
)

func Run() {
	l := logger.New("info")

	server := api.NewServer(*l)
	server.CreateServer()
	server.SetRoutes()
	server.Run()
}
