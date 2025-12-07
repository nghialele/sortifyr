package spotify

import (
	"context"

	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/pkg/utils"
)

// syncUser updates the information for every relevant user (for the given user)
func (c *client) syncUser(ctx context.Context, user model.User) error {
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
		newUser, err := c.api.UserGet(ctx, user, userDB)
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

// userCheck creates the user if it doesn't exist yet
func (c *client) userCheck(ctx context.Context, userUID string) error {
	user, err := c.user.GetByUID(ctx, userUID)
	if err != nil {
		return err
	}
	if user != nil {
		return nil
	}

	user = &model.User{
		UID: userUID,
	}

	if err := c.user.Create(ctx, user); err != nil {
		return err
	}

	return nil
}
