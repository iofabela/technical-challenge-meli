package load_file

import (
	"fmt"

	"github.com/iofabela/technical-challenge-meli/cmd/api/models/items"
)

var ToReprocess []*items.FailedItem

func Reprocess(failedItem *items.FailedItem) {
	ToReprocess = append(ToReprocess, failedItem)
	if failedItem.Response != nil {
		fmt.Printf("FailedItem - %s%s - %d\n", failedItem.Site, failedItem.ID, failedItem.Response.StatusCode)

	} else {
		fmt.Printf("FailedItem - %s%s - %s\n", failedItem.Site, failedItem.ID, failedItem.Error.Error())
	}
}
