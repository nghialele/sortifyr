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

type Album struct {
	repo Repository
}

func (r *Repository) NewAlbum() *Album {
	return &Album{
		repo: *r,
	}
}

func (a *Album) GetAll(ctx context.Context) ([]*model.Album, error) {
	albums, err := a.repo.queries(ctx).AlbumGetAll(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get albums %w", err)
	}

	return utils.SliceMap(albums, model.AlbumModel), nil
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

func (a *Album) GetByUser(ctx context.Context, userID int) ([]*model.Album, error) {
	albums, err := a.repo.queries(ctx).AlbumGetByUser(ctx, int32(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get albums by user  %d | %w", userID, err)
	}

	return utils.SliceMap(albums, model.AlbumModel), nil
}

func (a *Album) Create(ctx context.Context, album *model.Album) error {
	id, err := a.repo.queries(ctx).AlbumCreate(ctx, sqlc.AlbumCreateParams{
		SpotifyID:   album.SpotifyID,
		Name:        toString(album.Name),
		TrackAmount: toInt(album.TrackAmount),
		Popularity:  toInt(album.Popularity),
		CoverID:     toString(album.CoverID),
		CoverUrl:    toString(album.CoverURL),
	})
	if err != nil {
		return fmt.Errorf("create album %+v | %w", *album, err)
	}

	album.ID = int(id)

	return nil
}

func (a *Album) CreateUser(ctx context.Context, user *model.AlbumUser) error {
	id, err := a.repo.queries(ctx).AlbumUserCreate(ctx, sqlc.AlbumUserCreateParams{
		UserID:  int32(user.UserID),
		AlbumID: int32(user.AlbumID),
	})
	if err != nil {
		return fmt.Errorf("create album user %+v | %w", *user, err)
	}

	user.ID = int(id)

	return nil
}

func (a *Album) CreateArtist(ctx context.Context, artist *model.AlbumArtist) error {
	id, err := a.repo.queries(ctx).AlbumArtistCreate(ctx, sqlc.AlbumArtistCreateParams{
		AlbumID:  int32(artist.AlbumID),
		ArtistID: int32(artist.ArtistID),
	})
	if err != nil {
		return fmt.Errorf("create album artist %+v | %w", *artist, err)
	}

	artist.ID = int(id)

	return nil
}

func (a *Album) Update(ctx context.Context, album model.Album) error {
	if err := a.repo.queries(ctx).AlbumUpdate(ctx, sqlc.AlbumUpdateParams{
		ID:          int32(album.ID),
		Name:        toString(album.Name),
		TrackAmount: toInt(album.TrackAmount),
		Popularity:  toInt(album.Popularity),
		CoverID:     toString(album.CoverID),
		CoverUrl:    toString(album.CoverURL),
	}); err != nil {
		return fmt.Errorf("update album %+v | %w", album, err)
	}

	return nil
}

func (a *Album) DeleteUserByUserAlbum(ctx context.Context, user model.AlbumUser) error {
	if err := a.repo.queries(ctx).AlbumUserDeleteByUserAlbum(ctx, sqlc.AlbumUserDeleteByUserAlbumParams{
		UserID:  int32(user.UserID),
		AlbumID: int32(user.AlbumID),
	}); err != nil {
		return fmt.Errorf("delete album user %+v | %w", user, err)
	}

	return nil
}

func (a *Album) DeleteArtistByArtistAlbum(ctx context.Context, artist model.AlbumArtist) error {
	if err := a.repo.queries(ctx).AlbumArtistDeleteByArtistAlbum(ctx, sqlc.AlbumArtistDeleteByArtistAlbumParams{
		ArtistID: int32(artist.ArtistID),
		AlbumID:  int32(artist.AlbumID),
	}); err != nil {
		return fmt.Errorf("delete album artist %v | %w", artist, err)
	}

	return nil
}
