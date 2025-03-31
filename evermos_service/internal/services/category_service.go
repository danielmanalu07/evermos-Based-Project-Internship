package services

import (
	"errors"
	"evermos-app/internal/dtos"
	"evermos-app/internal/models"
	"evermos-app/internal/repositories"
)

type CategoryService struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (s *CategoryService) CreateCategory(req dtos.CreateCategoryRequest) (*dtos.CategoryResponse, error) {

	category := &models.Category{
		NamaCategory: req.NamaCategory,
	}

	if err := s.categoryRepo.Create(category); err != nil {
		return nil, err
	}

	categoryResponse := &dtos.CategoryResponse{}
	categoryResponse.FromModel(category)
	return categoryResponse, nil
}

func (s *CategoryService) GetCategoryById(categoryID uint) (*dtos.CategoryResponse, error) {
	category, err := s.categoryRepo.FindById(categoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	categoryResponse := &dtos.CategoryResponse{}
	categoryResponse.FromModel(category)
	return categoryResponse, nil
}

func (s *CategoryService) GetAllCategories() ([]dtos.CategoryResponse, error) {
	categories, err := s.categoryRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var categoryResponses []dtos.CategoryResponse
	for _, category := range categories {
		categoryResponse := dtos.CategoryResponse{}
		categoryResponse.FromModel(&category)
		categoryResponses = append(categoryResponses, categoryResponse)
	}
	return categoryResponses, nil
}

func (s *CategoryService) UpdateCategory(categoryID uint, req dtos.UpdateCategoryRequest) (*dtos.CategoryResponse, error) {
	category, err := s.categoryRepo.FindById(categoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	if req.NamaCategory != "" {
		category.NamaCategory = req.NamaCategory
	}

	if err := s.categoryRepo.Update(category); err != nil {
		return nil, err
	}

	categoryResponse := &dtos.CategoryResponse{}
	categoryResponse.FromModel(category)
	return categoryResponse, nil
}

func (s *CategoryService) DeleteCategory(categoryID uint) error {
	if _, err := s.categoryRepo.FindById(categoryID); err != nil {
		return errors.New("category not found")
	}

	return s.categoryRepo.Delete(categoryID)
}
