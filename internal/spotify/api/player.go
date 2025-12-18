package api

import (
	"context"
	"net/http"

	"github.com/topvennie/sortifyr/internal/database/model"
)

type playerHistoryResponse struct {
	Items []History `json:"items"`
}

func (c *Client) PlayerGetHistory(ctx context.Context, user model.User) ([]History, error) {
	var resp playerHistoryResponse
	if err := c.request(ctx, user, http.MethodGet, "me/player/recently-played?limit=50", http.NoBody, &resp); err != nil {
		return nil, err
	}

	return resp.Items, nil
}

type playerCurrentResponse struct {
	Current
}

func (c *Client) PlayerGetCurrent(ctx context.Context, user model.User) (Current, error) {
	var resp playerCurrentResponse
	if err := c.request(ctx, user, http.MethodGet, "me/player/currently-playing", http.NoBody, &resp); err != nil {
		return Current{}, err
	}

	return resp.Current, nil
}
