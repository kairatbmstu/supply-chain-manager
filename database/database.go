package database

import (
	"database/sql"
)

var DB *sql.DB

func ConnectDatabase() (*sql.DB, error) {
	connStr := "postgresql://supplier_portal:123456@localhost/todos?sslmode=disable"
	DB, err := sql.Open("postgres", connStr)
	return DB, err
}
