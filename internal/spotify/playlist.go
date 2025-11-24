// Package spotify connects with the spotify API
package spotify

import (
	"context"
	"fmt"

	"github.com/topvennie/spotify_organizer/internal/database/model"
	"github.com/topvennie/spotify_organizer/pkg/utils"
)

func (c *client) playlistSync(ctx context.Context, user model.User) error {
	playlistsDB, err := c.playlist.GetByUserPopulated(ctx, user.ID)
	if err != nil {
		return err
	}

	playlistsSpotify, err := c.playlistGetAll(ctx, user)
	if err != nil {
		return err
	}

	toCreate := make([]model.Playlist, 0)
	toUpdate := make([]model.Playlist, 0)
	toDelete := make([]model.Playlist, 0)

	// Find the playlists that need to be created or updated
	for i := range playlistsSpotify {
		playlistDB, ok := utils.SliceFind(playlistsDB, func(p *model.Playlist) bool { return p.Equal(playlistsSpotify[i]) })
		if !ok {
			// Playlist doesn't exist yet
			// Create it
			toCreate = append(toCreate, playlistsSpotify[i])
			continue
		}

		// Playlist already exist
		// But is it still completely the same?
		if !(*playlistDB).EqualEntry(playlistsSpotify[i]) {
			// Not completely the same anymore
			// Update it
			toUpdate = append(toUpdate, playlistsSpotify[i])
		}
	}

	for i := range toCreate {
		if err := c.playlistUserCheck(ctx, toCreate[i].OwnerUID); err != nil {
			return err
		}
		if err := c.playlist.Create(ctx, &toCreate[i]); err != nil {
			return err
		}
	}

	for i := range toUpdate {
		if err := c.playlistUserCheck(ctx, toUpdate[i].OwnerUID); err != nil {
			return err
		}
		if err := c.playlist.Update(ctx, toUpdate[i]); err != nil {
			return err
		}
	}

	// New and updated entries are now in the database
	// Let's bring our local copy up to date
	playlistsDB, err = c.playlist.GetByUserPopulated(ctx, user.ID)
	if err != nil {
		return err
	}

	// Find the playlists that need to be deleted
	for _, playlistDB := range playlistsDB {
		_, ok := utils.SliceFind(playlistsSpotify, func(p model.Playlist) bool { return p.SpotifyID == playlistDB.SpotifyID })
		if !ok {
			// Playlist no longer exists in the user's account
			// So delete it
			toDelete = append(toDelete, *playlistDB)
		}
	}

	for i := range toDelete {
		if err := c.playlist.Delete(ctx, toDelete[i].ID); err != nil {
			return err
		}
	}

	return nil
}

type playlist struct {
	SpotifyID string `json:"id"`
	Owner     struct {
		UID         string `json:"id"`
		DisplayName string `json:"display_name"`
	} `json:"owner"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
	Tracks      struct {
		Total int `json:"total"`
	} `json:"tracks"`
	Collaborative bool `json:"collaborative"`
}

func (p *playlist) toModel(user model.User) *model.Playlist {
	return &model.Playlist{
		UserID:        user.ID,
		SpotifyID:     p.SpotifyID,
		OwnerUID:      p.Owner.UID,
		Name:          p.Name,
		Description:   p.Description,
		Public:        p.Public,
		Tracks:        p.Tracks.Total,
		Collaborative: p.Collaborative,
		Owner: model.User{
			UID:         p.Owner.UID,
			DisplayName: p.Owner.DisplayName,
		},
	}
}

type playListResponse struct {
	Total int        `json:"total"`
	Items []playlist `json:"items"`
}

func (c *client) playlistGetAll(ctx context.Context, user model.User) ([]model.Playlist, error) {
	playlists := make([]model.Playlist, 0)

	limit := 50
	offset := 0

	resp, err := c.playlistGet(ctx, user, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get playlist with limit %d and offset %d | %w", limit, offset, err)
	}
	playlists = append(playlists, utils.SliceMap(resp.Items, func(p playlist) model.Playlist { return *p.toModel(user) })...)

	total := resp.Total

	for offset+limit < total {
		offset += limit

		resp, err := c.playlistGet(ctx, user, limit, offset)
		if err != nil {
			return nil, fmt.Errorf("get playlist with limit %d and offset %d | %w", limit, offset, err)
		}
		playlists = append(playlists, utils.SliceMap(resp.Items, func(p playlist) model.Playlist { return *p.toModel(user) })...)
	}

	return playlists, nil
}

func (c *client) playlistGet(ctx context.Context, user model.User, limit, offset int) (playListResponse, error) {
	var resp playListResponse

	if err := c.request(ctx, user, fmt.Sprintf("me/playlists?offset=%d&limit=%d", offset, limit), &resp); err != nil {
		return resp, fmt.Errorf("get playlist %w", err)
	}

	return resp, nil
}

// playlistUserCheck creates the user if it doesn't exist yet
func (c *client) playlistUserCheck(ctx context.Context, userUID string) error {
	user, err := c.user.GetByUID(ctx, userUID)
	if err != nil {
		return err
	}
	if user != nil {
		return nil
	}

	user = &model.User{
		UID: userUID,
	}

	if err := c.user.Create(ctx, user); err != nil {
		return err
	}

	return nil
}
