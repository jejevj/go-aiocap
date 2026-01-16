package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/go-aiocap/dto"
	"github.com/jejevj/go-aiocap/service"
	"github.com/jejevj/go-aiocap/utils"
)

type (
	CourseController interface {
		AddCourse(ctx *fiber.Ctx) error
	}

	courseController struct {
		userService service.CourseService
	}
)

func NewCourseController(us service.CourseService) CourseController {
	return &courseController{
		userService: us,
	}
}

func (c *courseController) AddCourse(ctx *fiber.Ctx) error {
	var user dto.CourseCreateRequest

	if err := ctx.BodyParser(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.userService.AddCourse(ctx.Context(), user)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}
