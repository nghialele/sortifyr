package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/topvennie/spotify_organizer/internal/database/model"
	"github.com/topvennie/spotify_organizer/pkg/sqlc"
	"github.com/topvennie/spotify_organizer/pkg/utils"
)

type Playlist struct {
	repo Repository
}

func (r *Repository) NewPlaylist() *Playlist {
	return &Playlist{
		repo: *r,
	}
}

func (p *Playlist) GetByUserPopulated(ctx context.Context, userID int) ([]*model.Playlist, error) {
	playlists, err := p.repo.queries(ctx).PlaylistGetByUserWithOwner(ctx, int32(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get playlists by user %d | %w", userID, err)
	}

	return utils.SliceMap(playlists, func(r sqlc.PlaylistGetByUserWithOwnerRow) *model.Playlist {
		return model.PlaylistModelPopulated(r.Playlist, r.User)
	}), nil
}

func (p *Playlist) Create(ctx context.Context, playlist *model.Playlist) error {
	id, err := p.repo.queries(ctx).PlaylistCreate(ctx, sqlc.PlaylistCreateParams{
		UserID:        int32(playlist.UserID),
		SpotifyID:     playlist.SpotifyID,
		OwnerUid:      playlist.OwnerUID,
		Name:          playlist.Name,
		Description:   pgtype.Text{String: playlist.Description, Valid: playlist.Description != ""},
		Public:        playlist.Public,
		Tracks:        int32(playlist.Tracks),
		Collaborative: playlist.Collaborative,
	})
	if err != nil {
		return fmt.Errorf("create playlist %+v | %w", *playlist, err)
	}

	playlist.ID = int(id)

	return nil
}

func (p *Playlist) Update(ctx context.Context, playlist model.Playlist) error {
	if err := p.repo.queries(ctx).PlaylistUpdateBySpotify(ctx, sqlc.PlaylistUpdateBySpotifyParams{
		SpotifyID:     playlist.SpotifyID,
		OwnerUid:      playlist.OwnerUID,
		Name:          playlist.Name,
		Description:   pgtype.Text{String: playlist.Description, Valid: playlist.Description != ""},
		Public:        playlist.Public,
		Tracks:        int32(playlist.Tracks),
		Collaborative: playlist.Collaborative,
	}); err != nil {
		return fmt.Errorf("update playlist %+v | %w", playlist, err)
	}

	return nil
}

func (p *Playlist) Delete(ctx context.Context, id int) error {
	if err := p.repo.queries(ctx).PlaylistDelete(ctx, int32(id)); err != nil {
		return fmt.Errorf("delete playlist %d", id)
	}

	return nil
}
