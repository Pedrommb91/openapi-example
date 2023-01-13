package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	server *http.Server
}

func NewServer() *Server {
	handler := gin.New()
	return &Server{
		engine: handler,
		server: &http.Server{
			Addr:              ":8080",
			Handler:           handler,
			ReadHeaderTimeout: time.Second * 30,
		},
	}
}
