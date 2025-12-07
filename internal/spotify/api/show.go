package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/topvennie/sortifyr/internal/database/model"
)

func (c *Client) ShowGet(ctx context.Context, user model.User, spotifyID string) (Show, error) {
	var resp Show

	if err := c.request(ctx, user, http.MethodGet, "shows/"+spotifyID, http.NoBody, &resp); err != nil {
		return Show{}, fmt.Errorf("get show %s | %w", spotifyID, err)
	}

	return resp, nil
}
