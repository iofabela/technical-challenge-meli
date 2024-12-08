package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iofabela/technical-challenge-meli/cmd/api/infrastructure/rest"
	"github.com/iofabela/technical-challenge-meli/cmd/api/models/env"
)

type Router struct {
	r      *gin.Engine
	rg     *gin.RouterGroup
	config *RouterConfig
}

type RouterConfig struct {
	Port      string
	GinMode   string
	Scope     string
	EnvConfig env.EnviromentConfig
	Rest      *rest.Client
}

func NewRouter(r *gin.Engine, config *RouterConfig) *Router {
	return &Router{
		r:      r,
		rg:     &r.RouterGroup,
		config: config,
	}
}

func (r *Router) MapRoutes() {
	// Path begining
	r.setGroup()
	// Build routes
	r.buildPing()
	r.buildLoadFileRoute()
}
