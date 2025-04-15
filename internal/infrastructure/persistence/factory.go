package persistence

import (
	"log"
	"os"
)

func GetConnector() DBConnector {
	switch os.Getenv("DB_TYPE") {
	case "postgres":
		return &PostgresConnector{}

	default:
		log.Fatal("Unsupported DB_TYPE or not set")
		return nil
	}
}
