package spotify

import (
	"context"
	"fmt"

	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/internal/task"
	"github.com/topvennie/sortifyr/pkg/config"
)

const (
	taskPlaylistUID = "task-playlist"
	taskTrackUID    = "task-track"
	taskUserUID     = "task-user"
	taskHistoryUID  = "task-history"
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

	if err := task.Manager.Add(context.Background(), task.NewTask(
		taskHistoryUID,
		"History: Synchronize",
		config.GetDefaultDuration("task.history_s", 10*60),
		c.taskWrap(c.taskHistory),
	)); err != nil {
		return err
	}

	return nil
}

func (c *client) taskWrap(fn func(context.Context, model.User) (string, error)) func(context.Context, []model.User) []task.TaskResult {
	return func(ctx context.Context, users []model.User) []task.TaskResult {
		results := make([]task.TaskResult, 0, len(users))

		for _, user := range users {
			msg, err := fn(ctx, user)
			results = append(results, task.TaskResult{
				User:    user,
				Message: msg,
				Error:   err,
			})
		}

		return results
	}
}

func (c *client) taskPlaylist(ctx context.Context, user model.User) (string, error) {
	msg1, err := c.playlistSync(ctx, user)
	if err != nil {
		return "", fmt.Errorf("synchronize playlists %w", err)
	}

	msg2, err := c.playlistCoverSync(ctx, user)
	if err != nil {
		return "", fmt.Errorf("synchronize playlist covers %w", err)
	}

	return fmt.Sprintf("%s | %s", msg1, msg2), nil
}

func (c *client) taskTrack(ctx context.Context, user model.User) (string, error) {
	msg1, err := c.playlistTrackSync(ctx, user)
	if err != nil {
		return "", fmt.Errorf("synchronize tracks %w", err)
	}

	msg2, err := c.tracksSync(ctx, user)
	if err != nil {
		return "", fmt.Errorf("update playlist tracks based on links %w", err)
	}

	return fmt.Sprintf("%s | %s", msg1, msg2), nil
}

func (c *client) taskUser(ctx context.Context, user model.User) (string, error) {
	if err := c.syncUser(ctx, user); err != nil {
		return "", fmt.Errorf("synchronize users %w", err)
	}

	return "", nil
}

func (c *client) taskHistory(ctx context.Context, user model.User) (string, error) {
	if _, err := c.historySync(ctx, user); err != nil {
		return "", fmt.Errorf("get history %w", err)
	}

	return "", nil
}
