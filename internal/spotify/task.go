package spotify

import (
	"context"
	"fmt"

	"github.com/topvennie/spotify_organizer/internal/database/model"
	"github.com/topvennie/spotify_organizer/internal/task"
	"github.com/topvennie/spotify_organizer/pkg/config"
)

const (
	taskPlaylistUID = "task-playlist"
	taskTrackUID    = "task-track"
	taskUserUID     = "task-user"
)

func (c *client) taskRegister() error {
	if err := task.Manager.Add(context.Background(), task.NewTask(
		taskPlaylistUID,
		"Playlist: Synchronize",
		config.GetDefaultDuration("task.playlist_s", 6*60*60),
		c.taskWrap(c.taskPlaylist),
	)); err != nil {
		return err
	}

	if err := task.Manager.Add(context.Background(), task.NewTask(
		taskTrackUID,
		"Track: Synchronize & update playlists by links",
		config.GetDefaultDuration("task.track_s", 60*60),
		c.taskWrap(c.taskTrack),
	)); err != nil {
		return err
	}

	if err := task.Manager.Add(context.Background(), task.NewTask(
		taskUserUID,
		"User: Synchronize",
		config.GetDefaultDuration("task.user_s", 24*60*60),
		c.taskWrap(c.taskUser),
	)); err != nil {
		return err
	}

	return nil
}

func (c *client) taskWrap(fn func(context.Context, model.User) error) func(context.Context, *model.User) error {
	return func(ctx context.Context, user *model.User) error {
		if user != nil {
			return fn(ctx, *user)
		}

		users, err := c.user.GetActualAll(ctx)
		if err != nil {
			return err
		}

		for _, user := range users {
			if err := fn(ctx, *user); err != nil {
				return err
			}
		}

		return nil
	}
}

func (c *client) taskPlaylist(ctx context.Context, user model.User) error {
	if err := c.syncPlaylist(ctx, user); err != nil {
		return fmt.Errorf("synchronize playlists %w", err)
	}

	if err := c.syncPlaylistCover(ctx, user); err != nil {
		return fmt.Errorf("synchronize playlist covers %w", err)
	}

	return nil
}

func (c *client) taskTrack(ctx context.Context, user model.User) error {
	if err := c.syncPlaylistTrack(ctx, user); err != nil {
		return fmt.Errorf("synchronize tracks %w", err)
	}

	if err := c.syncLink(ctx, user); err != nil {
		return fmt.Errorf("update playlist tracks based on links %w", err)
	}

	return nil
}

func (c *client) taskUser(ctx context.Context, user model.User) error {
	if err := c.syncUser(ctx, user); err != nil {
		return fmt.Errorf("synchronize users %w", err)
	}

	return nil
}
