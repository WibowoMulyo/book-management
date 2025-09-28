package models

import (
	"time"
)

type Book struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" validate:"required,min=1,max=255"`
	Description string    `json:"description" db:"description"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	ReleaseYear int       `json:"release_year" db:"release_year" validate:"required,min=1980,max=2024"`
	Price       int       `json:"price" db:"price" validate:"required,min=0"`
	TotalPage   int       `json:"total_page" db:"total_page" validate:"required,min=1"`
	Thickness   string    `json:"thickness" db:"thickness"`
	CategoryID  int       `json:"category_id" db:"category_id" validate:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	CreatedBy   string    `json:"created_by" db:"created_by"`
	ModifiedAt  time.Time `json:"modified_at" db:"modified_at"`
	ModifiedBy  string    `json:"modified_by" db:"modified_by"`
}

type BookWithCategory struct {
	Book
	CategoryName string `json:"category_name" db:"category_name"`
}

type CreateBookRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	ReleaseYear int    `json:"release_year" validate:"required,min=1980,max=2024"`
	Price       int    `json:"price" validate:"required,min=0"`
	TotalPage   int    `json:"total_page" validate:"required,min=1"`
	CategoryID  int    `json:"category_id" validate:"required"`
}

type UpdateBookRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	ReleaseYear int    `json:"release_year" validate:"required,min=1980,max=2024"`
	Price       int    `json:"price" validate:"required,min=0"`
	TotalPage   int    `json:"total_page" validate:"required,min=1"`
	CategoryID  int    `json:"category_id" validate:"required"`
}

// CalculateThickness menghitung ketebalan buku berdasarkan total halaman
func (b *Book) CalculateThickness() {
	if b.TotalPage >= 100 {
		b.Thickness = "tebal"
	} else {
		b.Thickness = "tipis"
	}
}
