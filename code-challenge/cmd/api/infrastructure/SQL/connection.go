package internal_sql

import (
	"database/sql"

	"github.com/iofabela/technical-challenge-meli/cmd/api/models/items"
	_ "github.com/mattn/go-sqlite3"
)

type SQL struct {
	DB *sql.DB
}

func NewSQL(db *sql.DB) *SQL {
	return &SQL{
		DB: db,
	}
}

func Connect(database string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./"+database)
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	return db, err
}

func (s *SQL) SaveItem(item *items.SaveItem) error {
	_, err := s.DB.Exec("INSERT INTO items (site_id, id, start_time, name, description, nickname) VALUES (?, ?, ?, ?, ?, ?)", item.SiteID, item.ID, item.StartTime, item.Name, item.Description, item.Nickname)
	if err != nil {
		return err
	}
	return nil
}
