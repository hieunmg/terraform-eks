package storage

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}
