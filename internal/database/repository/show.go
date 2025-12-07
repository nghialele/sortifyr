package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/pkg/sqlc"
)

type Show struct {
	repo Repository
}

func (r *Repository) NewShow() *Show {
	return &Show{
		repo: *r,
	}
}

func (s *Show) GetBySpotify(ctx context.Context, spotifyID string) (*model.Show, error) {
	show, err := s.repo.queries(ctx).ShowGetBySpotify(ctx, spotifyID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get show by spotify %s | %w", spotifyID, err)
	}

	return model.ShowModel(show), nil
}

func (s *Show) Create(ctx context.Context, show *model.Show) error {
	id, err := s.repo.queries(ctx).ShowCreate(ctx, sqlc.ShowCreateParams{
		SpotifyID:     show.SpotifyID,
		Name:          show.Name,
		EpisodeAmount: int32(show.EpisodeAmount),
	})
	if err != nil {
		return fmt.Errorf("create show %+v | %w", *show, err)
	}

	show.ID = int(id)

	return nil
}

func (s *Show) Update(ctx context.Context, show model.Show) error {
	if err := s.repo.queries(ctx).ShowUpdate(ctx, sqlc.ShowUpdateParams{
		ID:            int32(show.ID),
		Name:          show.Name,
		EpisodeAmount: int32(show.EpisodeAmount),
	}); err != nil {
		return fmt.Errorf("update show %+v | %w", show, err)
	}

	return nil
}
