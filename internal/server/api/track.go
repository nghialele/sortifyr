package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/topvennie/sortifyr/internal/server/dto"
	"github.com/topvennie/sortifyr/internal/server/service"
)

type History struct {
	router fiber.Router

	track service.Track
}

func NewHistory(router fiber.Router, service service.Service) *History {
	api := &History{
		router: router.Group("/track"),
		track:  *service.NewTrack(),
	}

	api.createRoutes()

	return api
}

func (r *History) createRoutes() {
	r.router.Get("/history", r.getHistory)
}

func (r *History) getHistory(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(int)
	if !ok {
		return fiber.ErrUnauthorized
	}

	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 0)
	if limit < 1 || page < 0 {
		return fiber.ErrBadRequest
	}

	history, err := r.track.GetHistory(c.Context(), dto.HistoryFilter{
		UserID: userID,
		Limit:  limit,
		Offset: page * limit,
	})
	if err != nil {
		return err
	}

	return c.JSON(history)
}
