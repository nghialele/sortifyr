package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/topvennie/sortifyr/internal/server/service"
)

type Playlist struct {
	router fiber.Router

	playlist service.Playlist
}

func NewPlaylist(router fiber.Router, service service.Service) *Playlist {
	api := &Playlist{
		router:   router.Group("/playlist"),
		playlist: *service.NewPlaylist(),
	}

	api.routes()

	return api
}

func (p *Playlist) routes() {
	p.router.Get("/", p.getAll)
	p.router.Get("/cover/:id", p.getCover)
	p.router.Get("/duplicate", p.getDuplicates)
	p.router.Post("/duplicate", p.removeDuplicates)
}

func (p *Playlist) getAll(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(int)
	if !ok {
		return fiber.ErrUnauthorized
	}

	playlists, err := p.playlist.GetByUser(c.Context(), userID)
	if err != nil {
		return err
	}

	return c.JSON(playlists)
}

func (p *Playlist) getCover(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	cover, err := p.playlist.GetCover(c.Context(), id)
	if err != nil {
		return err
	}

	c.Set("Content-Type", mimeWEBP)

	return sendCached(c, cover)
}

func (p *Playlist) getDuplicates(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(int)
	if !ok {
		return fiber.ErrUnauthorized
	}

	playlists, err := p.playlist.GetDuplicates(c.Context(), userID)
	if err != nil {
		return err
	}

	return c.JSON(playlists)
}

func (p *Playlist) removeDuplicates(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(int)
	if !ok {
		return fiber.ErrUnauthorized
	}

	if err := p.playlist.RemoveDuplicates(c.Context(), userID); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
