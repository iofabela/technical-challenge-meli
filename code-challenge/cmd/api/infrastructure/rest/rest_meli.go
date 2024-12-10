package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/iofabela/technical-challenge-meli/cmd/api/models/items"
)

type Endpoints struct {
	Items_price string
	Items_time  string
	Categories  string
	Currencies  string
	Sellers     string
}

// Client …
type Client struct {
	BaseURL      string
	PostPipeName string
	HTTP         *http.Client
	Endpoints    Endpoints
}

func NewClient(baseURL string, endpoints Endpoints) *Client {
	return &Client{
		BaseURL:      baseURL,
		PostPipeName: "post_pipe",
		HTTP:         http.DefaultClient,
		Endpoints:    endpoints,
	}
}

func RestMeli_Items(site string, id string, client *Client) (*items.SaveItem, *items.FailedItem) {

	var (
		item         items.Item
		category     items.Categories
		currency     items.Currencies
		seller       items.Sellers
		errorsValues []string
	)
	// ITEM
	respItem, errRespItem := http.Get(client.BaseURL + client.Endpoints.Items_price + site + id)
	if errRespItem != nil {
		return nil, &items.FailedItem{ID: id, Site: site, Error: fmt.Errorf("RestMeli_Items - Item:Error - %s", errRespItem.Error())}
	}
	defer respItem.Body.Close()

	if respItem.StatusCode != http.StatusOK {
		return nil, &items.FailedItem{ID: id, Site: site, Response: respItem, Error: fmt.Errorf("RestMeli_Items - Item:Status - %d", respItem.StatusCode)}
	}
	// Parsear el JSON
	if errJsonItem := json.NewDecoder(respItem.Body).Decode(&item); errJsonItem != nil {
		// return nil, &items.FailedItem{ID: id, Site: site, Response: respItem, Error: fmt.Errorf("RestMeli_Items - Item:Error parsing JSON: %v", errJsonItem.Error())}
		errorsValues = append(errorsValues, fmt.Sprintf("RestMeli_Items with %s%s - Item:Error parsing JSON", site, id))
	}

	// CATEGORY
	respCategory, errRespCategory := http.Get(client.BaseURL + client.Endpoints.Categories + item.CategoryID)
	if errRespCategory != nil {
		return nil, &items.FailedItem{ID: id, Site: site, Response: respCategory, Error: fmt.Errorf("RestMeli_Items - Category:Error - %s", errRespCategory.Error())}
	}
	defer respCategory.Body.Close()
	if respCategory.StatusCode != http.StatusOK {
		// return nil, &items.FailedItem{ID: id, Site: site, Response: respCategory, Error: fmt.Errorf("RestMeli_Items - Category:Status - %d", respCategory.StatusCode)}
		// fmt.Printf("RestMeli_Items - Category:Status - %d\n", respCategory.StatusCode)
		errorsValues = append(errorsValues, fmt.Sprintf("RestMeli_Items with ID %s%s - Category:Status - %d", site, id, respCategory.StatusCode))
	}
	// Parsear el JSON
	if errJsonCategory := json.NewDecoder(respCategory.Body).Decode(&category); errJsonCategory != nil {
		// return nil, &items.FailedItem{ID: id, Site: site, Response: respCategory, Error: fmt.Errorf("RestMeli_Items - Category:Error parsing JSON: %v", errJsonCategory.Error())}
		// fmt.Printf("RestMeli_Items - Category:Error parsing JSON: %s\n", errJsonCategory.Error())
		errorsValues = append(errorsValues, fmt.Sprintf("RestMeli_Items with ID %s%s - Category:Error parsing JSON", site, id))
	}

	// CURRENCY
	respCurrency, errRespCurrency := http.Get(client.BaseURL + client.Endpoints.Currencies + item.CurrencyID)
	if errRespCurrency != nil {
		return nil, &items.FailedItem{ID: id, Site: site, Response: respCurrency, Error: fmt.Errorf("RestMeli_Items - Currency:Error - %s", errRespCurrency.Error())}
	}
	defer respCurrency.Body.Close()
	if respCurrency.StatusCode != http.StatusOK {
		// return nil, &items.FailedItem{ID: id, Site: site, Response: respCurrency, Error: fmt.Errorf("RestMeli_Items - Currency:Status - %d", respCurrency.StatusCode)}
		// fmt.Printf("RestMeli_Items - Currency:Status - %d \n", respCurrency.StatusCode)
		errorsValues = append(errorsValues, fmt.Sprintf("RestMeli_Items with ID %s%s - Currency:Status - %d", site, id, respCurrency.StatusCode))
	}
	// Parsear el JSON
	if errJsonCurrency := json.NewDecoder(respCurrency.Body).Decode(&currency); errJsonCurrency != nil {
		// return nil, &items.FailedItem{ID: id, Site: site, Response: respCurrency, Error: fmt.Errorf("RestMeli_Items - Currency:Error parsing JSON: %v", errJsonCurrency.Error())}
		// fmt.Printf("RestMeli_Items - Currency:Error parsing JSON: %s\n", errJsonCurrency.Error())
		errorsValues = append(errorsValues, fmt.Sprintf("RestMeli_Items with ID %s%s - Currency:Error parsing JSON", site, id))
	}

	// SELLER
	respSeller, errRespSeller := http.Get(client.BaseURL + client.Endpoints.Sellers + strconv.Itoa(item.SellerID))
	if errRespSeller != nil {
		return nil, &items.FailedItem{ID: id, Site: site, Response: respSeller, Error: fmt.Errorf("RestMeli_Items - Seller:Error - %s", errRespSeller.Error())}
	}
	defer respSeller.Body.Close()
	if respSeller.StatusCode != http.StatusOK {
		// return nil, &items.FailedItem{ID: id, Site: site, Response: respSeller, Error: fmt.Errorf("RestMeli_Items - Seller:Status - %d", respSeller.StatusCode)}
		// fmt.Printf("RestMeli_Items - Seller:Status - %d\n", respSeller.StatusCode)
		errorsValues = append(errorsValues, fmt.Sprintf("RestMeli_Items with ID %s%s - Seller:Status - %d", site, id, respSeller.StatusCode))
	}
	// Parsear el JSON
	if errJsonSeller := json.NewDecoder(respSeller.Body).Decode(&seller); errJsonSeller != nil {
		// return nil, &items.FailedItem{ID: id, Site: site, Response: respSeller, Error: fmt.Errorf("RestMeli_Items - Seller:Error parsing JSON: %v", errJsonSeller.Error())}
		// fmt.Printf("RestMeli_Items - Seller:Error parsing JSON: %s\n", errJsonSeller.Error())
		errorsValues = append(errorsValues, fmt.Sprintf("RestMeli_Items with ID %s%s - Seller:Error parsing JSON", site, id))
	}

	if len(errorsValues) > 0 {
		combinedErrors := strings.Join(errorsValues, "; ")
		fmt.Println(combinedErrors)
	}

	saveItem := items.SaveItem{
		ID:          id,
		SiteID:      site,
		Price:       item.Price,
		StartTime:   item.StartTime,
		Name:        category.Name,
		Description: currency.Description,
		Nickname:    seller.Nickname,
	}
	saveItem = saveItem.Validate()

	return &saveItem, &items.FailedItem{}
}
