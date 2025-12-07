package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/topvennie/sortifyr/internal/database/model"
)

func (c *Client) TrackGet(ctx context.Context, user model.User, spotifyID string) (Track, error) {
	var resp Track

	if err := c.request(ctx, user, http.MethodGet, "tracks/"+spotifyID, http.NoBody, &resp); err != nil {
		return Track{}, fmt.Errorf("get track %s | %w", spotifyID, err)
	}

	return resp, nil
}
