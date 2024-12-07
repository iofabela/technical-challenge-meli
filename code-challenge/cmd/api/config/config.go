package config

import (
	"github.com/gin-gonic/gin"
)

func CreateConfig() (*gin.Engine, error) {
	r, err := createRouter()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func createRouter() (*gin.Engine, error) {
	engine, err := NewConfig().setConfig()
	if err != nil {
		return nil, err
	}

	return engine, nil
}
