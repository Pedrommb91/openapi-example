package router

import (
	"github.com/Pedrommb91/openapi-example/internal/api/openapi"
	"github.com/Pedrommb91/openapi-example/internal/api/service"
	"github.com/Pedrommb91/openapi-example/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
)

func NewRouter(handler *gin.Engine, l logger.Interface) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.StaticFile("/swagger", "./spec/openapi.yaml")

	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger", Path: "/swagger/ui"}
	sh := middleware.SwaggerUI(opts, nil)
	handler.GET("/swagger/ui", func(ctx *gin.Context) {
		sh.ServeHTTP(ctx.Writer, ctx.Request)
	})

	mid := make([]openapi.MiddlewareFunc, 0)
	opt := openapi.GinServerOptions{
		BaseURL:     "/api/v1/",
		Middlewares: mid,
	}
	openapi.RegisterHandlersWithOptions(handler, service.NewClient(l), opt)
}
