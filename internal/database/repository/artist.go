package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/pkg/sqlc"
	"github.com/topvennie/sortifyr/pkg/utils"
)

type Artist struct {
	repo Repository
}

func (r *Repository) NewArtist() *Artist {
	return &Artist{
		repo: *r,
	}
}

func (a *Artist) GetAll(ctx context.Context) ([]*model.Artist, error) {
	artists, err := a.repo.queries(ctx).ArtistGetAll(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get all artist %w", err)
	}

	return utils.SliceMap(artists, model.ArtistModel), nil
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

func (a *Artist) GetByAlbum(ctx context.Context, albumID int) ([]*model.Artist, error) {
	artists, err := a.repo.queries(ctx).ArtistGetByAlbum(ctx, int32(albumID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get artists by album %d | %w", albumID, err)
	}

	return utils.SliceMap(artists, model.ArtistModel), nil
}

func (a *Artist) GetByTrack(ctx context.Context, trackID int) ([]*model.Artist, error) {
	artists, err := a.repo.queries(ctx).ArtistGetByTrack(ctx, int32(trackID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get artists by track %d | %w", trackID, err)
	}

	return utils.SliceMap(artists, model.ArtistModel), nil
}

func (a *Artist) Create(ctx context.Context, artist *model.Artist) error {
	id, err := a.repo.queries(ctx).ArtistCreate(ctx, sqlc.ArtistCreateParams{
		SpotifyID:  artist.SpotifyID,
		Name:       toString(artist.Name),
		Followers:  toInt(artist.Followers),
		Popularity: toInt(artist.Popularity),
		CoverID:    toString(artist.CoverID),
		CoverUrl:   toString(artist.CoverURL),
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
		Name:       toString(artist.Name),
		Followers:  toInt(artist.Followers),
		Popularity: toInt(artist.Popularity),
		CoverID:    toString(artist.CoverID),
		CoverUrl:   toString(artist.CoverURL),
	}); err != nil {
		return fmt.Errorf("update artist %+v | %w", artist, err)
	}

	return nil
}
