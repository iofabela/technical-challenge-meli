package load_file

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iofabela/technical-challenge-meli/cmd/api/utils/web"
)

type Repository struct {
	db       *sql.DB
	file     *[]byte
	fileName string
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) LoadFile(ctx *gin.Context, file *multipart.FileHeader) (*multipart.FileHeader, error) {

	// Abre el archivo en memoria
	uploadedFile, err := file.Open()
	if err != nil {
		web.Error(ctx, http.StatusInternalServerError, "No se pudo abrir el archivo")
		return nil, fmt.Errorf("repository.loadFile - %w", err)
	}
	defer uploadedFile.Close()

	// Lee el contenido del archivo en memoria
	buf := new(strings.Builder)
	_, err = io.Copy(buf, uploadedFile)
	if err != nil {
		web.Error(ctx, http.StatusInternalServerError, "Error al leer el archivo")
		return nil, fmt.Errorf("repository.loadFile - %w", err)
	}

	// Procesa el contenido del archivo como CSV
	csvReader := csv.NewReader(strings.NewReader(buf.String()))
	rowCount := 0
	for {
		row, err := csvReader.Read()
		if err != nil {
			if err.Error() == "EOF" { // Fin del archivo
				break
			}
			web.Error(ctx, http.StatusBadRequest, "Error al procesar el archivo CSV")
			return nil, fmt.Errorf("repository.loadFile - %w", err)
		}

		// Procesa cada línea (aquí solo la imprime)
		fmt.Printf("Fila %d: %v\n", rowCount, row)
		rowCount++
	}

	return file, nil
}
