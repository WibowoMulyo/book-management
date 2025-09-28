package services

import (
	"database/sql"
	"errors"

	"book-management/internal/models"
	"book-management/internal/repositories"
	"book-management/internal/utils"
)

type CategoryService struct {
	categoryRepo *repositories.CategoryRepository
}

func NewCategoryService(categoryRepo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	categories, err := s.categoryRepo.GetAll()
	if err != nil {
		return nil, errors.New("failed to get categories")
	}

	return categories, nil
}

func (s *CategoryService) GetCategoryByID(id int) (*models.Category, error) {
	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("failed to get category")
	}

	if category == nil {
		return nil, errors.New("category not found")
	}

	return category, nil
}

func (s *CategoryService) CreateCategory(req *models.CreateCategoryRequest, username string) (*models.Category, error) {
	// Validate input
	if err := utils.ValidateStruct(req); err != nil {
		return nil, errors.New("validation failed: " + err.Error())
	}

	category := &models.Category{
		Name:       req.Name,
		CreatedBy:  username,
		ModifiedBy: username,
	}

	err := s.categoryRepo.Create(category)
	if err != nil {
		return nil, errors.New("failed to create category")
	}

	return category, nil
}

func (s *CategoryService) UpdateCategory(id int, req *models.UpdateCategoryRequest, username string) (*models.Category, error) {
	// Validate input
	if err := utils.ValidateStruct(req); err != nil {
		return nil, errors.New("validation failed: " + err.Error())
	}

	// Check if category exists
	existingCategory, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("failed to get category")
	}

	if existingCategory == nil {
		return nil, errors.New("category not found")
	}

	// Update category
	existingCategory.Name = req.Name
	existingCategory.ModifiedBy = username

	err = s.categoryRepo.Update(existingCategory)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("category not found")
		}
		return nil, errors.New("failed to update category")
	}

	return existingCategory, nil
}

func (s *CategoryService) DeleteCategory(id int) error {
	err := s.categoryRepo.Delete(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("category not found")
		}
		return errors.New("failed to delete category")
	}

	return nil
}

func (s *CategoryService) GetBooksByCategory(categoryID int) ([]models.BookWithCategory, error) {
	// Check if category exists
	category, err := s.categoryRepo.GetByID(categoryID)
	if err != nil {
		return nil, errors.New("failed to get category")
	}

	if category == nil {
		return nil, errors.New("category not found")
	}

	// Get books by category
	books, err := s.categoryRepo.GetBooksByCategory(categoryID)
	if err != nil {
		return nil, errors.New("failed to get books")
	}

	return books, nil
}
