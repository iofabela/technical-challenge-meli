package routes

import (
	"net/http"

	_ "github.com/iofabela/technical-challenge-meli/docs/guide"

	"github.com/gin-gonic/gin"
	"github.com/iofabela/technical-challenge-meli/cmd/api/handler"
	"github.com/iofabela/technical-challenge-meli/cmd/api/services/load_file"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (r *Router) setGroup() {
	r.rg = r.r.Group(r.config.Scope + "/api/")
}

// Build ping of the App
func (r *Router) buildPing() {
	r.r.GET("/ping", func(c *gin.Context) { c.Writer.WriteHeader(http.StatusOK); c.Next() })
}

func (r *Router) buildLoadFileRoute() {

	repo := load_file.NewRepository(r.config.EnvConfig.SQL, r.config.Rest, r.config.EnvConfig.SqlService)
	service := load_file.NewService(repo)
	handler := handler.NewLoadFile(service)
	loadFileRoute := r.rg.Group("/")
	{
		loadFileRoute.POST("/load_file", handler.LoadData())
	}
}

func (r *Router) buildSwaggerRoutes() {
	// docs.SwaggerInfo.Host = "localhost:8080"
	// docs.SwaggerInfo.BasePath = "/"
	r.r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
