package env

import (
	"database/sql"

	internal_sql "github.com/iofabela/technical-challenge-meli/cmd/api/infrastructure/SQL"
)

type EnviromentConfig struct {
	DBName     string
	SqlService *internal_sql.SQL
	SQL        *sql.DB
}
