package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/go-aiocap/constants"
	"github.com/jejevj/go-aiocap/dto"
	"github.com/jejevj/go-aiocap/service"
	"github.com/jejevj/go-aiocap/utils"
	"github.com/xuri/excelize/v2"
)

type (
	CourseCustomerController interface {
		AddCourseCustomer(ctx *fiber.Ctx) error
		Update(ctx *fiber.Ctx) error
		GetAllCC(ctx *fiber.Ctx) error
		GetSingle(ctx *fiber.Ctx) error
		ExportExcel(ctx *fiber.Ctx) error
	}

	courseCustomerController struct {
		userService service.CourseCustomerService
		us2         service.UserService
	}
)

func NewCourseCustomerController(us service.CourseCustomerService, us2 service.UserService) CourseCustomerController {
	return &courseCustomerController{
		userService: us,
		us2:         us2,
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

func (c *courseCustomerController) Update(ctx *fiber.Ctx) error {
	var req dto.CourseCustomerUpdateRequest

	userId := ctx.Locals("user_id").(string)
	res2, err := c.us2.GetUserById(ctx.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	if res2.Role != constants.ENUM_ROLE_ADMIN {
		res := utils.BuildResponseFailed(dto.MESSAGE_ROLE_ADMIN, "Forbidden Action", nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.userService.UpdateCourseCustomer(ctx.Context(), req, req.ID.String())
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *courseCustomerController) GetAllCC(ctx *fiber.Ctx) error {
	var req dto.PaginationRequest
	if err := ctx.BodyParser(&req); err != nil {
		fmt.Printf("Body parsing error: %v\n", err)
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	fmt.Printf("Parsed request: %+v\n", req)
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.userService.GetAllCourseCustomer(ctx.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_USER,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	return ctx.Status(http.StatusOK).JSON(resp)
}

func (c *courseCustomerController) GetSingle(ctx *fiber.Ctx) error {
	var req dto.CourseCustomerGetDetailsRequest

	// Parse the request body into req
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.userService.GetCourseCustomerById(ctx.Context(), req.ID.String())
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

func (c *courseCustomerController) ExportExcel(ctx *fiber.Ctx) error {
	// Parse pagination request
	var req dto.PaginationRequest
	if err := ctx.QueryParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Get all course customers with pagination
	result, err := c.userService.GetAllCourseCustomer(ctx.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	// Create a new Excel file
	f := excelize.NewFile()

	// Set sheet name
	sheetName := "Course Customers"
	f.SetSheetName("Sheet1", sheetName)

	// Set headers
	headers := []string{
		"ID",
		"Customer Name",
		"Customer Email",
		"Contact Name",
		"Phone Number",
		"Customer Address",
		"Created By ID",
		"Changed By ID",
	}

	// Set header row
	for i, header := range headers {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellValue(sheetName, cell, header)
	}

	// Set header style
	style, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#4472C4"},
			Pattern: 1,
		},
	})
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create style", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(res)
	}

	// Apply header style
	for i := 0; i < len(headers); i++ {
		cell := fmt.Sprintf("%c1", 'A'+i)
		f.SetCellStyle(sheetName, cell, cell, style)
	}

	// Populate data
	for i, customer := range result.Data {
		row := i + 2 // Start from row 2 (row 1 is header)

		// Set cell values
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), customer.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), customer.CustomerName)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), customer.CustomerEmail)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), customer.ContactName)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), customer.PhoneNumber)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), customer.CustomerAddress)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), customer.CreatedByID.String())
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), customer.ChangedByID.String())

	}

	// Auto-size columns
	for i := 0; i < len(headers); i++ {
		f.SetColWidth(sheetName, string(rune('A'+i)), string(rune('A'+i)), 20)
	}

	// Set response headers
	ctx.Response().Header.SetContentType("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Response().Header.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"course_customers_%s.xlsx\"", time.Now().Format("20060102_150405")))

	// Write file to response
	if err := f.Write(ctx.Response().BodyWriter()); err != nil {
		res := utils.BuildResponseFailed("Failed to write Excel file", err.Error(), nil)
		return ctx.Status(http.StatusInternalServerError).JSON(res)
	}

	return nil
}
