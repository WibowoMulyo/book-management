package controllers

import (
	"strconv"

	"book-management/internal/models"
	"book-management/internal/services"
	"book-management/internal/utils"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService *services.BookService
}

func NewBookController(bookService *services.BookService) *BookController {
	return &BookController{
		bookService: bookService,
	}
}

// GetAllBooks godoc
// @Summary Get all books
// @Description Get a list of all books with category information
// @Tags books
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response{data=[]models.BookWithCategory}
// @Failure 401 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/books [get]
func (ctrl *BookController) GetAllBooks(c *gin.Context) {
	books, err := ctrl.bookService.GetAllBooks()
	if err != nil {
		utils.InternalServerError(c, err.Error(), nil)
		return
	}

	utils.OK(c, "Books retrieved successfully", books)
}

// GetBookByID godoc
// @Summary Get book by ID
// @Description Get a specific book by its ID with category information
// @Tags books
// @Produce json
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Success 200 {object} utils.Response{data=models.BookWithCategory}
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/books/{id} [get]
func (ctrl *BookController) GetBookByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequest(c, "Invalid book ID", err.Error())
		return
	}

	book, err := ctrl.bookService.GetBookByID(id)
	if err != nil {
		if err.Error() == "book not found" {
			utils.NotFound(c, "Book not found")
			return
		}
		utils.InternalServerError(c, err.Error(), nil)
		return
	}

	utils.OK(c, "Book retrieved successfully", book)
}

// CreateBook godoc
// @Summary Create new book
// @Description Create a new book
// @Tags books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.CreateBookRequest true "Book data"
// @Success 201 {object} utils.Response{data=models.Book}
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/books [post]
func (ctrl *BookController) CreateBook(c *gin.Context) {
	var req models.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request body", err.Error())
		return
	}

	username := c.GetString("username")
	book, err := ctrl.bookService.CreateBook(&req, username)
	if err != nil {
		if err.Error()[:10] == "validation" {
			errors := utils.FormatValidationErrors(err)
			utils.BadRequest(c, "Validation failed", errors)
			return
		}
		if err.Error() == "category not found" {
			utils.BadRequest(c, "Category not found", nil)
			return
		}
		utils.InternalServerError(c, err.Error(), nil)
		return
	}

	utils.Created(c, "Book created successfully", book)
}

// UpdateBook godoc
// @Summary Update book
// @Description Update an existing book
// @Tags books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Param request body models.UpdateBookRequest true "Book data"
// @Success 200 {object} utils.Response{data=models.Book}
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/books/{id} [put]
func (ctrl *BookController) UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequest(c, "Invalid book ID", err.Error())
		return
	}

	var req models.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request body", err.Error())
		return
	}

	username := c.GetString("username")
	book, err := ctrl.bookService.UpdateBook(id, &req, username)
	if err != nil {
		if err.Error() == "book not found" {
			utils.NotFound(c, "Book not found")
			return
		}
		if err.Error() == "category not found" {
			utils.BadRequest(c, "Category not found", nil)
			return
		}
		if err.Error()[:10] == "validation" {
			errors := utils.FormatValidationErrors(err)
			utils.BadRequest(c, "Validation failed", errors)
			return
		}
		utils.InternalServerError(c, err.Error(), nil)
		return
	}

	utils.OK(c, "Book updated successfully", book)
}

// DeleteBook godoc
// @Summary Delete book
// @Description Delete a book by ID
// @Tags books
// @Produce json
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/books/{id} [delete]
func (ctrl *BookController) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequest(c, "Invalid book ID", err.Error())
		return
	}

	err = ctrl.bookService.DeleteBook(id)
	if err != nil {
		if err.Error() == "book not found" {
			utils.NotFound(c, "Book not found")
			return
		}
		utils.InternalServerError(c, err.Error(), nil)
		return
	}

	utils.OK(c, "Book deleted successfully", nil)
}
