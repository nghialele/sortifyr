package spotify

import (
	"context"

	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/internal/spotify/api"
	"github.com/topvennie/sortifyr/pkg/utils"
)

// trackUpdate updates local track instances to match the spotify data.
// It updates all tracks, regardless of the user given.
// However the given user's access token is used.
func (c *client) trackUpdate(ctx context.Context, user model.User) error {
	tracksDB, err := c.track.GetAll(ctx)
	if err != nil {
		return err
	}

	tracksSpotifyAPI, err := c.api.TrackGetAll(ctx, user, utils.SliceMap(tracksDB, func(t *model.Track) string { return t.SpotifyID }))
	if err != nil {
		return nil
	}
	tracksSpotify := utils.SliceMap(tracksSpotifyAPI, func(t api.Track) model.Track { return t.ToModel() })

	for i := range tracksSpotify {
		trackDB, ok := utils.SliceFind(tracksDB, func(t *model.Track) bool { return t.Equal(tracksSpotify[i]) })
		if !ok {
			// Track not found
			continue
		}

		tracksSpotify[i].ID = (*trackDB).ID

		// Bring track up to date
		if !(*trackDB).EqualEntry(tracksSpotify[i]) {
			if err := c.track.Update(ctx, tracksSpotify[i]); err != nil {
				return err
			}
		}

		// Bring the track artists up to date
		artistsDB, err := c.artist.GetByTrack(ctx, (*trackDB).ID)
		if err != nil {
			return err
		}

		artistsSpotify := utils.SliceMap(tracksSpotifyAPI[i].Artists, func(a api.Artist) model.Artist { return a.ToModel() })

		if err := syncUserData(syncUserDataStruct[model.Artist]{
			DB:     utils.SliceDereference(artistsDB),
			API:    artistsSpotify,
			Equal:  func(a1, a2 model.Artist) bool { return a1.Equal(a2) },
			Get:    func(a model.Artist) (*model.Artist, error) { return c.artist.GetBySpotify(ctx, a.SpotifyID) },
			Create: func(a *model.Artist) error { return c.artist.Create(ctx, a) },
			CreateUserLink: func(a model.Artist) error {
				return c.track.CreateArtist(ctx, &model.TrackArtist{TrackID: (*trackDB).ID, ArtistID: a.ID})
			},
			DeleteUserLink: func(a model.Artist) error {
				return c.track.DeleteArtistByArtistTrack(ctx, model.TrackArtist{TrackID: (*trackDB).ID, ArtistID: a.ID})
			},
		}); err != nil {
			return err
		}
	}

	return nil
}
