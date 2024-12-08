package load_file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iofabela/technical-challenge-meli/cmd/api/infrastructure/rest"
	"github.com/iofabela/technical-challenge-meli/cmd/api/models/items"
	"github.com/iofabela/technical-challenge-meli/cmd/api/utils/web"
)

func (r *Repository) ProcessFile(ctx *gin.Context, file *bufio.Scanner) error {
	switch r.FileConfig.Format {
	case "csv":
		return r.CSVReader(ctx, file)
	case "jsonlines":
		return r.JSONLinesReader(ctx, file)
	case "txt":
		return r.TXTReader(ctx, file)
	default:
		return fmt.Errorf("NewFileReader - format not supported: %s", r.FileConfig.Format)
	}
}

// Process the file content as CSV
func (r *Repository) CSVReader(ctx *gin.Context, file *bufio.Scanner) error {
	for file.Scan() {
		line := file.Text()
		if err := r.ProcessLine(line); err != nil && err != io.EOF {
			web.Error(ctx, http.StatusInternalServerError, "Error to read process line")
			return fmt.Errorf("Error to read process line: %w", err)
		}
	}
	return nil
}

func (r *Repository) TXTReader(ctx *gin.Context, file *bufio.Scanner) error {

	for file.Scan() {
		line := file.Text()
		if err := r.ProcessLine(line); err != nil && err != io.EOF {
			web.Error(ctx, http.StatusInternalServerError, "Error to read process line")
			return fmt.Errorf("Error to read process line: %w", err)
		}
	}
	return nil
}
func (r *Repository) JSONLinesReader(ctx *gin.Context, file *bufio.Scanner) error {

	for file.Scan() {
		line := file.Text()
		if err := r.ProcessJson(line); err != nil && err != io.EOF {
			//TODO save the line with error in a file
			web.Error(ctx, http.StatusInternalServerError, "Error to read process line ")
			return fmt.Errorf("Error to read process line: %w", err)
		}
	}

	return nil
}

func (r *Repository) ProcessLine(line string) error {
	fmt.Println("Line:", strings.ReplaceAll(line, string(r.FileConfig.Separator), ""))
	dataLine := strings.Split(line, string(r.FileConfig.Separator))
	item, err := rest.RestMeli_Items(dataLine[0], dataLine[1], r.Client) // TODO get fields of site and id
	if err != nil {
		return err
	}
	fmt.Println("Item: ", item)
	// err = r.SqlService.SaveItem(item)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (r *Repository) ProcessJson(line string) error {
	var obj items.DataLine
	if err := json.Unmarshal([]byte(line), &obj); err != nil { // Process the JSON line
		return fmt.Errorf("Error to unmarshal JSON: %s", err.Error())
	}
	fmt.Println("JSON Line Object: ", obj)
	item, err := rest.RestMeli_Items(obj.Site, strconv.Itoa(obj.ID), r.Client)
	if err != nil {
		return err
	}

	fmt.Println("Item: ", obj.Site, strconv.Itoa(obj.ID), item)
	// err = r.SqlService.SaveItem(item)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func Reprocess(id string) error { // Reprocess the item if fail
	// Save in a file
	return nil
}
