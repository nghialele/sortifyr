package spotify

import (
	"context"

	"github.com/topvennie/spotify_organizer/internal/database/model"
	"github.com/topvennie/spotify_organizer/pkg/utils"
)

// userSync updates the information for every relevant user (for the given user)
func (c *client) userSync(ctx context.Context, user model.User) error {
	// Get all relevant users
	playlists, err := c.playlist.GetByUserPopulated(ctx, user.ID)
	if err != nil {
		return err
	}
	if playlists == nil {
		return nil
	}

	usersDB := utils.SliceMap(playlists, func(p *model.Playlist) model.User { return p.Owner })
	usersDB = utils.SliceUnique(usersDB)

	// Get all spotify users
	usersSpotify := make([]model.User, 0, len(usersDB))
	for _, userDB := range usersDB {
		newUser, err := c.userGet(ctx, user, userDB)
		if err != nil {
			return err
		}

		usersSpotify = append(usersSpotify, newUser)
	}

	toUpdate := make([]model.User, 0)

	for _, userSpotify := range usersSpotify {
		if _, ok := utils.SliceFind(usersDB, func(u model.User) bool { return u.Equal(userSpotify) }); !ok {
			toUpdate = append(toUpdate, userSpotify)
		}
	}

	for _, user := range toUpdate {
		if err := c.user.Update(ctx, user); err != nil {
			return err
		}
	}

	return nil
}

type userResponse struct {
	DisplayName string `json:"display_name"`
}

func (c *client) userGet(ctx context.Context, user, spotifyUser model.User) (model.User, error) {
	var resp userResponse

	if err := c.request(ctx, user, "users/"+spotifyUser.UID, &resp); err != nil {
		return model.User{}, err
	}

	spotifyUser.DisplayName = resp.DisplayName

	return spotifyUser, nil
}
