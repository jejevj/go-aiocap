package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/go-aiocap/dto"
	"github.com/jejevj/go-aiocap/service"
	"github.com/jejevj/go-aiocap/utils"
)

type (
	CourseCustomerController interface {
		AddCourseCustomer(ctx *fiber.Ctx) error
	}

	courseCustomerController struct {
		userService service.CourseCustomerService
	}
)

func NewCourseCustomerController(us service.CourseCustomerService) CourseCustomerController {
	return &courseCustomerController{
		userService: us,
	}
}

func (c *courseCustomerController) AddCourseCustomer(ctx *fiber.Ctx) error {
	var user dto.CourseCustomerCreateRequest

	if err := ctx.BodyParser(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.userService.AddCourseCustomer(ctx.Context(), user)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}
