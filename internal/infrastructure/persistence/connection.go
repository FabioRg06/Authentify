package persistence

import "database/sql"

type DBConnector interface {
	Connect() (*sql.DB, error)
}
