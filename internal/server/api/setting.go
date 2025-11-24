package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/topvennie/spotify_organizer/internal/server/service"
)

type Setting struct {
	router fiber.Router

	setting service.Setting
}

func NewSetting(router fiber.Router, service service.Service) *Setting {
	api := &Setting{
		router:  router.Group("/setting"),
		setting: *service.NewSetting(),
	}

	api.createRoutes()

	return api
}

func (s *Setting) createRoutes() {
	s.router.Get("/", s.get)
}

func (s *Setting) get(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(int)
	if !ok {
		return fiber.ErrUnauthorized
	}

	setting, err := s.setting.GetByUser(c.Context(), userID)
	if err != nil {
		return err
	}

	return c.JSON(setting)
}
