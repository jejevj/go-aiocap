package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/go-aiocap/dto"
	"github.com/jejevj/go-aiocap/service"
	"github.com/jejevj/go-aiocap/utils"
)

type (
	UserController interface {
		Register(ctx *fiber.Ctx) error
		Login(ctx *fiber.Ctx) error
		Me(ctx *fiber.Ctx) error
		GetAllUser(ctx *fiber.Ctx) error
		SendVerificationEmail(ctx *fiber.Ctx) error
		VerifyEmail(ctx *fiber.Ctx) error
		Update(ctx *fiber.Ctx) error
		Delete(ctx *fiber.Ctx) error
	}

	userController struct {
		userService service.UserService
	}
)

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

// @Summary Register new user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.UserCreateRequest true "User data"
// @Success 200 {object} dto.UserResponse
// @Router /api/user [post]
func (c *userController) Register(ctx *fiber.Ctx) error {
	var user dto.UserCreateRequest

	if err := ctx.BodyParser(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.userService.RegisterUser(ctx.Context(), user)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

// @Summary Get all users
// @Description Get list of all users with pagination
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param per_page query int false "Items per page"
// @Success 200 {object} dto.UserPaginationResponse
// @Router /api/user [get]
func (c *userController) GetAllUser(ctx *fiber.Ctx) error {
	var req dto.PaginationRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.userService.GetAllUserWithPagination(ctx.Context(), req)
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

// @Summary Get current user
// @Description Get user information for the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} dto.UserResponse
// @Router /api/user/me [get]
func (c *userController) Me(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(string)

	result, err := c.userService.GetUserById(ctx.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param login body dto.UserLoginRequest true "Login credentials"
// @Success 200 {object} dto.UserLoginResponse
// @Router /api/user/login [post]
func (c *userController) Login(ctx *fiber.Ctx) error {
	var req dto.UserLoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	result, err := c.userService.Verify(ctx.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN, result)
	return ctx.Status(http.StatusOK).JSON(res)
}


// @Summary Send verification email
// @Description Send verification email to user
// @Tags users
// @Accept json
// @Produce json
// @Param email body dto.SendVerificationEmailRequest true "Email address"
// @Success 200 {object} utils.Response
// @Router /api/user/send_verification_email [post]
func (c *userController) SendVerificationEmail(ctx *fiber.Ctx) error {
	var req dto.SendVerificationEmailRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	err := c.userService.SendVerificationEmail(ctx.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SEND_VERIFICATION_EMAIL_SUCCESS, nil)
	return ctx.Status(http.StatusOK).JSON(res)
}
// @Summary Verify email
// @Description Verify user email with token
// @Tags users
// @Accept json
// @Produce json
// @Param token body dto.VerifyEmailRequest true "Verification token"
// @Success 200 {object} dto.VerifyEmailResponse
// @Router /api/user/verify_email [post]
func (c *userController) VerifyEmail(ctx *fiber.Ctx) error {
	var req dto.VerifyEmailRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	result, err := c.userService.VerifyEmail(ctx.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_VERIFY_EMAIL, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_VERIFY_EMAIL, result)
	return ctx.Status(http.StatusOK).JSON(res)
}

// @Summary Update user
// @Description Update user information
// @Tags users
// @Accept json
// @Produce json
// @Security Bearer
// @Param user body dto.UserUpdateRequest true "User data"
// @Success 200 {object} dto.UserUpdateResponse
// @Router /api/user [patch]
func (c *userController) Update(ctx *fiber.Ctx) error {
	var req dto.UserUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	userId := ctx.Locals("user_id").(string)
	result, err := c.userService.UpdateUser(ctx.Context(), req, userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_USER, result)
	return ctx.Status(http.StatusOK).JSON(res)
}
// @Summary Delete user
// @Description Delete user account
// @Tags users
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} utils.Response
// @Router /api/user [delete]
func (c *userController) Delete(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(string)

	if err := c.userService.DeleteUser(ctx.Context(), userId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_USER, err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_USER, nil)
	return ctx.Status(http.StatusOK).JSON(res)
}
