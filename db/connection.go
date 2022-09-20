package db

import (
	"github.com/jmoiron/sqlx"
	"os"
)

// ConnectDB connects to the database
func ConnectDB() *sqlx.DB {
	return sqlx.MustConnect("postgres", os.Getenv("DB_CONNECTION"))
}
