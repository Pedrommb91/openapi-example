package service

import (
	"net/http"
	"time"

	"github.com/Pedrommb91/openapi-example/internal/api/error_handler"
	"github.com/Pedrommb91/openapi-example/internal/api/openapi"
	"github.com/Pedrommb91/openapi-example/pkg/errors"
	"github.com/gin-gonic/gin"
)

// GetPlaylist implements openapi.ServerInterface
func (c *client) GetPlaylist(ctx *gin.Context, id int) {
	playlist, err := GetPlayListById(id)
	if err != nil {
		errors.LogError(c.log, err)
		error_handler.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, playlist)
}

// CreatePlaylist implements openapi.ServerInterface
func (c *client) CreatePlaylist(ctx *gin.Context) {
	var playlist openapi.CreatePlaylistJSONRequestBody
	if err := ctx.ShouldBindJSON(&playlist); err != nil {
		ctx.JSON(http.StatusBadRequest,
			openapi.Error{
				Error:     err.Error(),
				Message:   "Bad Request",
				Path:      ctx.FullPath(),
				Status:    http.StatusBadRequest,
				Timestamp: time.Now(),
			},
		)
		return
	}

	id := AddPlayList(playlist)
	ctx.JSON(http.StatusCreated, openapi.PlaylistResponse{
		Id:       id,
		Playlist: playlist.Playlist,
	})
}
