-- +migrate Up
CREATE TABLE categories (
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(255) NOT NULL,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            created_by VARCHAR(255) DEFAULT 'system',
                            modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            modified_by VARCHAR(255) DEFAULT 'system'
);

-- Insert sample categories
INSERT INTO categories (name, created_by, modified_by) VALUES
                                                           ('Fiction', 'system', 'system'),
                                                           ('Non-Fiction', 'system', 'system'),
                                                           ('Science', 'system', 'system'),
                                                           ('Technology', 'system', 'system'),
                                                           ('History', 'system', 'system');

-- +migrate Down
DROP TABLE categories;