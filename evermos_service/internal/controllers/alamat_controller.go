package controllers

import (
	"evermos-app/internal/dtos"
	"evermos-app/internal/services"
	"evermos-app/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AlamatController struct {
	alamatService *services.AlamatService
}

func NewAlamatController(alamatService *services.AlamatService) *AlamatController {
	return &AlamatController{alamatService: alamatService}
}

func (ctrl *AlamatController) CreateAlamat(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)

	var req dtos.CreateAlamatRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	alamat, err := ctrl.alamatService.CreateAlamat(userID, req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Alamat created successfully", alamat)
}

func (ctrl *AlamatController) GetUserAlamats(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)

	alamats, err := ctrl.alamatService.GetAlamatByUserId(userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Alamats retrieved successfully", alamats)
}

func (ctrl *AlamatController) GetAlamat(c *fiber.Ctx) error {
	alamatID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid alamat ID")
	}

	alamat, err := ctrl.alamatService.GetAlamatById(uint(alamatID))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Alamat retrieved successfully", alamat)
}

func (ctrl *AlamatController) UpdateAlamat(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)
	alamatID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid alamat ID")
	}

	var req dtos.UpdateAlamatRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	alamat, err := ctrl.alamatService.UpdateAlamat(uint(alamatID), userID, req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Alamat updated successfully", alamat)
}

func (ctrl *AlamatController) DeleteAlamat(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)
	alamatID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid alamat ID")
	}

	if err := ctrl.alamatService.DeleteAlamat(uint(alamatID), userID); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Alamat deleted successfully", nil)
}
