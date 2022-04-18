package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func NewDB() (*sql.DB, error) {
	connStr := "postgresql://supply_chain_manager:123456@localhost/supply_chain_manager?sslmode=disable"
	DB, err := sql.Open("postgres", connStr)
	return DB, err
}
