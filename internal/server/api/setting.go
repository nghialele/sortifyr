package api

import (
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/topvennie/sortifyr/internal/server/service"
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

	api.routes()

	return api
}

func (s *Setting) routes() {
	s.router.Post("/export", s.export)
}

func (s *Setting) export(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(int)
	if !ok {
		return fiber.ErrUnauthorized
	}

	fileHeader, err := c.FormFile("zip")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	file, err := fileHeader.Open()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	defer func() {
		// nolint:errcheck // Unlucky if it fails
		_ = file.Close()
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	if err := s.setting.Export(c.Context(), userID, data); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
