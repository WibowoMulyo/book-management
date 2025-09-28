package models

import (
	"time"
)

type Category struct {
	ID         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name" validate:"required,min=1,max=255"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	CreatedBy  string    `json:"created_by" db:"created_by"`
	ModifiedAt time.Time `json:"modified_at" db:"modified_at"`
	ModifiedBy string    `json:"modified_by" db:"modified_by"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=1,max=255"`
}
