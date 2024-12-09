package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
		item     items.Item
		category items.Categories
		currency items.Currencies
		seller   items.Sellers
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
		return nil, &items.FailedItem{ID: id, Site: site, Response: respItem, Error: fmt.Errorf("RestMeli_Items - Item:Error parsing JSON: %v", errJsonItem.Error())}
	}

	// CATEGORY
	respCategory, errRespCategory := http.Get(client.BaseURL + client.Endpoints.Categories + item.CategoryID)
	if errRespCategory != nil {
		return nil, &items.FailedItem{ID: id, Site: site, Response: respCategory, Error: fmt.Errorf("RestMeli_Items - Category:Error - %s", errRespCategory.Error())}
	}
	defer respCategory.Body.Close()
	if respCategory.StatusCode != http.StatusOK {
		return nil, &items.FailedItem{ID: id, Site: site, Response: respCategory, Error: fmt.Errorf("RestMeli_Items - Category:Status - %d", respCategory.StatusCode)}
	}
	// Parsear el JSON
	if errJsonCategory := json.NewDecoder(respCategory.Body).Decode(&category); errJsonCategory != nil {
		return nil, &items.FailedItem{ID: id, Site: site, Response: respCategory, Error: fmt.Errorf("RestMeli_Items - Category:Error parsing JSON: %v", errJsonCategory.Error())}
	}

	// CURRENCY
	respCurrency, errRespCurrency := http.Get(client.BaseURL + client.Endpoints.Currencies + item.CurrencyID)
	if errRespCurrency != nil {
		return nil, &items.FailedItem{ID: id, Site: site, Response: respCurrency, Error: fmt.Errorf("RestMeli_Items - Currency:Error - %s", errRespCurrency.Error())}
	}
	defer respCurrency.Body.Close()
	if respCurrency.StatusCode != http.StatusOK {
		return nil, &items.FailedItem{ID: id, Site: site, Response: respCurrency, Error: fmt.Errorf("RestMeli_Items - Currency:Status - %d", respCurrency.StatusCode)}
	}
	// Parsear el JSON
	if errJsonCurrency := json.NewDecoder(respCurrency.Body).Decode(&currency); errJsonCurrency != nil {
		return nil, &items.FailedItem{ID: id, Site: site, Response: respCurrency, Error: fmt.Errorf("RestMeli_Items - Currency:Error parsing JSON: %v", errJsonCurrency.Error())}
	}

	// SELLER
	respSeller, errRespSeller := http.Get(client.BaseURL + client.Endpoints.Sellers + strconv.Itoa(item.SellerID))
	if errRespSeller != nil {
		return nil, &items.FailedItem{ID: id, Site: site, Response: respSeller, Error: fmt.Errorf("RestMeli_Items - Seller:Error - %s", errRespSeller.Error())}
	}
	defer respSeller.Body.Close()
	if respSeller.StatusCode != http.StatusOK {
		return nil, &items.FailedItem{ID: id, Site: site, Response: respSeller, Error: fmt.Errorf("RestMeli_Items - Seller:Status - %d", respSeller.StatusCode)}
	}
	// Parsear el JSON
	if errJsonSeller := json.NewDecoder(respSeller.Body).Decode(&seller); errJsonSeller != nil {
		return nil, &items.FailedItem{ID: id, Site: site, Response: respSeller, Error: fmt.Errorf("RestMeli_Items - Seller:Error parsing JSON: %v", errJsonSeller.Error())}
	}
	return &items.SaveItem{
		ID:          id,
		SiteID:      site,
		Price:       item.Price,
		StartTime:   item.StartTime,
		Name:        category.Name,
		Description: currency.Description,
		Nickname:    seller.Nickname,
	}, &items.FailedItem{}
}
