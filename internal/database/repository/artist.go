package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/pkg/sqlc"
)

type Artist struct {
	repo Repository
}

func (r *Repository) NewArtist() *Artist {
	return &Artist{
		repo: *r,
	}
}

func (a *Artist) GetBySpotify(ctx context.Context, spotifyID string) (*model.Artist, error) {
	artist, err := a.repo.queries(ctx).ArtistGetBySpotify(ctx, spotifyID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get artist by spotify %s | %w", spotifyID, err)
	}

	return model.ArtistModel(artist), nil
}

func (a *Artist) Create(ctx context.Context, artist *model.Artist) error {
	id, err := a.repo.queries(ctx).ArtistCreate(ctx, sqlc.ArtistCreateParams{
		SpotifyID:  artist.SpotifyID,
		Name:       artist.Name,
		Followers:  int32(artist.Followers),
		Popularity: int32(artist.Popularity),
	})
	if err != nil {
		return fmt.Errorf("create artist %+v | %w", *artist, err)
	}

	artist.ID = int(id)

	return nil
}

func (a *Artist) Update(ctx context.Context, artist model.Artist) error {
	if err := a.repo.queries(ctx).ArtistUpdate(ctx, sqlc.ArtistUpdateParams{
		ID:         int32(artist.ID),
		Name:       artist.Name,
		Followers:  int32(artist.Followers),
		Popularity: int32(artist.Popularity),
	}); err != nil {
		return fmt.Errorf("update artist %+v | %w", artist, err)
	}

	return nil
}
