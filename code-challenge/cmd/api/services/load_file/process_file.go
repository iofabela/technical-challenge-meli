package load_file

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iofabela/technical-challenge-meli/cmd/api/infrastructure/rest"
	"github.com/iofabela/technical-challenge-meli/cmd/api/models/items"
	"github.com/iofabela/technical-challenge-meli/cmd/api/utils/web"
)

func (r *Repository) DetectFileType(filename string) string {
	fmt.Println("Filename: ", filename)
	fmt.Println("FileType: ", filename[len(filename)-4:]) // TODO remove

	// Convertir el nombre del archivo a minúsculas para evitar problemas de sensibilidad
	filename = strings.ToLower(filename)

	// Definir expresiones regulares para cada tipo de archivo
	csvRegex := regexp.MustCompile(`\.csv$`)
	txtRegex := regexp.MustCompile(`\.txt$`)
	jsonlRegex := regexp.MustCompile(`\.jsonl$`)

	// Validar con las expresiones regulares
	switch {
	case csvRegex.MatchString(filename):
		return "csv"
	case txtRegex.MatchString(filename):
		return "txt"
	case jsonlRegex.MatchString(filename):
		return "jsonl"
	default:
		return "unknown"
	}
}

func (r *Repository) ProcessFile(ctx *gin.Context, uploadedFile *multipart.File, fileType string, scanner *bufio.Scanner) error {
	switch fileType {
	case "csv":
		return r.processCSV(ctx, *uploadedFile)
	case "jsonl":
		return r.processJSONLiner(ctx, scanner)
	case "txt":
		return r.processTXT(ctx, scanner)
	default:
		return fmt.Errorf("NewFileReader - format not supported: %s", fileType)
	}
}

// Process the file content as CSV
func (r *Repository) processCSV(ctx *gin.Context, file multipart.File) error {
	// Open the file in memory
	reader := csv.NewReader(file)
	// Read the first line of the CSV file (header)
	firstLine, err := reader.Read()
	if err != nil {
		return fmt.Errorf("error reading CSV header: %v", err)
	}
	// Detect the format and separator
	_, err = r.DetectFormatAndSeparator(strings.Join(firstLine, ","))
	if err != nil {
		return fmt.Errorf("error detecting CSV format: %v", err)
	}

	for { // Read the file line by line
		sliceLine, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			web.Error(ctx, http.StatusInternalServerError, "Error to read process line")
			return fmt.Errorf("Error to read process line: %w", err) // TODO save the line with error in a file
		}
		line := strings.Join(sliceLine, ",")

		if err = r.ProcessLine(line); err != nil { // TODO save the line with error in a file
			web.Error(ctx, http.StatusInternalServerError, "Error to read process line in CSV")
			return fmt.Errorf("processFile.CSV - Internal error to read process line: %s", err.Error())
		}
	}
	return nil
}

func (r *Repository) processTXT(ctx *gin.Context, file *bufio.Scanner) error {

	var firstLine string
	if file.Scan() { // If the first line exists
		firstLine = file.Text()
	}
	// Detect the format and separator
	_, err := r.DetectFormatAndSeparator(firstLine)
	if err != nil {
		return fmt.Errorf("error detecting CSV format: %v", err)
	}

	for file.Scan() { // Read the file line by line
		line := file.Text()
		if err := r.ProcessLine(line); err != nil && err != io.EOF {
			web.Error(ctx, http.StatusInternalServerError, "Error to read process line in TXT")
			return fmt.Errorf("processFile.TXT - Internal error to read process line: %s", err.Error())
		}
	}
	return nil
}
func (r *Repository) processJSONLiner(ctx *gin.Context, file *bufio.Scanner) error {

	for file.Scan() { // Read the file line by line
		line := file.Text()
		if err := r.ProcessJson(line); err != nil && err != io.EOF {
			web.Error(ctx, http.StatusInternalServerError, "Error to read process line in JSON")
			return fmt.Errorf("processFile.JSON - Internal error to read process line: %s", err.Error())
		}
	}

	return nil
}

func (r *Repository) ProcessLine(line string) error {
	fmt.Println("Line:", line) // TODO remove

	dataLine := strings.Split(line, string(r.FileConfig.Separator))
	item, err := rest.RestMeli_Items(dataLine[0], dataLine[1], r.Client) // TODO get fields of site and id
	if err != nil {
		//TODO save the line with error in a file
		return err
	}
	fmt.Println("Item: ", item) // TODO remove
	err = r.SqlService.SaveItem(item)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) ProcessJson(line string) error {
	var obj items.DataLine
	if err := json.Unmarshal([]byte(line), &obj); err != nil { // Process the JSON line
		return fmt.Errorf("Error to unmarshal JSON: %s", err.Error())
	}
	fmt.Println("JSON Line Object: ", obj) // TODO remove
	item, err := rest.RestMeli_Items(obj.Site, strconv.Itoa(obj.ID), r.Client)
	if err != nil {
		//TODO save the line with error in a file
		return err
	}

	fmt.Println("Item: ", obj.Site, strconv.Itoa(obj.ID), item)
	err = r.SqlService.SaveItem(item)
	if err != nil {
		return err
	}
	return nil
}

func Reprocess(id string) error { // Reprocess the item if fail
	// Save in a file
	return nil
}
