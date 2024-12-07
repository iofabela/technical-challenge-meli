package app

import (
	"github.com/iofabela/technical-challenge-meli/cmd/api/config"
)

func Run() error {
	app, err := config.CreateConfig()
	if err != nil {
		return err
	}
	return app.Run()
}
