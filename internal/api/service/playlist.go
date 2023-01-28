package service

import (
	"net/http"

	"github.com/Pedrommb91/openapi-example/internal/api/error_handler"
	"github.com/Pedrommb91/openapi-example/pkg/errors"
	"github.com/gin-gonic/gin"
)

// GetPlaylist implements openapi.ServerInterface
func (c *client) GetPlaylist(ctx *gin.Context, id int) {
	playlist, err := GetMockPlayListById(id)
	if err != nil {
		errors.LogError(c.log, err)
		error_handler.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, playlist)
}
