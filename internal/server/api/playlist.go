package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/topvennie/spotify_organizer/internal/server/service"
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
