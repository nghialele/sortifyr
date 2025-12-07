package spotify

import (
	"context"

	"github.com/topvennie/sortifyr/internal/database/model"
)

func (c *client) showCheck(ctx context.Context, show *model.Show) error {
	showDB, err := c.show.GetBySpotify(ctx, show.SpotifyID)
	if err != nil {
		return err
	}

	if showDB == nil {
		return c.show.Create(ctx, show)
	}

	show.ID = showDB.ID

	if !showDB.EqualEntry(*show) {
		return c.show.Update(ctx, *show)
	}

	return nil
}
