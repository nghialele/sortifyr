// Package api interacts with the spotify api
package api

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/pkg/config"
	"github.com/topvennie/sortifyr/pkg/redis"
)

type Client struct {
	clientID     string
	clientSecret string
}

func New() (*Client, error) {
	clientID := config.GetString("auth.spotify.client.id")
	clientSecret := config.GetString("auth.spotify.client.secret")

	if clientID == "" || clientSecret == "" {
		return nil, errors.New("client id or client secret not set")
	}

	return &Client{
		clientID:     clientID,
		clientSecret: clientSecret,
	}, nil
}

func (c *Client) NewUser(ctx context.Context, user model.User, accessToken, refreshToken string, expiresIn time.Duration) error {
	if _, err := redis.C.Set(ctx, accessKey(user), accessToken, expiresIn).Result(); err != nil {
		return fmt.Errorf("set access token %w", err)
	}

	if _, err := redis.C.Set(ctx, refreshKey(user), refreshToken, 0).Result(); err != nil {
		return fmt.Errorf("set refresh token %w", err)
	}

	return nil
}
