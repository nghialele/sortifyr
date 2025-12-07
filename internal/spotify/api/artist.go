package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/topvennie/sortifyr/internal/database/model"
)

func (c *Client) ArtistGet(ctx context.Context, user model.User, spotifyID string) (Artist, error) {
	var resp Artist

	if err := c.request(ctx, user, http.MethodGet, "artists/"+spotifyID, http.NoBody, &resp); err != nil {
		return Artist{}, fmt.Errorf("get artist %s | %w", spotifyID, err)
	}

	return resp, nil
}
