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

func (h *History) GetPopulatedFiltered(ctx context.Context, filter model.HistoryFilter) ([]*model.History, error) {
	history, err := h.repo.queries(ctx).HistoryGetPopulatedFiltered(ctx, sqlc.HistoryGetPopulatedFilteredParams{
		Column1: int32(filter.UserID),
		Limit:   int32(filter.Limit),
		Offset:  int32(filter.Offset),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get filtered populated history %+v | %w", filter, err)
	}

	return utils.SliceMap(history, func(h sqlc.HistoryGetPopulatedFilteredRow) *model.History {
		history := model.HistoryModel(h.History)
		history.Track = *model.TrackModel(h.Track)

		return history
	}), nil
}

func (h *History) Create(ctx context.Context, history *model.History) error {
	id, err := h.repo.queries(ctx).HistoryCreate(ctx, sqlc.HistoryCreateParams{
		UserID:     int32(history.UserID),
		TrackID:    int32(history.TrackID),
		PlayedAt:   toTime(history.PlayedAt),
		AlbumID:    toInt(history.AlbumID),
		ArtistID:   toInt(history.ArtistID),
		PlaylistID: toInt(history.PlaylistID),
		ShowID:     toInt(history.ShowID),
	})
	if err != nil {
		return fmt.Errorf("create history %+v | %w", *history, err)
	}

	history.ID = int(id)

	return nil
}
