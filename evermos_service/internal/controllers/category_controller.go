package controllers

import (
	"evermos-app/internal/dtos"
	"evermos-app/internal/services"
	"evermos-app/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	categoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

func (ctrl *CategoryController) CreateCategory(c *fiber.Ctx) error {
	var req dtos.CreateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	category, err := ctrl.categoryService.CreateCategory(req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Category created successfully", category)
}

func (ctrl *CategoryController) GetCategory(c *fiber.Ctx) error {
	categoryID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid category ID")
	}

	category, err := ctrl.categoryService.GetCategoryById(uint(categoryID))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Category retrieved successfully", category)
}

func (ctrl *CategoryController) GetAllCategories(c *fiber.Ctx) error {
	categories, err := ctrl.categoryService.GetAllCategories()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Categories retrieved successfully", categories)
}

func (ctrl *CategoryController) UpdateCategory(c *fiber.Ctx) error {
	categoryID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid category ID")
	}

	var req dtos.UpdateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	category, err := ctrl.categoryService.UpdateCategory(uint(categoryID), req)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Category updated successfully", category)
}

func (ctrl *CategoryController) DeleteCategory(c *fiber.Ctx) error {
	categoryID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid category ID")
	}

	if err := ctrl.categoryService.DeleteCategory(uint(categoryID)); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Category deleted successfully", nil)
}
