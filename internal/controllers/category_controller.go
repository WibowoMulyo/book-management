package controllers

import (
	"strconv"

	"book-management/internal/models"
	"book-management/internal/services"
	"book-management/internal/utils"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Get a list of all categories
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response{data=[]models.Category}
// @Failure 401 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/categories [get]
func (ctrl *CategoryController) GetAllCategories(c *gin.Context) {
	categories, err := ctrl.categoryService.GetAllCategories()
	if err != nil {
		utils.InternalServerError(c, err.Error(), nil)
		return
	}

	utils.OK(c, "Categories retrieved successfully", categories)
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Description Get a specific category by its ID
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} utils.Response{data=models.Category}
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/categories/{id} [get]
func (ctrl *CategoryController) GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequest(c, "Invalid category ID", err.Error())
		return
	}

	category, err := ctrl.categoryService.GetCategoryByID(id)
	if err != nil {
		if err.Error() == "category not found" {
			utils.NotFound(c, "Category not found")
			return
		}
		utils.InternalServerError(c, err.Error(), nil)
		return
	}

	utils.OK(c, "Category retrieved successfully", category)
}

// CreateCategory godoc
// @Summary Create new category
// @Description Create a new category
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.CreateCategoryRequest true "Category data"
// @Success 201 {object} utils.Response{data=models.Category}
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/categories [post]
func (ctrl *CategoryController) CreateCategory(c *gin.Context) {
	var req models.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request body", err.Error())
		return
	}

	username := c.GetString("username")
	category, err := ctrl.categoryService.CreateCategory(&req, username)
	if err != nil {
		if err.Error()[:10] == "validation" {
			errors := utils.FormatValidationErrors(err)
			utils.BadRequest(c, "Validation failed", errors)
			return
		}
		utils.InternalServerError(c, err.Error(), nil)
		return
	}

	utils.Created(c, "Category created successfully", category)
}

// UpdateCategory godoc
// @Summary Update category
// @Description Update an existing category
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Param request body models.UpdateCategoryRequest true "Category data"
// @Success 200 {object} utils.Response{data=models.Category}
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/categories/{id} [put]
func (ctrl *CategoryController) UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequest(c, "Invalid category ID", err.Error())
		return
	}

	var req models.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request body", err.Error())
		return
	}

	username := c.GetString("username")
	category, err := ctrl.categoryService.UpdateCategory(id, &req, username)
	if err != nil {
		if err.Error() == "category not found" {
			utils.NotFound(c, "Category not found")
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

	utils.OK(c, "Category updated successfully", category)
}

// DeleteCategory godoc
// @Summary Delete category
// @Description Delete a category by ID
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/categories/{id} [delete]
func (ctrl *CategoryController) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequest(c, "Invalid category ID", err.Error())
		return
	}

	err = ctrl.categoryService.DeleteCategory(id)
	if err != nil {
		if err.Error() == "category not found" {
			utils.NotFound(c, "Category not found")
			return
		}
		utils.InternalServerError(c, err.Error(), nil)
		return
	}

	utils.OK(c, "Category deleted successfully", nil)
}

// GetBooksByCategory godoc
// @Summary Get books by category
// @Description Get all books in a specific category
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} utils.Response{data=[]models.BookWithCategory}
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/categories/{id}/books [get]
func (ctrl *CategoryController) GetBooksByCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequest(c, "Invalid category ID", err.Error())
		return
	}

	books, err := ctrl.categoryService.GetBooksByCategory(id)
	if err != nil {
		if err.Error() == "category not found" {
			utils.NotFound(c, "Category not found")
			return
		}
		utils.InternalServerError(c, err.Error(), nil)
		return
	}

	utils.OK(c, "Books retrieved successfully", books)
}
