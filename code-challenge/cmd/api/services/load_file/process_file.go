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
	"sync"

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
		r.FileConfig.Format = "csv"
		return "csv"
	case txtRegex.MatchString(filename):
		r.FileConfig.Format = "txt"
		return "txt"
	case jsonlRegex.MatchString(filename):
		r.FileConfig.Format = "jsonl"
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

	var wg sync.WaitGroup
	for { // Read the file line by line
		sliceLine, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			web.Error(ctx, http.StatusInternalServerError, "Error to read process line")
			return fmt.Errorf("Error to read process line: %w", err)
		}
		wg.Add(1)
		go func(sliceLine []string) {
			defer wg.Done()
			line := strings.Join(sliceLine, ",")
			if itemFailed := r.ProcessLine(line); itemFailed.Error != nil {
				if r.isCriticalError(itemFailed.Error) {
					web.Error(ctx, http.StatusInternalServerError, "Error to read process line in CSV - %s", line)
					// return fmt.Errorf("processFile.CSV - Internal error to read process line: %s", itemFailed.Error.Error())
				}
				Reprocess(itemFailed)
			}
		}(sliceLine)
	}
	wg.Wait()
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
		return fmt.Errorf("error detecting TXT format: %v", err)
	}

	var wg sync.WaitGroup
	for file.Scan() { // Read the file line by line
		line := file.Text()
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			if itemFailed := r.ProcessLine(line); itemFailed.Error != nil && itemFailed.Error != io.EOF {
				if r.isCriticalError(itemFailed.Error) {
					web.Error(ctx, http.StatusInternalServerError, "Error to read process line in TXT - %s", line)
					// return fmt.Errorf("processFile.TXT - Internal error to read process line: %s", itemFailed.Error.Error())
				}
				Reprocess(itemFailed)
			}
		}(line)
	}
	wg.Wait()
	return nil
}
func (r *Repository) processJSONLiner(ctx *gin.Context, file *bufio.Scanner) error {

	var wg sync.WaitGroup
	for file.Scan() { // Read the file line by line
		line := file.Text()
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			if itemFailed := r.ProcessJson(line); itemFailed.Error != nil && itemFailed.Error != io.EOF {
				if r.isCriticalError(itemFailed.Error) {
					web.Error(ctx, http.StatusInternalServerError, "Error to read process line in JSON")
					// return fmt.Errorf("processFile.JSON - Internal error to read process line: %s", itemFailed.Error.Error())
				}
				Reprocess(itemFailed)
			}
		}(line)
		wg.Wait()
	}
	return nil
}

func (r *Repository) ProcessLine(line string) *items.FailedItem {

	dataLine := strings.Split(line, string(r.FileConfig.Separator))
	item, failedItem := rest.RestMeli_Items(dataLine[0], dataLine[1], r.Client)
	if failedItem.Error != nil {
		return failedItem
	}
	err := r.SqlService.SaveItem(item)
	if err != nil {
		return &items.FailedItem{
			ID:    item.ID,
			Site:  item.SiteID,
			Error: err,
		}
	}
	return &items.FailedItem{}
}

func (r *Repository) ProcessJson(line string) *items.FailedItem {
	var obj items.DataLine
	if err := json.Unmarshal([]byte(line), &obj); err != nil { // Process the JSON line
		return &items.FailedItem{Error: err}
	}

	item, failedItem := rest.RestMeli_Items(obj.Site, strconv.Itoa(obj.ID), r.Client)
	if failedItem.Error != nil {
		return failedItem
	}

	fmt.Println("Item: ", obj.Site, strconv.Itoa(obj.ID), item)
	err := r.SqlService.SaveItem(item)
	if err != nil {
		return &items.FailedItem{
			ID:    item.ID,
			Site:  item.SiteID,
			Error: err,
		}
	}

	return &items.FailedItem{}
}

func (r *Repository) isCriticalError(err error) bool {
	if err == nil {
		return false
	}
	errorMessage := err.Error()
	return strings.Contains(errorMessage, "datatype mismatch") ||
		strings.Contains(errorMessage, "Error to save save item")
}
