package env

import "database/sql"

type EnviromentConfig struct {
	DBName string
	SQL    *sql.DB
}
