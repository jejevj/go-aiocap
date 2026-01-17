package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/go-aiocap/service"
	"github.com/jejevj/go-aiocap/utils"
)

type LogController interface {
	GetLogs(ctx *fiber.Ctx) error
	GetLogByID(ctx *fiber.Ctx) error
}

type logController struct {
	logService service.LogService
}

func NewLogController(logService service.LogService) LogController {
	return &logController{
		logService: logService,
	}
}

func (c *logController) GetLogs(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	logs, err := c.logService.GetLogs(ctx.Context(), page, limit)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to get logs", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess("Successfully retrieved logs", logs)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *logController) GetLogByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		res := utils.BuildResponseFailed("Invalid ID", "ID parameter is required", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	log, err := c.logService.GetLogByID(ctx.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to get log", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess("Successfully retrieved log", log)
	return ctx.Status(http.StatusOK).JSON(res)
}
