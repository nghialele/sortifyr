package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/topvennie/sortifyr/internal/database/model"
)

func (c *Client) TrackGet(ctx context.Context, user model.User, spotifyID string) (Track, error) {
	var resp Track

	if err := c.request(ctx, user, http.MethodGet, "tracks/"+spotifyID, http.NoBody, &resp); err != nil {
		return Track{}, fmt.Errorf("get track %s | %w", spotifyID, err)
	}

	return resp, nil
}

type trackAllResponse struct {
	Tracks []Track `json:"tracks"`
}

func (c *Client) TrackGetAll(ctx context.Context, user model.User, trackIDs []string) ([]Track, error) {
	tracks := make([]Track, 0, len(trackIDs))

	limit := 50

	for i := 0; i < len(trackIDs); i += limit {
		var resp trackAllResponse

		url := "tracks?ids=" + strings.Join(trackIDs[i:min(len(trackIDs), i+limit)], ",")
		if err := c.request(ctx, user, http.MethodGet, url, http.NoBody, &resp); err != nil {
			return nil, err
		}

		tracks = append(tracks, resp.Tracks...)
	}

	return tracks, nil
}
