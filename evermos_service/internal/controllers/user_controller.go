package controllers

import (
	"evermos-app/internal/dtos"
	"evermos-app/internal/services"
	"evermos-app/internal/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *services.UserService
	jwtSecret   string
}

func NewUserController(userService *services.UserService, jwtSecret string) *UserController {
	return &UserController{userService: userService, jwtSecret: jwtSecret}
}

func (ctrl *UserController) RegisterUser(c *fiber.Ctx) error {
	var req dtos.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user, err := ctrl.userService.Register(req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var userResponse dtos.UserResponse
	userResponse.FromModel(user)

	return utils.SuccessResponse(c, fiber.StatusCreated, "User Registered Successfully", userResponse)
}

func (ctrl *UserController) LoginUser(c *fiber.Ctx) error {
	var req dtos.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user, err := ctrl.userService.Login(req, ctrl.jwtSecret)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "User Logged in Successfully", user)
}

func (ctrl *UserController) GetProfile(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	user, err := ctrl.userService.GetProfile(userId)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Get user successfully", user)
}

func (ctrl *UserController) UpdateProfile(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	var req dtos.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	user, err := ctrl.userService.UpdateProfile(userId, req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "User Updated Successfully", user)
}

func (ctrl *UserController) LogoutUser(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid token format")
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")

	if err := ctrl.userService.Logout(token); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "User logged out successfully", nil)
}

func (ctrl *UserController) DeleteAccount(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid token format")
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")

	if err := ctrl.userService.DeleteAccount(userID, token); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Account deleted successfully", nil)
}
