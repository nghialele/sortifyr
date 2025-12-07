package spotify

import (
	"context"

	"github.com/topvennie/sortifyr/internal/database/model"
)

func (c *client) artistCheck(ctx context.Context, artist *model.Artist) error {
	artistDB, err := c.artist.GetBySpotify(ctx, artist.SpotifyID)
	if err != nil {
		return err
	}

	if artistDB == nil {
		return c.artist.Create(ctx, artist)
	}

	artist.ID = artistDB.ID

	if !artistDB.EqualEntry(*artist) {
		return c.artist.Update(ctx, *artist)
	}

	return nil
}
