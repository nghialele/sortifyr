// Package spotify connects with the spotify API
package spotify

import (
	"context"
	"time"

	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/internal/database/repository"
	"github.com/topvennie/sortifyr/internal/spotify/api"
)

type client struct {
	api api.Client

	album     repository.Album
	artist    repository.Artist
	directory repository.Directory
	history   repository.History
	link      repository.Link
	playlist  repository.Playlist
	show      repository.Show
	track     repository.Track
	user      repository.User
}

var C *client

func Init(repo repository.Repository) error {
	apiClient, err := api.New()
	if err != nil {
		return err
	}

	C = &client{
		api:       *apiClient,
		album:     *repo.NewAlbum(),
		artist:    *repo.NewArtist(),
		directory: *repo.NewDirectory(),
		history:   *repo.NewHistory(),
		link:      *repo.NewLink(),
		playlist:  *repo.NewPlaylist(),
		show:      *repo.NewShow(),
		track:     *repo.NewTrack(),
		user:      *repo.NewUser(),
	}

	if err := C.taskRegister(); err != nil {
		return err
	}

	return nil
}

func (c *client) NewUser(ctx context.Context, user model.User, accessToken, refreshToken string, expiresIn time.Duration) error {
	return c.api.NewUser(ctx, user, accessToken, refreshToken, expiresIn)
}
