package service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/topvennie/spotify_organizer/internal/database/model"
	"github.com/topvennie/spotify_organizer/internal/database/repository"
	"github.com/topvennie/spotify_organizer/internal/server/dto"
	"go.uber.org/zap"
)

type Setting struct {
	service Service

	setting repository.Setting
}

func (s *Service) NewSetting() *Setting {
	return &Setting{
		service: *s,
		setting: *s.repo.NewSetting(),
	}
}

func (s *Setting) GetByUser(ctx context.Context, userID int) (dto.Setting, error) {
	setting, err := s.setting.GetByUser(ctx, userID)
	if err != nil {
		zap.S().Error(err)
		return dto.Setting{}, fiber.ErrInternalServerError
	}
	if setting == nil {
		return dto.Setting{}, fiber.ErrNotFound
	}

	return dto.SettingDTO(setting), nil
}

func (s *Setting) Create(ctx context.Context, user dto.User) (dto.Setting, error) {
	setting := model.Setting{UserID: user.ID}

	if err := s.setting.Create(ctx, &setting); err != nil {
		zap.S().Error(err)
		return dto.Setting{}, fiber.ErrInternalServerError
	}

	return dto.SettingDTO(&setting), nil
}
