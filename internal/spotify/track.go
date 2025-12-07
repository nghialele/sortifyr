package spotify

import (
	"context"
	"fmt"

	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/pkg/utils"
)

func (c *client) tracksSync(ctx context.Context, user model.User) (string, error) {
	directories, err := c.directory.GetByUserPopulated(ctx, user.ID)
	if err != nil {
		return "", err
	}

	playlists, err := c.playlist.GetByUserPopulated(ctx, user.ID)
	if err != nil {
		return "", err
	}

	links, err := c.link.GetAllByUser(ctx, user.ID)
	if err != nil {
		return "", err
	}

	totalAdded := 0

	for _, link := range links {
		var sources []model.Playlist
		var targets []model.Playlist

		switch {
		case link.SourceDirectoryID != 0:
			directory, ok := utils.SliceFind(directories, func(d *model.Directory) bool { return d.ID == link.SourceDirectoryID })
			if !ok {
				return "", fmt.Errorf("database foreign key reference error (source directory) for link %+v", *link)
			}
			sources = (*directory).Playlists

		case link.SourcePlaylistID != 0:
			playlist, ok := utils.SliceFind(playlists, func(p *model.Playlist) bool { return p.ID == link.SourcePlaylistID })
			if !ok {
				return "", fmt.Errorf("database foreign key reference error (source playlist) for link %+v", *link)
			}
			sources = []model.Playlist{**playlist}

		default:
			return "", fmt.Errorf("database foreign key reference error (source) for link %+v", *link)
		}

		switch {
		case link.TargetDirectoryID != 0:
			directory, ok := utils.SliceFind(directories, func(d *model.Directory) bool { return d.ID == link.TargetDirectoryID })
			if !ok {
				return "", fmt.Errorf("database foreign key reference error (target directory) for link %+v", *link)
			}
			targets = (*directory).Playlists

		case link.TargetPlaylistID != 0:
			playlist, ok := utils.SliceFind(playlists, func(p *model.Playlist) bool { return p.ID == link.TargetPlaylistID })
			if !ok {
				return "", fmt.Errorf("database foreign key reference error (target playlist) for link %+v", *link)
			}
			targets = []model.Playlist{**playlist}

		default:
			return "", fmt.Errorf("database foreign key reference error (target) for link %+v", *link)
		}

		for i := range sources {
			for j := range targets {
				added, err := c.trackOneSync(ctx, user, sources[i], targets[j])
				if err != nil {
					return "", err
				}

				totalAdded += added
			}
		}
	}

	return fmt.Sprintf("Added %d", totalAdded), nil
}

func (c *client) trackOneSync(ctx context.Context, user model.User, source, target model.Playlist) (int, error) {
	if source.Equal(target) {
		return 0, nil
	}

	tracksSource, err := c.track.GetByPlaylist(ctx, source.ID)
	if err != nil {
		return 0, err
	}

	tracksTarget, err := c.track.GetByPlaylist(ctx, target.ID)
	if err != nil {
		return 0, err
	}

	toAdd := make([]model.Track, 0)

	for _, trackSource := range tracksSource {
		if _, ok := utils.SliceFind(tracksTarget, func(t *model.Track) bool { return t.Equal(*trackSource) }); !ok {
			toAdd = append(toAdd, *trackSource)
		}
	}

	if err := c.api.PlaylistPostTrackAll(ctx, user, target.SpotifyID, toAdd); err != nil {
		return 0, err
	}

	return len(toAdd), nil
}

// trackCheck creates or updates the track if needed
func (c *client) trackCheck(ctx context.Context, track *model.Track) error {
	trackDB, err := c.track.GetBySpotify(ctx, track.SpotifyID)
	if err != nil {
		return err
	}

	if trackDB == nil {
		return c.track.Create(ctx, track)
	}

	track.ID = trackDB.ID

	if !trackDB.EqualEntry(*track) {
		return c.track.UpdateBySpotify(ctx, *track)
	}

	return nil
}
