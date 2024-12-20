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

	// Get the scanner
	scanner := bufio.NewScanner(uploadedFile)

	// Detect the file type
	fileType := r.DetectFileType(file.Filename)
	if fileType == "unknown" {
		web.Error(ctx, http.StatusInternalServerError, "It was not possible to detect the file type")
		return nil, fmt.Errorf("repository.loadFile - %s", fileType)
	}

	// Process the file
	if err := r.ProcessFile(ctx, &uploadedFile, fileType, scanner); err != nil {
		return nil, fmt.Errorf("repository.loadFile - %w", err)
	}
	fmt.Println("Total failed items: ", len(ToReprocess))
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

	// Regex for JSON Lines (looks for common separators: { } , ; | \t)
	jsonRegex := regexp.MustCompile(`^\s*{\s*.*\s*}$`)

	// Regex for CSV (looks for common separators: , ; | \t)
	csvRegex := regexp.MustCompile(`^.*([,;|\t]).*`)

	// Regex for TXT (looks for common separators: , ; | \t)
	txtRegex := regexp.MustCompile(`^.*([,;|\t]).*`)

	// Format detection
	if jsonRegex.MatchString(firstLine) {
		r.FileConfig = FileConfig{Format: "jsonlines", Separator: 0}
		fmt.Print("JSON Separator: {}")
		return FileConfig{Format: "jsonlines", Separator: 0}, nil
	} else if csvRegex.MatchString(firstLine) || r.FileConfig.Format == "csv" {
		// Detect separator most likely in CSV
		match := csvRegex.FindStringSubmatch(firstLine)
		if len(match) > 1 {
			separator := rune(match[1][0])
			r.FileConfig = FileConfig{Format: "csv", Separator: separator}
			fmt.Println("CSV Separator: ", string(separator))
			return FileConfig{Format: "csv", Separator: separator}, nil
		}
	} else if txtRegex.MatchString(firstLine) || r.FileConfig.Format == "txt" {
		match := txtRegex.FindStringSubmatch(firstLine)
		if len(match) > 1 {
			separator := rune(match[1][0])
			r.FileConfig = FileConfig{Format: "txt", Separator: separator}
			fmt.Println("TXT Separator: ", string(separator))
			return FileConfig{Format: "txt", Separator: separator}, nil
		}
	}

	return FileConfig{}, fmt.Errorf("detectFormatAndSeparator - It was not possible to determine the file format")
}
