package main

import (
	"log"

	"book-management/internal/config"
	"book-management/internal/controllers"
	"book-management/internal/middleware"
	"book-management/internal/repositories"
	"book-management/internal/services"
	"book-management/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	defer cfg.DB.Close()

	// Initialize JWT manager
	jwtManager := utils.NewJWTManager(cfg.JWTSecret, cfg.JWTExpire)

	// Initialize repositories
	userRepo := repositories.NewUserRepository(cfg.DB)
	categoryRepo := repositories.NewCategoryRepository(cfg.DB)
	bookRepo := repositories.NewBookRepository(cfg.DB)

	// Initialize services
	authService := services.NewAuthService(userRepo, jwtManager)
	categoryService := services.NewCategoryService(categoryRepo)
	bookService := services.NewBookService(bookRepo)

	// Initialize controllers
	authController := controllers.NewAuthController(authService)
	categoryController := controllers.NewCategoryController(categoryService)
	bookController := controllers.NewBookController(bookService)

	// Initialize Gin router
	router := gin.Default()

	// Add middleware
	router.Use(middleware.CORSMiddleware())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		utils.OK(c, "Server is running", map[string]string{
			"status":  "healthy",
			"service": "book-management-api",
		})
	})

	// API routes
	api := router.Group("/api")
	{
		// Auth routes (no authentication required)
		users := api.Group("/users")
		{
			users.POST("/login", authController.Login)
		}

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.JWTAuthMiddleware(authService))
		{
			// Categories routes
			categories := protected.Group("/categories")
			{
				categories.GET("", categoryController.GetAllCategories)
				categories.POST("", categoryController.CreateCategory)
				categories.GET("/:id", categoryController.GetCategoryByID)
				categories.PUT("/:id", categoryController.UpdateCategory)
				categories.DELETE("/:id", categoryController.DeleteCategory)
				categories.GET("/:id/books", categoryController.GetBooksByCategory)
			}

			// Books routes
			books := protected.Group("/books")
			{
				books.GET("", bookController.GetAllBooks)
				books.POST("", bookController.CreateBook)
				books.GET("/:id", bookController.GetBookByID)
				books.PUT("/:id", bookController.UpdateBook)
				books.DELETE("/:id", bookController.DeleteBook)
			}
		}
	}

	log.Printf("Server starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
