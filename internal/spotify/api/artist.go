package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/topvennie/sortifyr/internal/database/model"
)

func (c *Client) ArtistGet(ctx context.Context, user model.User, spotifyID string) (Artist, error) {
	var resp Artist

	if err := c.request(ctx, user, http.MethodGet, "artists/"+spotifyID, http.NoBody, &resp); err != nil {
		return Artist{}, fmt.Errorf("get artist %s | %w", spotifyID, err)
	}

	return resp, nil
}

type artistAllResponse struct {
	Artists []Artist `json:"artists"`
}

func (c *Client) ArtistGetAll(ctx context.Context, user model.User, artistIDs []string) ([]Artist, error) {
	artists := make([]Artist, 0, len(artistIDs))

	limit := 50

	for i := 0; i < len(artistIDs); i += limit {
		var resp artistAllResponse

		url := "artists?ids=" + strings.Join(artistIDs[i:min(len(artistIDs), i+limit)], ",")
		if err := c.request(ctx, user, http.MethodGet, url, http.NoBody, &resp); err != nil {
			return nil, err
		}

		artists = append(artists, resp.Artists...)
	}

	return artists, nil
}
