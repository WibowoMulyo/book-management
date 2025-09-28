-- +migrate Up
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       created_by VARCHAR(255) DEFAULT 'system',
                       modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       modified_by VARCHAR(255) DEFAULT 'system'
);

-- Insert default admin user (password: admin123)
INSERT INTO users (username, password, created_by, modified_by)
VALUES ('admin', '$2a$10$ILAGgLASFiUrPpOy70E8MOHsLcznfsEQo8YXme1OeRR7.kO6Uzera', 'system', 'system');

-- +migrate Down
DROP TABLE users;