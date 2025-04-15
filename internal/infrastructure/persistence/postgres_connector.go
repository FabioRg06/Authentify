package persistence

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type PostgresConnector struct{}

func (p *PostgresConnector) Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	return sql.Open("postgres", connStr)
}
