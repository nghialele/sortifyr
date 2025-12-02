package service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/internal/database/repository"
	"github.com/topvennie/sortifyr/internal/server/dto"
	"github.com/topvennie/sortifyr/pkg/storage"
	"github.com/topvennie/sortifyr/pkg/utils"
	"go.uber.org/zap"
)

type Playlist struct {
	service Service

	playlist repository.Playlist
}

func (s *Service) NewPlaylist() *Playlist {
	return &Playlist{
		service:  *s,
		playlist: *s.repo.NewPlaylist(),
	}
}

func (p *Playlist) GetByUser(ctx context.Context, userID int) ([]dto.Playlist, error) {
	playlistsDB, err := p.playlist.GetByUserPopulated(ctx, userID)
	if err != nil {
		zap.S().Error(err)
		return nil, fiber.ErrInternalServerError
	}
	if playlistsDB == nil {
		return []dto.Playlist{}, nil
	}

	return utils.SliceMap(playlistsDB, func(p *model.Playlist) dto.Playlist { return dto.PlaylistDTO(p, &p.Owner) }), nil
}

func (p *Playlist) GetCover(ctx context.Context, playlistID int) ([]byte, error) {
	playlist, err := p.playlist.Get(ctx, playlistID)
	if err != nil {
		zap.S().Error(err)
		return nil, fiber.ErrInternalServerError
	}
	if playlist == nil {
		return nil, fiber.ErrNotFound
	}
	if playlist.CoverID == "" {
		return nil, fiber.ErrNotFound
	}

	cover, err := storage.S.Get(playlist.CoverID)
	if err != nil {
		zap.S().Error(err)
		return nil, fiber.ErrInternalServerError
	}

	return cover, nil
}
