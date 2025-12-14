package spotify

import (
	"context"

	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/internal/spotify/api"
	"github.com/topvennie/sortifyr/pkg/utils"
)

// artistUpdate updates local artist instances to match the spotify data.
// It updates all artists, regardless of the user given.
// However the given user's access token is used.
func (c *client) artistUpdate(ctx context.Context, user model.User) error {
	artistsDB, err := c.artist.GetAll(ctx)
	if err != nil {
		return err
	}

	artistsSpotifyAPI, err := c.api.ArtistGetAll(ctx, user, utils.SliceMap(artistsDB, func(a *model.Artist) string { return a.SpotifyID }))
	if err != nil {
		return err
	}
	artistsSpotify := utils.SliceMap(artistsSpotifyAPI, func(a api.Artist) model.Artist { return a.ToModel() })

	for i := range artistsSpotify {
		artistDB, ok := utils.SliceFind(artistsDB, func(a *model.Artist) bool { return a.Equal(artistsSpotify[i]) })
		if !ok {
			// Artist not found
			continue
		}

		artistsSpotify[i].ID = (*artistDB).ID

		// Bring the artist data up to date
		if !(*artistDB).EqualEntry(artistsSpotify[i]) {
			if err := c.artist.Update(ctx, artistsSpotify[i]); err != nil {
				return err
			}
		}
	}

	return nil
}
