package service

import (
	"github.com/Pedrommb91/openapi-example/internal/api/openapi"
	"github.com/Pedrommb91/openapi-example/pkg/logger"
)

type client struct {
	log logger.Interface
}

func NewClient(l logger.Interface) openapi.ServerInterface {
	NewPlayListInMemoryDatabase()

	return &client{
		log: l,
	}
}
