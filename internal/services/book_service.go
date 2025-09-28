package services

import (
	"database/sql"
	"errors"

	"book-management/internal/models"
	"book-management/internal/repositories"
	"book-management/internal/utils"
)

type BookService struct {
	bookRepo *repositories.BookRepository
}

func NewBookService(bookRepo *repositories.BookRepository) *BookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

func (s *BookService) GetAllBooks() ([]models.BookWithCategory, error) {
	books, err := s.bookRepo.GetAll()
	if err != nil {
		return nil, errors.New("failed to get books")
	}

	return books, nil
}

func (s *BookService) GetBookByID(id int) (*models.BookWithCategory, error) {
	book, err := s.bookRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("failed to get book")
	}

	if book == nil {
		return nil, errors.New("book not found")
	}

	return book, nil
}

func (s *BookService) CreateBook(req *models.CreateBookRequest, username string) (*models.Book, error) {
	// Validate input
	if err := utils.ValidateStruct(req); err != nil {
		return nil, errors.New("validation failed: " + err.Error())
	}

	// Check if category exists
	categoryExists, err := s.bookRepo.CheckCategoryExists(req.CategoryID)
	if err != nil {
		return nil, errors.New("failed to validate category")
	}

	if !categoryExists {
		return nil, errors.New("category not found")
	}

	book := &models.Book{
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		ReleaseYear: req.ReleaseYear,
		Price:       req.Price,
		TotalPage:   req.TotalPage,
		CategoryID:  req.CategoryID,
		CreatedBy:   username,
		ModifiedBy:  username,
	}

	// Calculate thickness based on total pages
	book.CalculateThickness()

	err = s.bookRepo.Create(book)
	if err != nil {
		return nil, errors.New("failed to create book")
	}

	return book, nil
}

func (s *BookService) UpdateBook(id int, req *models.UpdateBookRequest, username string) (*models.Book, error) {
	// Validate input
	if err := utils.ValidateStruct(req); err != nil {
		return nil, errors.New("validation failed: " + err.Error())
	}

	// Check if book exists
	existingBook, err := s.bookRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("failed to get book")
	}

	if existingBook == nil {
		return nil, errors.New("book not found")
	}

	// Check if category exists
	categoryExists, err := s.bookRepo.CheckCategoryExists(req.CategoryID)
	if err != nil {
		return nil, errors.New("failed to validate category")
	}

	if !categoryExists {
		return nil, errors.New("category not found")
	}

	// Update book
	updatedBook := &models.Book{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		ReleaseYear: req.ReleaseYear,
		Price:       req.Price,
		TotalPage:   req.TotalPage,
		CategoryID:  req.CategoryID,
		ModifiedBy:  username,
		CreatedAt:   existingBook.CreatedAt,
		CreatedBy:   existingBook.CreatedBy,
	}

	// Calculate thickness based on total pages
	updatedBook.CalculateThickness()

	err = s.bookRepo.Update(updatedBook)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("book not found")
		}
		return nil, errors.New("failed to update book")
	}

	return updatedBook, nil
}

func (s *BookService) DeleteBook(id int) error {
	err := s.bookRepo.Delete(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("book not found")
		}
		return errors.New("failed to delete book")
	}

	return nil
}
