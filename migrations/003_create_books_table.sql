-- +migrate Up
CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       description TEXT,
                       image_url VARCHAR(255),
                       release_year INTEGER CHECK (release_year >= 1980 AND release_year <= 2024),
                       price INTEGER,
                       total_page INTEGER,
                       thickness VARCHAR(10) CHECK (thickness IN ('tipis', 'tebal')),
                       category_id INTEGER NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       created_by VARCHAR(255) DEFAULT 'system',
                       modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       modified_by VARCHAR(255) DEFAULT 'system',
                       FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE RESTRICT
);

-- Create index for better performance
CREATE INDEX idx_books_category_id ON books(category_id);
CREATE INDEX idx_books_release_year ON books(release_year);

-- Insert sample books
INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by, modified_by) VALUES
                                                                                                                                        ('The Go Programming Language', 'Comprehensive guide to Go programming', 'https://example.com/go-book.jpg', 2015, 500000, 380, 'tebal', 4, 'system', 'system'),
                                                                                                                                        ('Clean Code', 'A handbook of agile software craftsmanship', 'https://example.com/clean-code.jpg', 2008, 450000, 464, 'tebal', 4, 'system', 'system'),
                                                                                                                                        ('The Pragmatic Programmer', 'Your journey to mastery', 'https://example.com/pragmatic.jpg', 2019, 400000, 352, 'tebal', 4, 'system', 'system');

-- +migrate Down
DROP INDEX IF EXISTS idx_books_release_year;
DROP INDEX IF EXISTS idx_books_category_id;
DROP TABLE books;