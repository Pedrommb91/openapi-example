package service

import (
	"net/http"

	"github.com/Pedrommb91/openapi-example/internal/api/openapi"
	"github.com/gin-gonic/gin"
)

// GetGenres implements openapi.ServerInterface
func (*client) GetGenres(c *gin.Context) {
	c.JSON(http.StatusOK, openapi.GenreListResponse{
		Genres: &[]struct {
			Genre string `json:"genre"`
		}{
			{Genre: "Rock"},
			{Genre: "Pop"},
			{Genre: "Hip-pop"},
			{Genre: "Blues"},
			{Genre: "Soul"},
			{Genre: "Reggae"},
			{Genre: "Country"},
			{Genre: "Funk"},
			{Genre: "Folk"},
			{Genre: "Middle Eastern"},
			{Genre: "Jazz"},
			{Genre: "Disco"},
			{Genre: "Classical"},
			{Genre: "Electronic"},
		},
	})
}
