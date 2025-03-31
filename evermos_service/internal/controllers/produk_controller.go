package controllers

import (
	"evermos-app/internal/dtos"
	"evermos-app/internal/services"
	"evermos-app/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (ctrl *ProductController) CreateProduct(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)

	var req dtos.CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	form, err := c.MultipartForm()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse multipart form: "+err.Error())
	}
	files := form.File["foto_produk"]

	if len(files) == 0 {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "At least one foto produk is required")
	}

	product, err := ctrl.productService.CreateProduct(userID, req, files)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Product created successfully", product)
}

func (ctrl *ProductController) GetAllProducts(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)

	products, err := ctrl.productService.GetAllProducts(userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Products retrieved successfully", products)
}

func (ctrl *ProductController) GetProductsByToko(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)
	tokoID, err := strconv.ParseUint(c.Params("tokoId"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid toko ID")
	}

	products, err := ctrl.productService.GetProductsByTokoId(uint(tokoID), userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusForbidden, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Products retrieved successfully", products)
}

func (ctrl *ProductController) GetProduct(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)
	productID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}

	product, err := ctrl.productService.GetProductById(uint(productID), userID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusForbidden, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Product retrieved successfully", product)
}

func (ctrl *ProductController) UpdateProduct(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)
	productID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}

	var req dtos.UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	product, err := ctrl.productService.UpdateProduct(uint(productID), userID, req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusForbidden, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Product updated successfully", product)
}

func (ctrl *ProductController) DeleteProduct(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)
	productID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}

	if err := ctrl.productService.DeleteProduct(uint(productID), userID); err != nil {
		return utils.ErrorResponse(c, fiber.StatusForbidden, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Product deleted successfully", nil)
}

func (ctrl *ProductController) GetPageAndFilterProducts(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)

	// Ambil parameter pagination
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	// Ambil parameter filter
	filter := &dtos.ProductFilter{
		NamaProduk: c.Query("nama_produk"),
	}
	if idTokoStr := c.Query("id_toko"); idTokoStr != "" {
		idToko, err := strconv.ParseUint(idTokoStr, 10, 32)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid id_toko")
		}
		idTokoUint := uint(idToko)
		filter.IdToko = &idTokoUint
	}
	if idCategoryStr := c.Query("id_category"); idCategoryStr != "" {
		idCategory, err := strconv.ParseUint(idCategoryStr, 10, 32)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid id_category")
		}
		idCategoryUint := uint(idCategory)
		filter.IdCategory = &idCategoryUint
	}

	paginatedResp, err := ctrl.productService.GetPageAndFilterProducts(userID, page, limit, filter)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Products retrieved successfully", paginatedResp)
}
