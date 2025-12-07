package spotify

import (
	"context"

	"github.com/topvennie/sortifyr/internal/database/model"
)

func (c *client) albumCheck(ctx context.Context, album *model.Album) error {
	albumDB, err := c.album.GetBySpotify(ctx, album.SpotifyID)
	if err != nil {
		return err
	}

	if albumDB == nil {
		return c.album.Create(ctx, album)
	}

	album.ID = albumDB.ID

	if !albumDB.EqualEntry(*album) {
		return c.album.Update(ctx, *album)
	}

	return nil
}
