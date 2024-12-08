package load_file

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	internal_sql "github.com/iofabela/technical-challenge-meli/cmd/api/infrastructure/SQL"
	"github.com/iofabela/technical-challenge-meli/cmd/api/infrastructure/rest"
	"github.com/iofabela/technical-challenge-meli/cmd/api/utils/web"
)

type Repository struct {
	db         *sql.DB
	SqlService *internal_sql.SQL
	file       *[]byte
	fileName   string
	FileConfig FileConfig
	Client     *rest.Client
}

// Configuration of the reader
type FileConfig struct {
	Format    string // Format: csv, jsonlines, txt
	Separator rune   // Separator detected (para CSV)
}

// Implementación para CSV
type CSVReader struct {
	Separator rune
	Reader    *csv.Reader
}

func NewRepository(db *sql.DB, client *rest.Client, sqlService *internal_sql.SQL) *Repository {
	return &Repository{
		db:         db,
		Client:     client,
		SqlService: sqlService,
	}
}

func (r *Repository) LoadFile(ctx *gin.Context, file *multipart.FileHeader) (*multipart.FileHeader, error) {

	// Open the file in memory
	uploadedFile, err := file.Open()
	if err != nil {
		web.Error(ctx, http.StatusInternalServerError, "It was not possible to open the file")
		return nil, fmt.Errorf("repository.loadFile - %w", err)
	}
	defer uploadedFile.Close()

	// Get the config file
	scanner, err := r.GetConfigFile(ctx, &uploadedFile)
	if err != nil {
		return nil, fmt.Errorf("repository.loadFile - %w", err)
	}

	// Process the file
	if err := r.ProcessFile(ctx, scanner); err != nil {
		return nil, fmt.Errorf("repository.loadFile - %w", err)
	}

	return file, nil
}

func (r *Repository) GetConfigFile(ctx *gin.Context, uploadedFile *multipart.File) (*bufio.Scanner, error) {
	var buf bytes.Buffer
	bufReader := io.TeeReader(*uploadedFile, &buf)

	// Read the first line of the file
	var firstLine string
	scanner := bufio.NewScanner(bufReader)
	if scanner.Scan() { // If the first line exists
		firstLine = scanner.Text()
	}
	// Check for errors while scanning
	if err := scanner.Err(); err != nil {
		web.Error(ctx, http.StatusInternalServerError, "Error to read the file")
		return nil, fmt.Errorf("detectFormatAndSeparator - %w", err)
	}

	// Detect the format and separator
	fileConfig, err := r.DetectFormatAndSeparator(firstLine)
	if err != nil {
		web.Error(ctx, http.StatusInternalServerError, "Error to detect the file format")
		return nil, fmt.Errorf("repository.loadFile - %w", err)
	}

	fmt.Println("Format detected: ", fileConfig.Format)               // TODO remove
	fmt.Println("Separator detected: ", string(fileConfig.Separator)) // TODO remove
	r.FileConfig = fileConfig
	return scanner, nil
}

// Detect automatically the format and separator using regex
func (r *Repository) DetectFormatAndSeparator(firstLine string) (FileConfig, error) {

	fmt.Println("First line of the file: ", firstLine) // TODO remove

	// Regex for JSON Lines (looks for common separators: { } , ; | \t)
	jsonRegex := regexp.MustCompile(`^\s*{\s*.*\s*}$`)

	// Regex for CSV (looks for common separators: , ; | \t)
	csvRegex := regexp.MustCompile(`^.*([,;|\t]).*`)

	// Format detection
	if jsonRegex.MatchString(firstLine) {
		return FileConfig{Format: "jsonlines", Separator: 0}, nil
	} else if csvRegex.MatchString(firstLine) {
		// Detect separator most likely in CSV
		match := csvRegex.FindStringSubmatch(firstLine)
		if len(match) > 1 {
			separator := rune(match[1][0])
			return FileConfig{Format: "csv", Separator: separator}, nil
		}
	} else {
		return FileConfig{Format: "txt", Separator: 0}, nil
	}

	return FileConfig{}, fmt.Errorf("detectFormatAndSeparator - It was not possible to determine the file format")
}
