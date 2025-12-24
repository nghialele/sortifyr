package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/pkg/concurrent"
	"github.com/topvennie/sortifyr/pkg/utils"
)

func (c *Client) PlaylistGet(ctx context.Context, user model.User, spotifyID string) (Playlist, error) {
	var resp Playlist

	if err := c.request(ctx, user, http.MethodGet, "playlists/"+spotifyID, http.NoBody, &resp); err != nil {
		return Playlist{}, fmt.Errorf("get playlist %s | %w", spotifyID, err)
	}

	return resp, nil
}

type playlistUserResponse struct {
	Total int        `json:"total"`
	Items []Playlist `json:"items"`
}

func (c *Client) PlaylistGetUser(ctx context.Context, user model.User) ([]Playlist, error) {
	playlists := make([]Playlist, 0)

	total := 51
	limit := 50

	for i := 0; i < total; i += limit {
		var resp playlistUserResponse

		if err := c.request(ctx, user, http.MethodGet, fmt.Sprintf("me/playlists?offset=%d&limit=%d", i, limit), http.NoBody, &resp); err != nil {
			return nil, fmt.Errorf("get playlists with limit %d and offset %d | %w", limit, i, err)
		}

		playlists = append(playlists, resp.Items...)
		total = resp.Total
	}

	return playlists, nil
}

type playlistTrackAPI struct {
	Track Track `json:"track"`
}

type playlistTrackResponse struct {
	Total int                `json:"total"`
	Items []playlistTrackAPI `json:"items"`
}

func (c *Client) PlaylistGetTrackAll(ctx context.Context, user model.User, spotifyID string) ([]Track, error) {
	wg := concurrent.NewLimitedWaitGroup(12)

	var mu sync.Mutex
	var errs []error

	tracks := make([]Track, 0)
	total := 0
	limit := 50

	// Do the first request to get the total
	var resp playlistTrackResponse
	if err := c.request(ctx, user, http.MethodGet, fmt.Sprintf("playlists/%s/tracks?offset=%d&limit=%d", spotifyID, 0, limit), http.NoBody, &resp); err != nil {
		return nil, fmt.Errorf("get playlist tracks with limit %d and offset %d | %w", limit, 0, err)
	}

	tracks = append(tracks, utils.SliceMap(resp.Items, func(t playlistTrackAPI) Track { return t.Track })...)
	total = resp.Total

	for i := limit; i < total; i += limit {
		wg.Go(func() {
			var resp playlistTrackResponse

			if err := c.request(ctx, user, http.MethodGet, fmt.Sprintf("playlists/%s/tracks?offset=%d&limit=%d", spotifyID, i, limit), http.NoBody, &resp); err != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("get playlist tracks with limit %d and offset %d | %w", limit, i, err))
				mu.Unlock()
				return
			}

			mu.Lock()
			tracks = append(tracks, utils.SliceMap(resp.Items, func(t playlistTrackAPI) Track { return t.Track })...)
			mu.Unlock()
		})
	}

	wg.Wait()

	return tracks, nil
}

type playlistTrackPayload struct {
	URIs []string `json:"uris"`
}

func (c *Client) PlaylistPostTrackAll(ctx context.Context, user model.User, spotifyID string, tracks []model.Track) error {
	current := 0
	total := len(tracks)

	for current < total {
		end := min(current+100, total)

		payload := playlistTrackPayload{
			URIs: utils.SliceMap(tracks[current:end], func(t model.Track) string { return "spotify:track:" + t.SpotifyID }),
		}

		data, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("marshal tracks payload: %w", err)
		}

		body := bytes.NewReader(data)

		if err := c.request(ctx, user, http.MethodPost, fmt.Sprintf("playlists/%s/tracks", spotifyID), body, noResp); err != nil {
			return err
		}

		current = end
	}

	return nil
}
