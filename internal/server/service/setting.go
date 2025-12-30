package service

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/topvennie/sortifyr/internal/database/repository"
	"github.com/topvennie/sortifyr/internal/spotify"
	"github.com/topvennie/sortifyr/internal/task"
	"go.uber.org/zap"
)

type Setting struct {
	service Service

	user repository.User
}

func (s *Service) NewSetting() *Setting {
	return &Setting{
		service: *s,
		user:    *s.repo.NewUser(),
	}
}

func (s *Setting) Export(ctx context.Context, userID int, zip []byte) error {
	user, err := s.user.GetByID(ctx, userID)
	if err != nil {
		zap.S().Error(err)
		return fiber.ErrInternalServerError
	}
	if user == nil {
		return fiber.ErrUnauthorized
	}

	if err := spotify.C.TaskExport(ctx, *user, zip); err != nil {
		if errors.Is(err, task.ErrTaskExists) {
			return fiber.NewError(fiber.StatusBadRequest, "Task is already running")
		}
		zap.S().Error(err)
		return fiber.ErrInternalServerError
	}

	return nil
}
