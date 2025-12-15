package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/topvennie/sortifyr/internal/server/dto"
	"github.com/topvennie/sortifyr/internal/server/service"
)

type Track struct {
	router fiber.Router

	track service.Track
}

func NewTrack(router fiber.Router, service service.Service) *Track {
	api := &Track{
		router: router.Group("/track"),
		track:  *service.NewTrack(),
	}

	api.createRoutes()

	return api
}

func (r *Track) createRoutes() {
	r.router.Get("/history", r.getHistory)
	r.router.Get("/added", r.getAdded)
	r.router.Get("/deleted", r.getDeleted)
}

func (r *Track) getHistory(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(int)
	if !ok {
		return fiber.ErrUnauthorized
	}

	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	if limit < 1 || page < 1 {
		return fiber.ErrBadRequest
	}

	history, err := r.track.GetHistory(c.Context(), dto.HistoryFilter{
		UserID: userID,
		Limit:  limit,
		Offset: (page - 1) * limit,
	})
	if err != nil {
		return err
	}

	return c.JSON(history)
}

func (r *Track) getAdded(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(int)
	if !ok {
		return fiber.ErrUnauthorized
	}

	playlistID := c.QueryInt("playlist_id")

	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	if limit < 1 || page < 1 {
		return fiber.ErrBadRequest
	}

	tracks, err := r.track.GetAdded(c.Context(), dto.TrackFilter{
		UserID:     userID,
		PlaylistID: playlistID,
		Limit:      limit,
		Offset:     (page - 1) * limit,
	})
	if err != nil {
		return err
	}

	return c.JSON(tracks)
}

func (r *Track) getDeleted(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(int)
	if !ok {
		return fiber.ErrUnauthorized
	}

	playlistID := c.QueryInt("playlist_id")

	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	if limit < 1 || page < 1 {
		return fiber.ErrBadRequest
	}

	tracks, err := r.track.GetDeleted(c.Context(), dto.TrackFilter{
		UserID:     userID,
		PlaylistID: playlistID,
		Limit:      limit,
		Offset:     (page - 1) * limit,
	})
	if err != nil {
		return err
	}

	return c.JSON(tracks)
}
