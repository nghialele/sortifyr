package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/topvennie/sortifyr/internal/database/model"
)

func (c *Client) AlbumGet(ctx context.Context, user model.User, spotifyID string) (Album, error) {
	var resp Album

	if err := c.request(ctx, user, http.MethodGet, "albums/"+spotifyID, http.NoBody, &resp); err != nil {
		return Album{}, fmt.Errorf("get album %s | %w", spotifyID, err)
	}

	return resp, nil
}
