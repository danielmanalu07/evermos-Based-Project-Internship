package dtos

import (
	"evermos-app/internal/models"
	"time"
)

type CreateCategoryRequest struct {
	NamaCategory string `json:"nama_category" validate:"required"`
}

type UpdateCategoryRequest struct {
	NamaCategory string `json:"nama_category,omitempty"`
}

type CategoryResponse struct {
	ID           uint      `json:"id"`
	NamaCategory string    `json:"nama_category"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (cr *CategoryResponse) FromModel(category *models.Category) {
	cr.ID = category.ID
	cr.NamaCategory = category.NamaCategory
	cr.CreatedAt = category.CreatedAt
	cr.UpdatedAt = category.UpdatedAt
}
