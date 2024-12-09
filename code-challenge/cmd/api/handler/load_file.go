package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iofabela/technical-challenge-meli/cmd/api/services/load_file"
	"github.com/iofabela/technical-challenge-meli/cmd/api/utils/web"
)

// LoadFile struct for services
type LoadFile struct {
	loadFileService load_file.Service
}

// NewLoadFile Handler
func NewLoadFile(l load_file.Service) *LoadFile {
	return &LoadFile{
		loadFileService: l,
	}
}

// GetLoadData …
// @Summary Load File by a request form and save it in the database SQLite
// @Description Load File by a request form and save it in the database SQLite
// @Tags LoadFile
// @Accept multipart/form-data
// @Produce json
// @Param file formData file  true "File to be loaded"
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /local/api/load_file [post]
func (l *LoadFile) LoadData() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("Path: %s, Method: %s", ctx.Request.URL.Path, ctx.Request.Method)
		// Get the file from the form
		file, err := ctx.FormFile("file")
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "It was not possible to get the file")
			return
		}
		// Load the file in the database
		_, err = l.loadFileService.LoadFile(ctx, file)
		if err != nil {
			fmt.Printf("handler.%s - %v\n", ctx.Request.URL.Path, err)
			return
		}

		fmt.Println("LoadFile successfully | Status 200 OK")
		web.Success(ctx, http.StatusOK, gin.H{"message": "Archive successfully processed"})
	}
}
