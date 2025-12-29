package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/topvennie/sortifyr/internal/database/model"
	"github.com/topvennie/sortifyr/internal/server/dto"
	"github.com/topvennie/sortifyr/internal/server/service"
)

type Task struct {
	router fiber.Router

	task service.Task
}

func NewTask(router fiber.Router, service service.Service) *Task {
	api := &Task{
		router: router.Group("/task"),
		task:   *service.NewTask(),
	}

	api.createRoutes()

	return api
}

func (r *Task) createRoutes() {
	r.router.Get("/", r.getTasks)
	r.router.Get("/history", r.getHistory)
	r.router.Post("/start/:uid", r.start)
}

func (r *Task) getTasks(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(int)
	if !ok {
		return fiber.ErrUnauthorized
	}

	tasks, err := r.task.GetTasks(c.Context(), userID)
	if err != nil {
		return err
	}

	return c.JSON(tasks)
}

func (r *Task) getHistory(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(int)
	if !ok {
		return fiber.ErrUnauthorized
	}

	uid := c.Query("uid")

	var result *model.TaskResult
	if v := c.Query("result"); v != "" {
		switch v {
		case string(model.TaskSuccess), string(model.TaskFailed):
			r := model.TaskResult(v)
			result = &r
		}
	}

	var recurring *bool
	if v := c.Query("recurring"); v != "" {
		if b, err := strconv.ParseBool(v); err == nil {
			recurring = &b
		}
	}

	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	if limit < 1 || page < 1 {
		return fiber.ErrBadRequest
	}

	tasks, err := r.task.GetHistory(c.Context(), dto.TaskFilter{
		UserID:    userID,
		TaskUID:   uid,
		Result:    result,
		Limit:     limit,
		Recurring: recurring,
		Offset:    (page - 1) * limit,
	})
	if err != nil {
		return err
	}

	return c.JSON(tasks)
}

func (r *Task) start(c *fiber.Ctx) error {
	uid := c.Params("uid")

	userID, ok := c.Locals("userID").(int)
	if !ok {
		return fiber.ErrUnauthorized
	}

	if err := r.task.Start(c.Context(), userID, uid); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusAccepted)
}
