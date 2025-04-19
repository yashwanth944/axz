package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/user/apigateway/auth"
	"github.com/user/apigateway/models"
)

// InitDB establishes a connection to the PostgreSQL database
func InitDB() (*sqlx.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "apigateway"
	}

	// Construct connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	// Connect to database
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Create the users table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreateUser creates a new user in the database
func CreateUser(db *sqlx.DB, user *models.User) error {
	// Check if user already exists
	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM users WHERE email = $1", user.Email)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// Insert user into database
	query := `
	INSERT INTO users (email, password, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at
	`

	row := db.QueryRowx(query, user.Email, hashedPassword, user.FirstName, user.LastName)
	
	return row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

// GetUserByEmail retrieves a user by their email address
func GetUserByEmail(db *sqlx.DB, email string) (*models.User, error) {
	var user models.User
	
	query := "SELECT * FROM users WHERE email = $1"
	err := db.Get(&user, query, email)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	
	return &user, nil
}

// GetUserByID retrieves a user by their ID
func GetUserByID(db *sqlx.DB, id uint) (*models.User, error) {
	var user models.User
	
	query := "SELECT * FROM users WHERE id = $1"
	err := db.Get(&user, query, id)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	
	return &user, nil
} 