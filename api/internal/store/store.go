package store

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"fmt"
)

type Store struct {
	db *sqlx.DB
}

func New(cfg DatabaseConfig) (*Store, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
	)
	
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	
	return &Store{db: db}, nil
} 