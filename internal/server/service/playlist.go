package service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/internal/database/repository"
	"github.com/topvennie/sortifyr/internal/server/dto"
	"github.com/topvennie/sortifyr/internal/spotify"
	"github.com/topvennie/sortifyr/pkg/storage"
	"github.com/topvennie/sortifyr/pkg/utils"
	"go.uber.org/zap"
)

type Playlist struct {
	service Service

	playlist repository.Playlist
	user     repository.User
}

func (s *Service) NewPlaylist() *Playlist {
	return &Playlist{
		service:  *s,
		playlist: *s.repo.NewPlaylist(),
		user:     *s.repo.NewUser(),
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

func (p *Playlist) GetDuplicates(ctx context.Context, userID int) ([]dto.PlaylistDuplicate, error) {
	playlists, err := p.playlist.GetDuplicateTracksByUser(ctx, userID)
	if err != nil {
		zap.S().Error(err)
		return nil, fiber.ErrInternalServerError
	}

	return utils.SliceMap(playlists, func(p *model.Playlist) dto.PlaylistDuplicate {
		return dto.PlaylistDuplicateDTO(p, &p.Owner, p.Duplicates)
	}), nil
}

func (p *Playlist) RemoveDuplicates(ctx context.Context, userID int) error {
	user, err := p.user.GetByID(ctx, userID)
	if err != nil {
		zap.S().Error(err)
		return fiber.ErrInternalServerError
	}
	if user == nil {
		return fiber.ErrUnauthorized
	}

	go func() {
		if err := spotify.C.PlaylistRemoveDuplicates(context.Background(), *user); err != nil {
			zap.S().Error(err)
		}
	}()

	return nil
}
