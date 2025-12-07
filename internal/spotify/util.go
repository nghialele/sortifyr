package spotify

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/pkg/image"
	"go.uber.org/zap"
)

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

func uriToID(uri string) string {
	parts := strings.Split(uri, ":")
	if len(parts) != 3 {
		return ""
	}

	return parts[2]
}
