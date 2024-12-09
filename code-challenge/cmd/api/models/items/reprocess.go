package items

import "net/http"

type Reprocess struct {
	Items []FailedItem `json:"items"`
}

type FailedItem struct {
	Site     string         `json:"site"`
	ID       string         `json:"id"`
	Response *http.Response `json:"response"`
	Error    error          `json:"error"`
}
