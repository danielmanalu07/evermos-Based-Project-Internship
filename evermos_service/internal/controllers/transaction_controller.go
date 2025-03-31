package controllers

import (
	"evermos-app/internal/dtos"
	"evermos-app/internal/services"
	"evermos-app/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	transactionService *services.TransactionService
}

func NewTransactionController(transactionService *services.TransactionService) *TransactionController {
	return &TransactionController{transactionService: transactionService}
}

func (ctrl *TransactionController) CreateTransaksi(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)

	var req dtos.CreateTransaksiRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	transaksi, err := ctrl.transactionService.CreateTransaksi(userID, req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Transaksi created successfully", transaksi)
}

func (ctrl *TransactionController) GetTransaksi(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)
	transaksiID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid transaksi ID")
	}

	transaksi, err := ctrl.transactionService.GetTransaksiById(uint(transaksiID), userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusForbidden, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Transaksi retrieved successfully", transaksi)
}

func (ctrl *TransactionController) GetAllTransaksis(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)

	transaksis, err := ctrl.transactionService.GetAllTransaksis(userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Transaksis retrieved successfully", transaksis)
}

func (ctrl *TransactionController) GetPageAndFilterTransaksis(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)

	// Ambil parameter pagination
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	// Ambil parameter filter
	filter := &dtos.TransactionFilter{
		KodeInvoice: c.Query("kode_invoice"),
		MethodBayar: c.Query("method_bayar"),
	}

	paginatedResp, err := ctrl.transactionService.GetPageAndFilterTransaksis(userID, page, limit, filter)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Transaksis retrieved successfully", paginatedResp)
}
