package service

import (
	"fmt"

	"github.com/Pedrommb91/openapi-example/internal/api/openapi"
	"github.com/Pedrommb91/openapi-example/pkg/errors"
	"github.com/rs/zerolog"
)

type PlayList struct {
	id       int
	playlist openapi.PlaylistResponse
}

func GetMockPlaylists() []openapi.PlaylistResponse {
	return []openapi.PlaylistResponse{
		{
			Id: 1,
			Playlist: []openapi.Music{
				{
					Album:  "Master of Puppets",
					Artist: "Metallica",
					Genre:  "Metal",
					Name:   "Master of Puppets",
				},
				{
					Album:  "Metallica",
					Artist: "Metallica",
					Genre:  "Metal",
					Name:   "Enter Sandman",
				},
				{
					Album:  "The Living Infinite",
					Artist: "Soilwrk",
					Genre:  "Metal",
					Name:   "This Momentary Bliss",
				},
				{
					Album:  "Doomsday",
					Artist: "Architects",
					Genre:  "Metal",
					Name:   "Holly Hell",
				},
			},
		},
		{
			Id: 2,
			Playlist: []openapi.Music{
				{
					Album:  "Blizzard of Ozz",
					Artist: "Ozzy Osbourne",
					Genre:  "Metal",
					Name:   "Crazy Train",
				},
				{
					Album:  "From the fires",
					Artist: "Greta van fleet",
					Genre:  "Rock",
					Name:   "Black Smoke Rising",
				},
			},
		},
	}
}

func GetMockPlayListById(id int) (openapi.PlaylistResponse, error) {
	const op = "service.GetMockPlayListById"
	for _, playlist := range GetMockPlaylists() {
		if playlist.Id == id {
			return playlist, nil
		}
	}
	return openapi.PlaylistResponse{}, errors.Build(
		errors.KindNoContent(),
		errors.WithOp(op),
		errors.WithError(fmt.Errorf("Playlist not found")),
		errors.WithSeverity(zerolog.WarnLevel),
	)
}
