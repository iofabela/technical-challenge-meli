package internal_sql

import (
	"database/sql"
	"fmt"
	"sync"

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

var dbMutex sync.Mutex

func Connect(database string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./"+database)
	if err != nil {
		return nil, err
	}
	createTable := `
	CREATE TABLE IF NOT EXISTS items (
		id STRING NOT NULL,
		site TEXT NOT NULL,
		price LONG NOT NULL,
		start_time TEXT NOT NULL,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		nickname TEXT NOT NULL
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		return nil, fmt.Errorf("sql.Connect - Error to create table: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("sql.Connect - Error to ping database: %s", err.Error())
	}

	return db, err
}

func (s *SQL) SaveItem(item *items.SaveItem) error {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	_, err := s.DB.
		Exec("INSERT INTO items ( id, site, price, start_time, name, description, nickname) VALUES (?, ?, ?, ?, ?, ?, ?)",
			item.ID, item.SiteID, item.Price, item.StartTime, item.Name, item.Description, item.Nickname)
	if err != nil {
		return fmt.Errorf("sql.SaveItem - Error to save item: %s", err.Error())
	}
	return nil
}
