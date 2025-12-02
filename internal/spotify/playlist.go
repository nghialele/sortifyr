// Package spotify connects with the spotify API
package spotify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/pkg/utils"
)

type playlistImageAPI struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type playlistAPI struct {
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
	Collaborative bool               `json:"collaborative"`
	Images        []playlistImageAPI `json:"images"`
}

func (p *playlistAPI) toModel(user model.User) model.Playlist {
	url := ""
	maxWidth := -1
	for _, image := range p.Images {
		if image.Width > maxWidth {
			url = image.URL
			maxWidth = image.Width
		}
	}

	return model.Playlist{
		UserID:        user.ID,
		SpotifyID:     p.SpotifyID,
		OwnerUID:      p.Owner.UID,
		Name:          p.Name,
		Description:   p.Description,
		Public:        p.Public,
		TrackAmount:   p.Tracks.Total,
		Collaborative: p.Collaborative,
		CoverURL:      url,
		Owner: model.User{
			UID:         p.Owner.UID,
			DisplayName: p.Owner.DisplayName,
		},
	}
}

type playlistResponse struct {
	Total int           `json:"total"`
	Items []playlistAPI `json:"items"`
}

func (c *client) playlistGetAll(ctx context.Context, user model.User) ([]model.Playlist, error) {
	playlists := make([]model.Playlist, 0)

	limit := 50
	offset := 0

	resp, err := c.playlistGet(ctx, user, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get playlist with limit %d and offset %d | %w", limit, offset, err)
	}
	playlists = append(playlists, utils.SliceMap(resp.Items, func(p playlistAPI) model.Playlist { return p.toModel(user) })...)

	total := resp.Total

	for offset+limit < total {
		offset += limit

		resp, err := c.playlistGet(ctx, user, limit, offset)
		if err != nil {
			return nil, fmt.Errorf("get playlist with limit %d and offset %d | %w", limit, offset, err)
		}
		playlists = append(playlists, utils.SliceMap(resp.Items, func(p playlistAPI) model.Playlist { return p.toModel(user) })...)
	}

	return playlists, nil
}

func (c *client) playlistGet(ctx context.Context, user model.User, limit, offset int) (playlistResponse, error) {
	var resp playlistResponse

	if err := c.request(ctx, user, http.MethodGet, fmt.Sprintf("me/playlists?offset=%d&limit=%d", offset, limit), http.NoBody, &resp); err != nil {
		return resp, fmt.Errorf("get playlist %w", err)
	}

	return resp, nil
}

type playlistTrackAPI struct {
	Track struct {
		SpotifyID  string `json:"id"`
		Name       string `json:"name"`
		Popularity int    `json:"popularity"`
	} `json:"track"`
}

func (p *playlistTrackAPI) toModel() model.Track {
	return model.Track{
		SpotifyID:  p.Track.SpotifyID,
		Name:       p.Track.Name,
		Popularity: p.Track.Popularity,
	}
}

type playlistTrackResponse struct {
	Total int                `json:"total"`
	Items []playlistTrackAPI `json:"items"`
}

func (c *client) playlistGetTrackAll(ctx context.Context, user model.User, playlist model.Playlist) ([]model.Track, error) {
	tracks := make([]model.Track, 0)

	limit := 50
	offset := 0

	resp, err := c.playlistGetTrack(ctx, user, playlist, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get playlist tracks %+v with limit %d and offset %d | %w", playlist, limit, offset, err)
	}
	tracks = append(tracks, utils.SliceMap(resp.Items, func(t playlistTrackAPI) model.Track { return t.toModel() })...)

	total := resp.Total

	for offset+limit < total {
		offset += limit

		resp, err := c.playlistGetTrack(ctx, user, playlist, limit, offset)
		if err != nil {
			return nil, fmt.Errorf("get playlist tracks %+v with limit %d and offset %d | %w", playlist, limit, offset, err)
		}
		tracks = append(tracks, utils.SliceMap(resp.Items, func(t playlistTrackAPI) model.Track { return t.toModel() })...)
	}

	return tracks, nil
}

func (c *client) playlistGetTrack(ctx context.Context, user model.User, playlist model.Playlist, limit, offset int) (playlistTrackResponse, error) {
	var resp playlistTrackResponse

	if err := c.request(ctx, user, http.MethodGet, fmt.Sprintf("playlists/%s/tracks?offset=%d&limit=%d", playlist.SpotifyID, offset, limit), http.NoBody, &resp); err != nil {
		return resp, fmt.Errorf("get playlist tracks %w", err)
	}

	return resp, nil
}

func (c *client) playlistPostTrackAll(ctx context.Context, user model.User, playlist model.Playlist, tracks []model.Track) error {
	current := 0
	total := len(tracks)

	for current < total {
		end := current + 100
		if end > total {
			end = total
		}

		toAdd := tracks[current:end]
		if err := c.playlistPostTrack(ctx, user, playlist, toAdd); err != nil {
			return fmt.Errorf("add tracks %d-%d to playlist %+v | %w", current, end, playlist, err)
		}

		current = end
	}

	return nil
}

func (c *client) playlistPostTrack(ctx context.Context, user model.User, playlist model.Playlist, tracks []model.Track) error {
	payload := struct {
		URIS []string `json:"uris"`
	}{
		URIS: utils.SliceMap(tracks, func(t model.Track) string { return "spotify:track:" + t.SpotifyID }),
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal tracks payload: %w", err)
	}

	body := bytes.NewReader(data)

	if err := c.request(ctx, user, http.MethodPost, fmt.Sprintf("playlists/%s/tracks", playlist.SpotifyID), body, noResp); err != nil {
		return err
	}

	return nil
}
