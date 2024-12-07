package config

import (
	"github.com/gin-gonic/gin"
	sql "github.com/iofabela/technical-challenge-meli/cmd/api/infrastructure/SQL"
	"github.com/iofabela/technical-challenge-meli/cmd/api/models/env"
	"github.com/iofabela/technical-challenge-meli/cmd/api/routes"
)

type Config struct {
	Port      string
	GinMode   string
	Scope     string
	EnvConfig env.EnviromentConfig
}

func NewConfig() *Config {
	return &Config{
		Port:    "8080",
		GinMode: "debug",
		Scope:   "local",
		EnvConfig: env.EnviromentConfig{
			DBName: "app.db",
		},
	}
}
func (cfg *Config) getConfig() *routes.RouterConfig {
	return (*routes.RouterConfig)(cfg)
}

func (cfg *Config) setConfig() (*gin.Engine, error) {
	var err error
	cfg.EnvConfig.SQL, err = sql.Connect(cfg.EnvConfig.DBName)
	if err != nil {
		panic(err)
	}

	// Engine instance - router
	gin.SetMode(gin.DebugMode)

	// Routes config
	router := gin.Default()
	routes.NewRouter(router, cfg.getConfig()).MapRoutes() // Routes config

	return router, nil
}
