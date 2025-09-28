package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

type Config struct {
	DB        *sql.DB
	Port      string
	JWTSecret string
	JWTExpire int
}

func LoadConfig() (*Config, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Database configuration
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "book_management")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")

	// JWT configuration
	jwtSecret := getEnv("JWT_SECRET", "your_super_secret_jwt_key_here")
	jwtExpireStr := getEnv("JWT_EXPIRE_HOURS", "24")
	jwtExpire, err := strconv.Atoi(jwtExpireStr)
	if err != nil {
		jwtExpire = 24
	}

	// Server configuration
	port := getEnv("PORT", "8080")

	// Database connection
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Run migrations
	if err := runMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}

	return &Config{
		DB:        db,
		Port:      port,
		JWTSecret: jwtSecret,
		JWTExpire: jwtExpire,
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func runMigrations(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	log.Printf("Applied %d migrations", n)
	return nil
}
