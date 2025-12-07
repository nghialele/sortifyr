package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/pkg/sqlc"
	"github.com/topvennie/sortifyr/pkg/utils"
)

type History struct {
	repo Repository
}

func (r *Repository) NewHistory() *History {
	return &History{
		repo: *r,
	}
}

func (h *History) GetLatest(ctx context.Context, userID int) (*model.History, error) {
	history, err := h.repo.queries(ctx).HistoryGetLatestByUser(ctx, int32(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get latest history %w", err)
	}

	return model.HistoryModel(history), nil
}

func (h *History) GetByPlaylist(ctx context.Context, playlistID int) ([]*model.History, error) {
	history, err := h.repo.queries(ctx).HistoryGetByPlaylist(ctx, pgtype.Int4{Int32: int32(playlistID), Valid: playlistID != 0})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get history by playlist %d | %w", playlistID, err)
	}

	return utils.SliceMap(history, model.HistoryModel), nil
}

func (h *History) Create(ctx context.Context, history *model.History) error {
	id, err := h.repo.queries(ctx).HistoryCreate(ctx, sqlc.HistoryCreateParams{
		UserID:     int32(history.UserID),
		TrackID:    int32(history.TrackID),
		PlayedAt:   pgtype.Timestamptz{Time: history.PlayedAt, Valid: !history.PlayedAt.IsZero()},
		AlbumID:    pgtype.Int4{Int32: int32(history.AlbumID), Valid: history.AlbumID != 0},
		ArtistID:   pgtype.Int4{Int32: int32(history.ArtistID), Valid: history.ArtistID != 0},
		PlaylistID: pgtype.Int4{Int32: int32(history.PlaylistID), Valid: history.PlaylistID != 0},
		ShowID:     pgtype.Int4{Int32: int32(history.ShowID), Valid: history.ShowID != 0},
	})
	if err != nil {
		return fmt.Errorf("create history %+v | %w", *history, err)
	}

	history.ID = int(id)

	return nil
}
