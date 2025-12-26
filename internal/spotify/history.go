package spotify

import (
	"context"
	"time"

	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/internal/spotify/api"
)

func (c *client) historySync(ctx context.Context, user model.User) error {
	latest, err := c.history.GetLatest(ctx, user.ID)
	if err != nil {
		return err
	}
	if latest == nil {
		latest = &model.History{}
	}

	current, err := c.api.PlayerGetCurrent(ctx, user)
	if err != nil {
		return err
	}

	now := time.Now()

	if !current.IsPlaying {
		return nil
	}

	// Listen at least 20 seconds
	if current.ProgressMs < 20*1000 {
		return nil
	}

	// Add a 5 second buffer
	if latest.PlayedAt.Add(5 * time.Second).After(now.Add(time.Duration(-current.ProgressMs) * time.Millisecond)) {
		return nil
	}

	return c.historyOneSync(ctx, user, api.History{
		Track:    current.Track,
		PlayedAt: now.Add(time.Duration(-current.ProgressMs) * time.Millisecond),
		Context:  current.Context,
	})
}

func (c *client) historyOneSync(ctx context.Context, user model.User, history api.History) error {
	historyModel := history.ToModel(user)

	spotifyID := history.Track.SpotifyID
	if history.Track.LinkedFrom.SpotifyID != "" {
		spotifyID = history.Track.LinkedFrom.SpotifyID
	}

	track := model.Track{SpotifyID: spotifyID}
	if err := c.historyTrackCheck(ctx, &track); err != nil {
		return err
	}
	historyModel.TrackID = track.ID

	contextSpotifyID := uriToID(history.Context.URI)

	switch history.Context.Type {
	case "album":
		album := model.Album{SpotifyID: contextSpotifyID}
		if err := c.historyAlbumCheck(ctx, &album); err != nil {
			return err
		}
		historyModel.AlbumID = album.ID
	case "artist":
		artist := model.Artist{SpotifyID: contextSpotifyID}
		if err := c.historyArtistCheck(ctx, &artist); err != nil {
			return err
		}
		historyModel.ArtistID = artist.ID
	case "playlist":
		playlist := model.Playlist{SpotifyID: contextSpotifyID}
		if err := c.historyPlaylistCheck(ctx, &playlist); err != nil {
			return err
		}
		historyModel.PlaylistID = playlist.ID
	case "show":
		show := model.Show{SpotifyID: contextSpotifyID}
		if err := c.historyShowCheck(ctx, &show); err != nil {
			return err
		}
		historyModel.ShowID = show.ID
	}

	if err := c.history.Create(ctx, &historyModel); err != nil {
		return err
	}

	return nil
}

func (c *client) historyTrackCheck(ctx context.Context, track *model.Track) error {
	trackDB, err := c.track.GetBySpotify(ctx, track.SpotifyID)
	if err != nil {
		return err
	}

	if trackDB == nil {
		if err := c.track.Create(ctx, track); err != nil {
			return err
		}
	} else {
		track.ID = trackDB.ID
	}

	return nil
}

func (c *client) historyArtistCheck(ctx context.Context, artist *model.Artist) error {
	artistDB, err := c.artist.GetBySpotify(ctx, artist.SpotifyID)
	if err != nil {
		return err
	}

	if artistDB == nil {
		if err := c.artist.Create(ctx, artist); err != nil {
			return err
		}
	} else {
		artist.ID = artistDB.ID
	}

	return nil
}

func (c *client) historyAlbumCheck(ctx context.Context, album *model.Album) error {
	albumDB, err := c.album.GetBySpotify(ctx, album.SpotifyID)
	if err != nil {
		return err
	}

	if albumDB == nil {
		if err := c.album.Create(ctx, album); err != nil {
			return err
		}
	} else {
		album.ID = albumDB.ID
	}

	return nil
}

func (c *client) historyPlaylistCheck(ctx context.Context, playlist *model.Playlist) error {
	playlistDB, err := c.playlist.GetBySpotify(ctx, playlist.SpotifyID)
	if err != nil {
		return err
	}

	if playlistDB == nil {
		if err := c.playlist.Create(ctx, playlist); err != nil {
			return err
		}
	} else {
		playlist.ID = playlistDB.ID
	}

	return nil
}

func (c *client) historyShowCheck(ctx context.Context, show *model.Show) error {
	showDB, err := c.show.GetBySpotify(ctx, show.SpotifyID)
	if err != nil {
		return err
	}

	if showDB == nil {
		if err := c.show.Create(ctx, show); err != nil {
			return err
		}
	} else {
		show.ID = showDB.ID
	}

	return nil
}
