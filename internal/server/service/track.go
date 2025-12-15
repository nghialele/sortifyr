package service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/internal/database/repository"
	"github.com/topvennie/sortifyr/internal/server/dto"
	"github.com/topvennie/sortifyr/pkg/utils"
	"go.uber.org/zap"
)

type Track struct {
	service Service

	history repository.History
}

func (s *Service) NewTrack() *Track {
	return &Track{
		service: *s,
		history: *s.repo.NewHistory(),
	}
}

func (t *Track) GetHistory(ctx context.Context, filter dto.HistoryFilter) ([]dto.History, error) {
	history, err := t.history.GetPopulatedFiltered(ctx, *filter.ToModel())
	if err != nil {
		zap.S().Error(err)
		return nil, fiber.ErrInternalServerError
	}

	return utils.SliceMap(history, func(h *model.History) dto.History { return dto.HistoryDTO(&h.Track, h) }), nil
}
