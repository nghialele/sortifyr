package spotify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/topvennie/spotify_organizer/internal/database/model"
	"github.com/topvennie/spotify_organizer/pkg/redis"
	"go.uber.org/zap"
)

const (
	apiAccount = "https://accounts.spotify.com/api/token"
	apiSpotify = "https://api.spotify.com/v1"
)

type accountResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func (c *client) refreshToken(ctx context.Context, user model.User) error {
	zap.S().Info("Refreshing spotify access token")

	refreshToken, err := redis.C.Get(ctx, refreshKey(user)).Result()
	if err != nil {
		if !errors.Is(err, redis.ErrNil) {
			return fmt.Errorf("get redis key %s | %w", refreshKey(user), err)
		}
		return fmt.Errorf("user %+v refresh token not found", user)
	}

	form := url.Values{}
	form.Set("grant_type", "refresh_token")
	form.Set("refresh_token", refreshToken)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiAccount, strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("new http request %w", err)
	}

	creds := base64.StdEncoding.EncodeToString([]byte(c.clientID + ":" + c.clientSecret))

	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+creds)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("do http request %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest {
			return ErrUnauthorized
		}

		return fmt.Errorf("unexpected status code %s", resp.Status)
	}

	var account accountResponse
	if err := json.NewDecoder(resp.Body).Decode(&account); err != nil {
		return fmt.Errorf("decode account response %w", err)
	}

	if account.TokenType != "Bearer" {
		return fmt.Errorf("invalid token type %+v", account)
	}

	if _, err := redis.C.Set(ctx, accessKey(user), account.AccessToken, time.Duration(account.ExpiresIn)*time.Second).Result(); err != nil {
		return fmt.Errorf("set access token %w", err)
	}

	if account.RefreshToken != "" {
		if _, err := redis.C.Set(ctx, refreshKey(user), account.RefreshToken, 0).Result(); err != nil {
			return fmt.Errorf("set refresh token %w", err)
		}
	}

	return nil
}

func (c *client) request(ctx context.Context, user model.User, url string, target any) error {
	zap.S().Infof("do request for url %s", url)

	accessToken, err := redis.C.Get(ctx, accessKey(user)).Result()
	if err != nil {
		if !errors.Is(err, redis.ErrNil) {
			return fmt.Errorf("get redis key %s | %w", accessKey(user), err)
		}

		if err := c.refreshToken(ctx, user); err != nil {
			return err
		}
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", apiSpotify, url), http.NoBody)
	if err != nil {
		return fmt.Errorf("new http request %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("do http request %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	switch resp.StatusCode {
	case 401:
		return errors.New("bad or expired token")

	case 403:
		return errors.New("bad oauth request")

	case 429:
		zap.S().Info("rate limit hit")
		time.Sleep(5 * time.Second)

		return c.request(ctx, user, url, target)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("decode body to json %w", err)
	}

	return nil
}
