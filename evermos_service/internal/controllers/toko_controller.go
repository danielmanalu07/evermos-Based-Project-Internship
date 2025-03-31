package controllers

import (
	"evermos-app/internal/dtos"
	"evermos-app/internal/services"
	"evermos-app/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TokoController struct {
	tokoService *services.TokoService
}

func NewTokoController(tokoService *services.TokoService) *TokoController {
	return &TokoController{tokoService: tokoService}
}

func (ctrl *TokoController) GetUserTokos(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)

	tokos, err := ctrl.tokoService.GetTokoByUserId(userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Tokos retrieved successfully", tokos)
}

func (ctrl *TokoController) GetToko(c *fiber.Ctx) error {
	tokoID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid toko ID")
	}

	toko, err := ctrl.tokoService.GetTokoById(uint(tokoID))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Toko retrieved successfully", toko)
}

func (ctrl *TokoController) UpdateToko(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)
	tokoID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid toko ID")
	}

	namaToko := c.FormValue("nama_toko")
	file, err := c.FormFile("url_foto")

	req := dtos.CreateTokoRequest{
		NamaToko: namaToko,
	}

	if err == nil {
		req.UrlFoto = file
	}

	toko, err := ctrl.tokoService.UpdateToko(uint(tokoID), userID, dtos.UpdateTokoRequest(req))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Toko updated successfully", toko)
}

func (ctrl *TokoController) DeleteToko(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)
	tokoID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid toko ID")
	}

	if err := ctrl.tokoService.DeleteToko(uint(tokoID), userID); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Toko deleted successfully", nil)
}
