package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/topvennie/spotify_organizer/internal/database/model"
	"github.com/topvennie/spotify_organizer/pkg/sqlc"
)

type Setting struct {
	repo Repository
}

func (r *Repository) NewSetting() *Setting {
	return &Setting{
		repo: *r,
	}
}

func (s *Setting) GetByUser(ctx context.Context, userID int) (*model.Setting, error) {
	setting, err := s.repo.queries(ctx).SettingGetByUser(ctx, int32(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get sertting by user %d | %w", userID, err)
	}

	return model.SettingModel(setting), nil
}

func (s *Setting) Create(ctx context.Context, setting *model.Setting) error {
	id, err := s.repo.queries(ctx).SettingCreate(ctx, int32(setting.UserID))
	if err != nil {
		return fmt.Errorf("create setting %+v | %w", *setting, err)
	}

	setting.ID = int(id)

	return nil
}

func (s *Setting) Update(ctx context.Context, setting model.Setting) error {
	if err := s.repo.queries(ctx).SettingUpdateByUser(ctx, sqlc.SettingUpdateByUserParams{
		UserID:      int32(setting.UserID),
		LastUpdated: pgtype.Timestamptz{Time: setting.LastUpdate, Valid: !setting.LastUpdate.IsZero()},
	}); err != nil {
		return fmt.Errorf("update setting %+v | %w", setting, err)
	}

	return nil
}
