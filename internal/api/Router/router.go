package router

import (
	"github.com/Pedrommb91/openapi-example/internal/api/openapi"
	"github.com/Pedrommb91/openapi-example/internal/api/service"
	"github.com/Pedrommb91/openapi-example/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, l logger.Interface) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	mid := make([]openapi.MiddlewareFunc, 0)
	opt := openapi.GinServerOptions{
		BaseURL:     "/api/v1/",
		Middlewares: mid,
	}
	openapi.RegisterHandlersWithOptions(handler, service.NewClient(l), opt)
}
