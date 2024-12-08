package load_file

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/iofabela/technical-challenge-meli/cmd/api/utils/web"
)

type Repository struct {
	db       *sql.DB
	file     *[]byte
	fileName string
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

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) LoadFile(ctx *gin.Context, file *multipart.FileHeader) (*multipart.FileHeader, error) {

	// Open the file in memory
	uploadedFile, err := file.Open()
	if err != nil {
		web.Error(ctx, http.StatusInternalServerError, "It was not possible to open the file")
		return nil, fmt.Errorf("repository.loadFile - %w", err)
	}
	defer uploadedFile.Close()

	var buf bytes.Buffer
	bufReader := io.TeeReader(uploadedFile, &buf)

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
	fileConfig, err := detectFormatAndSeparator(firstLine)
	if err != nil {
		web.Error(ctx, http.StatusInternalServerError, "Error to detect the file format")
		return nil, fmt.Errorf("repository.loadFile - %w", err)
	}

	fmt.Println("Format detected: ", fileConfig.Format)
	fmt.Println("Separator detected: ", string(fileConfig.Separator))

	// Process the file
	if err := r.ProcessFile(ctx, fileConfig, scanner); err != nil {
		return nil, fmt.Errorf("repository.loadFile - %w", err)
	}

	return file, nil
}

// Detect automatically the format and separator using regex
func detectFormatAndSeparator(firstLine string) (FileConfig, error) {

	fmt.Println("Primera línea del archivo: ", firstLine)

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

func (r *Repository) ProcessFile(ctx *gin.Context, config FileConfig, file *bufio.Scanner) error {
	switch config.Format {
	case "csv":
		return r.CSVReader(ctx, file)
	case "jsonlines":
		return r.JSONLinesReader(ctx, file)
	case "txt":
		return r.TXTReader(ctx, file)
	default:
		return fmt.Errorf("NewFileReader - format not supported: %s", config.Format)
	}
}

// Process the file content as CSV
func (r *Repository) CSVReader(ctx *gin.Context, file *bufio.Scanner) error {
	// csvReader := csv.NewReader(*file)
	// rowCount := 1
	// for {
	// 	row, err := csvReader.Read()
	// 	if err != nil {
	// 		if err.Error() == "EOF" { // End of file
	// 			break
	// 		}
	// 		web.Error(ctx, http.StatusBadRequest, "Error al procesar el archivo CSV")
	// 		return fmt.Errorf("repository.loadFile - %w", err)
	// 	}

	// 	// Process each line (here only print)
	// 	fmt.Printf("Row %d: %v\n", rowCount, row)
	// 	rowCount++
	// }

	for file.Scan() {
		line := file.Text()
		if err := Read(line); err != nil && err != io.EOF {
			web.Error(ctx, http.StatusInternalServerError, "Error to read process line")
			return fmt.Errorf("Error to read process line: %w", err)
		}
	}
	return nil
}

func (r *Repository) TXTReader(ctx *gin.Context, file *bufio.Scanner) error {
	// csvReader := csv.NewReader(*file)
	// rowCount := 0
	// for {
	// 	row, err := csvReader.Read()
	// 	if err != nil {
	// 		if err.Error() == "EOF" { // End of file
	// 			break
	// 		}
	// 		web.Error(ctx, http.StatusBadRequest, "Error al procesar el archivo TXT")
	// 		return fmt.Errorf("repository.loadFile - %w", err)
	// 	}

	// 	// Process each line (here only print)
	// 	fmt.Printf("Fila %d: %v\n", rowCount, row)
	// 	rowCount++
	// }

	for file.Scan() {
		line := file.Text()
		if err := Read(line); err != nil && err != io.EOF {
			web.Error(ctx, http.StatusInternalServerError, "Error to read process line")
			return fmt.Errorf("Error to read process line: %w", err)
		}
	}
	return nil
}
func (r *Repository) JSONLinesReader(ctx *gin.Context, file *bufio.Scanner) error {

	for file.Scan() {
		line := file.Text()
		if err := ReadJson(line); err != nil && err != io.EOF {
			web.Error(ctx, http.StatusInternalServerError, "Error to read process line")
			return fmt.Errorf("Error to read process line: %w", err)
		}
	}

	return nil
}

func Read(line string) error {
	fmt.Println("Line:", line)
	return nil
}

func ReadJson(line string) error {
	var obj struct {
		Site string `json:"site"`
		ID   int    `json:"id"`
	}
	err := json.Unmarshal([]byte(line), &obj) // Process the JSON line
	if err != nil {
		return fmt.Errorf("Error to unmarshal JSON: %w", err)
	}
	fmt.Println("JSON Line Object: ", obj)
	return nil
}
