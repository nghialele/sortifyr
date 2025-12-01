package spotify

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/topvennie/spotify_organizer/internal/database/model"
	"github.com/topvennie/spotify_organizer/pkg/image"
	"go.uber.org/zap"
)

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

	if !trackDB.Equal(*track) {
		return c.track.UpdateBySpotify(ctx, *track)
	}

	return nil
}

func (c *client) getCover(playlist model.Playlist) ([]byte, error) {
	zap.S().Infof("Get cover image for %s", playlist.Name)

	if playlist.CoverURL == "" {
		return nil, nil
	}

	req, err := http.NewRequest(http.MethodGet, playlist.CoverURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("new http request %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get image data %+v | %w", playlist, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("unexpected status %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read image data %+v | %w", playlist, err)
	}

	webp, err := image.ToWebp(data)
	if err != nil {
		return nil, err
	}

	return webp, nil
}

func accessKey(user model.User) string {
	return user.UID + ":spotify:access_token"
}

func refreshKey(user model.User) string {
	return user.UID + ":spotify:refresh_token"
}
