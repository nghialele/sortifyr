package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/pkg/sqlc"
)

type Album struct {
	repo Repository
}

func (r *Repository) NewAlbum() *Album {
	return &Album{
		repo: *r,
	}
}

func (a *Album) GetBySpotify(ctx context.Context, spotifyID string) (*model.Album, error) {
	album, err := a.repo.queries(ctx).AlbumGetBySpotify(ctx, spotifyID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get album by spotify %s | %w", spotifyID, err)
	}

	return model.AlbumModel(album), nil
}

func (a *Album) Create(ctx context.Context, album *model.Album) error {
	id, err := a.repo.queries(ctx).AlbumCreate(ctx, sqlc.AlbumCreateParams{
		SpotifyID:   album.SpotifyID,
		Name:        album.Name,
		TrackAmount: int32(album.TrackAmount),
		Popularity:  int32(album.Popularity),
	})
	if err != nil {
		return fmt.Errorf("create album %+v | %w", *album, err)
	}

	album.ID = int(id)

	return nil
}

func (a *Album) Update(ctx context.Context, album model.Album) error {
	if err := a.repo.queries(ctx).AlbumUpdate(ctx, sqlc.AlbumUpdateParams{
		ID:          int32(album.ID),
		Name:        album.Name,
		TrackAmount: int32(album.TrackAmount),
		Popularity:  int32(album.Popularity),
	}); err != nil {
		return fmt.Errorf("update album %+v | %w", album, err)
	}

	return nil
}
