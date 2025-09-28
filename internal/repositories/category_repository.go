package repositories

import (
	"database/sql"
	"time"

	"book-management/internal/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	query := `
		SELECT id, name, created_at, created_by, modified_at, modified_by
		FROM categories 
		ORDER BY id ASC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.CreatedBy,
			&category.ModifiedAt,
			&category.ModifiedBy,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *CategoryRepository) GetByID(id int) (*models.Category, error) {
	query := `
		SELECT id, name, created_at, created_by, modified_at, modified_by
		FROM categories 
		WHERE id = $1
	`

	category := &models.Category{}
	err := r.db.QueryRow(query, id).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.CreatedBy,
		&category.ModifiedAt,
		&category.ModifiedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) Create(category *models.Category) error {
	query := `
		INSERT INTO categories (name, created_by, modified_by)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, modified_at
	`

	err := r.db.QueryRow(
		query,
		category.Name,
		category.CreatedBy,
		category.ModifiedBy,
	).Scan(&category.ID, &category.CreatedAt, &category.ModifiedAt)

	return err
}

func (r *CategoryRepository) Update(category *models.Category) error {
	query := `
		UPDATE categories 
		SET name = $1, modified_by = $2, modified_at = $3
		WHERE id = $4
	`

	category.ModifiedAt = time.Now()
	result, err := r.db.Exec(query, category.Name, category.ModifiedBy, category.ModifiedAt, category.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *CategoryRepository) Delete(id int) error {
	query := `DELETE FROM categories WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *CategoryRepository) GetBooksByCategory(categoryID int) ([]models.BookWithCategory, error) {
	query := `
		SELECT b.id, b.title, b.description, b.image_url, b.release_year, 
			   b.price, b.total_page, b.thickness, b.category_id,
			   b.created_at, b.created_by, b.modified_at, b.modified_by,
			   c.name as category_name
		FROM books b
		JOIN categories c ON b.category_id = c.id
		WHERE b.category_id = $1
		ORDER BY b.id ASC
	`

	rows, err := r.db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.BookWithCategory
	for rows.Next() {
		var book models.BookWithCategory
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageURL,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CategoryID,
			&book.CreatedAt,
			&book.CreatedBy,
			&book.ModifiedAt,
			&book.ModifiedBy,
			&book.CategoryName,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
