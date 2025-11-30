package service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/topvennie/spotify_organizer/internal/database/repository"
	"github.com/topvennie/spotify_organizer/internal/server/dto"
	"github.com/topvennie/spotify_organizer/internal/task"
	"github.com/topvennie/spotify_organizer/pkg/utils"
	"go.uber.org/zap"
)

type Task struct {
	service Service

	task repository.Task
	user repository.User
}

func (s *Service) NewTask() *Task {
	return &Task{
		service: *s,
		task:    *s.repo.NewTask(),
		user:    *s.repo.NewUser(),
	}
}

func (t *Task) GetTasks() ([]dto.Task, error) {
	tasks, err := task.Manager.Tasks()
	if err != nil {
		zap.S().Error(err)
		return nil, fiber.ErrInternalServerError
	}
	if tasks == nil {
		return []dto.Task{}, nil
	}

	return utils.SliceMap(tasks, dto.TaskDTO), nil
}

func (t *Task) GetHistory(ctx context.Context, filter dto.TaskFilter) ([]dto.TaskHistory, error) {
	tasks, err := t.task.GetFiltered(ctx, *filter.ToModel())
	if err != nil {
		zap.S().Error(err)
		return nil, fiber.ErrInternalServerError
	}
	if tasks == nil {
		return []dto.TaskHistory{}, nil
	}

	return utils.SliceMap(tasks, dto.TaskHistoryDTO), nil
}

func (t *Task) Start(ctx context.Context, userID int, taskUID string) error {
	user, err := t.user.GetByID(ctx, userID)
	if err != nil {
		zap.S().Error(err)
		return fiber.ErrInternalServerError
	}
	if user == nil {
		return fiber.ErrUnauthorized
	}

	taskModel, err := t.task.GetByUID(ctx, taskUID)
	if err != nil {
		zap.S().Error(err)
		return fiber.ErrInternalServerError
	}
	if taskModel == nil {
		return fiber.ErrNotFound
	}

	return task.Manager.RunByUID(taskUID, *user)
}
