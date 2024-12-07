package sql

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Connect(DB_NAME string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./"+DB_NAME)
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	return db, err
}
