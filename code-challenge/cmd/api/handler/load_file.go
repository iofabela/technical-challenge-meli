package handler

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iofabela/technical-challenge-meli/cmd/api/services/load_file"
	"github.com/iofabela/technical-challenge-meli/cmd/api/utils/web"
)

// LoadFile struct for services
type LoadFile struct {
	loadFileService load_file.Service
}

type file struct {
	FileContent *multipart.FileHeader `binding:"required"`
}

// NewLoadFile Handler
func NewLoadFile(l load_file.Service) *LoadFile {
	return &LoadFile{
		loadFileService: l,
	}
}

func (l *LoadFile) GetLoadData() gin.HandlerFunc {
	return func(ctx *gin.Context) {

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
