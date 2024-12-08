package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/iofabela/technical-challenge-meli/cmd/api/models/items"
	"github.com/iofabela/technical-challenge-meli/cmd/api/utils/web"
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

func RestMeli_Items(site string, id string, client *Client) (*items.SaveItem, error) {

	var (
		item     items.Item
		category items.Categories
		currency items.Currencies
		seller   items.Sellers
	)
	// ITEM
	respItem, err := http.Get(client.BaseURL + client.Endpoints.Items_price + site + id)
	if err != nil {
		fmt.Printf("Error al consultar %s: %v\n", "API-MELI", err)
	}
	defer respItem.Body.Close()

	if respItem.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("RestMeli_Items - Item:Status - %d", respItem.StatusCode)
	}
	// Parsear el JSON
	if err = json.NewDecoder(respItem.Body).Decode(&item); err != nil {
		return nil, fmt.Errorf("RestMeli_Items - Item:Error parsing JSON: %v", err.Error())
	}

	// CATEGORY
	respCategory, err := http.Get(client.BaseURL + client.Endpoints.Categories + item.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("RestMeli_Items - Category:Error - %s", err.Error())
	}
	defer respCategory.Body.Close()

	if respCategory.StatusCode != http.StatusOK {
		web.Error(nil, respCategory.StatusCode, "Error to consult the API")
		return nil, fmt.Errorf("RestMeli_Items - Category:Status - %d", respCategory.StatusCode)
	}
	// Parsear el JSON
	if err = json.NewDecoder(respCategory.Body).Decode(&category); err != nil {
		return nil, fmt.Errorf("RestMeli_Items - Category:Error parsing JSON: %v", err.Error())
	}

	// CURRENCY
	respCurrency, err := http.Get(client.BaseURL + client.Endpoints.Currencies + item.CurrencyID)
	if err != nil {
		return nil, fmt.Errorf("RestMeli_Items - Currency:Error - %s", err.Error())
	}
	defer respCurrency.Body.Close()

	if respCurrency.StatusCode != http.StatusOK {
		web.Error(nil, respCurrency.StatusCode, "Error to consult the API")
		return nil, fmt.Errorf("RestMeli_Items - Currency:Status - %d", respCurrency.StatusCode)
	}
	// Parsear el JSON
	if err = json.NewDecoder(respCurrency.Body).Decode(&currency); err != nil {
		return nil, fmt.Errorf("RestMeli_Items - Currency:Error parsing JSON: %v", err.Error())
	}

	// SELLER
	respSeller, err := http.Get(client.BaseURL + client.Endpoints.Sellers + strconv.Itoa(item.SellerID))
	if err != nil {
		return nil, fmt.Errorf("RestMeli_Items - Seller:Error - %s", err.Error())
	}
	defer respSeller.Body.Close()

	if respSeller.StatusCode != http.StatusOK {
		web.Error(nil, respSeller.StatusCode, "Error to consult the API")
		return nil, fmt.Errorf("RestMeli_Items - Seller:Status - %d", respSeller.StatusCode)
	}
	// Parsear el JSON
	if err = json.NewDecoder(respSeller.Body).Decode(&seller); err != nil {
		return nil, fmt.Errorf("RestMeli_Items - Seller:Error parsing JSON: %v", err.Error())
	}
	return &items.SaveItem{
		Price:       item.Price,
		StartTime:   item.StartTime,
		Name:        category.Name,
		Description: currency.Description,
		Nickname:    seller.Nickname,
	}, nil
}
