package repositories

import (
	"database/sql"
	"time"

	"book-management/internal/models"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetAll() ([]models.BookWithCategory, error) {
	query := `
		SELECT b.id, b.title, b.description, b.image_url, b.release_year, 
			   b.price, b.total_page, b.thickness, b.category_id,
			   b.created_at, b.created_by, b.modified_at, b.modified_by,
			   c.name as category_name
		FROM books b
		JOIN categories c ON b.category_id = c.id
		ORDER BY b.id ASC
	`

	rows, err := r.db.Query(query)
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

func (r *BookRepository) GetByID(id int) (*models.BookWithCategory, error) {
	query := `
		SELECT b.id, b.title, b.description, b.image_url, b.release_year, 
			   b.price, b.total_page, b.thickness, b.category_id,
			   b.created_at, b.created_by, b.modified_at, b.modified_by,
			   c.name as category_name
		FROM books b
		JOIN categories c ON b.category_id = c.id
		WHERE b.id = $1
	`

	book := &models.BookWithCategory{}
	err := r.db.QueryRow(query, id).Scan(
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
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return book, nil
}

func (r *BookRepository) Create(book *models.Book) error {
	query := `
		INSERT INTO books (title, description, image_url, release_year, price, 
						  total_page, thickness, category_id, created_by, modified_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, created_at, modified_at
	`

	err := r.db.QueryRow(
		query,
		book.Title,
		book.Description,
		book.ImageURL,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CategoryID,
		book.CreatedBy,
		book.ModifiedBy,
	).Scan(&book.ID, &book.CreatedAt, &book.ModifiedAt)

	return err
}

func (r *BookRepository) Update(book *models.Book) error {
	query := `
		UPDATE books 
		SET title = $1, description = $2, image_url = $3, release_year = $4,
			price = $5, total_page = $6, thickness = $7, category_id = $8,
			modified_by = $9, modified_at = $10
		WHERE id = $11
	`

	book.ModifiedAt = time.Now()
	result, err := r.db.Exec(
		query,
		book.Title,
		book.Description,
		book.ImageURL,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CategoryID,
		book.ModifiedBy,
		book.ModifiedAt,
		book.ID,
	)
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

func (r *BookRepository) Delete(id int) error {
	query := `DELETE FROM books WHERE id = $1`

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

func (r *BookRepository) CheckCategoryExists(categoryID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM categories WHERE id = $1)`

	var exists bool
	err := r.db.QueryRow(query, categoryID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
