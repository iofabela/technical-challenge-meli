package items

import "fmt"

type DataLine struct {
	Site string `json:"site" binding:"required"`
	ID   int    `json:"id" binding:"required"`
}

type Item struct {
	Price      float64 `json:"price" binding:"required"`
	StartTime  string  `json:"date_created" binding:"required"`
	CategoryID string  `json:"category_id" binding:"required"`
	CurrencyID string  `json:"currency_id" binding:"required"`
	SellerID   int     `json:"seller_id" binding:"required"`
}

type Categories struct {
	Name string `json:"name"`
}

type Currencies struct {
	Description string `json:"description"`
}

type Sellers struct {
	Nickname string `json:"nickname"`
}

type SaveItem struct {
	SiteID      string  `json:"site_id"`
	ID          string  `json:"id"`
	StartTime   string  `json:"start_time"`
	Price       float64 `json:"price"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Nickname    string  `json:"nickname"`
}

func (i Item) Validate() []error {

	errors := []error{}

	if i.Price <= 0 {
		errors = append(errors, fmt.Errorf("\n*price* must be greater than 0"))
	}
	if i.StartTime == "" {
		errors = append(errors, fmt.Errorf("\n*start_time* must be not empty"))
	}
	if i.CategoryID == "" {
		errors = append(errors, fmt.Errorf("\n*category_id* must be not empty"))
	}
	if i.CurrencyID == "" {
		errors = append(errors, fmt.Errorf("\n*currency_id* must be not empty"))
	}
	if i.SellerID == 0 {
		errors = append(errors, fmt.Errorf("\n*seller_id* must be not empty"))
	}
	return errors
}
