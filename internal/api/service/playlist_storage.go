package service

import (
	"fmt"

	"github.com/Pedrommb91/openapi-example/internal/api/openapi"
	"github.com/Pedrommb91/openapi-example/pkg/errors"
	"github.com/rs/zerolog"
)

var PlayListsInMemory []PlayListStorage
var lastId int

type PlayListStorage struct {
	id       int
	playlist []openapi.Music
}

func NewPlayListInMemoryDatabase() {
	mockPlaylist := GetMockPlaylists()
	for i, playlist := range mockPlaylist {
		PlayListsInMemory = append(PlayListsInMemory, PlayListStorage{
			id:       i + 1,
			playlist: playlist.Playlist,
		})
		lastId = i + 1
	}
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
					Artist: "Soilwork",
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

func GetPlayListById(id int) (openapi.PlaylistResponse, error) {
	const op = "service.GetPlayListById"
	for _, playlist := range PlayListsInMemory {
		if playlist.id == id {
			return openapi.PlaylistResponse{
				Id:       playlist.id,
				Playlist: playlist.playlist,
			}, nil
		}
	}
	return openapi.PlaylistResponse{}, errors.Build(
		errors.KindNotFound(),
		errors.WithOp(op),
		errors.WithError(fmt.Errorf("Playlist not found")),
		errors.WithSeverity(zerolog.WarnLevel),
	)
}

func AddPlayList(playlist openapi.CreatePlaylistRequestBody) int {
	lastId += 1
	PlayListsInMemory = append(PlayListsInMemory, PlayListStorage{
		id:       lastId,
		playlist: playlist.Playlist,
	})
	return lastId
}
