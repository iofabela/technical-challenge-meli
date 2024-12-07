package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iofabela/technical-challenge-meli/cmd/api/handler"
	"github.com/iofabela/technical-challenge-meli/cmd/api/services/load_file"
)

func (r *Router) setGroup() {
	r.rg = r.r.Group(r.config.Scope + "/api/")
}

// Build ping of the App
func (r *Router) buildPing() {
	r.r.GET("/ping", func(c *gin.Context) { c.Writer.WriteHeader(http.StatusOK); c.Next() })
}

func (r *Router) buildLoadFileRoute() {

	repo := load_file.NewRepository(r.config.EnvConfig.SQL)
	service := load_file.NewService(repo)
	handler := handler.NewLoadFile(service)
	loadFileRoute := r.rg.Group("/")
	{
		loadFileRoute.POST("/load_file", handler.GetLoadData())
	}
}
