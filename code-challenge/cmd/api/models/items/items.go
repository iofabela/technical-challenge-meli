package items

type DataLine struct {
	Site string `json:"site" binding:"required"`
	ID   int    `json:"id" binding:"required"`
}

type Item struct {
	Price      float64 `json:"price"`
	StartTime  string  `json:"date_created"`
	CategoryID string  `json:"category_id"`
	CurrencyID string  `json:"currency_id"`
	SellerID   int     `json:"seller_id"`
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
	SiteID      string  `json:"site"`
	ID          string  `json:"id"`
	Price       float64 `json:"price"`
	StartTime   string  `json:"start_time"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Nickname    string  `json:"nickname"`
}

func (i SaveItem) Validate() SaveItem {

	if i.Price <= 0 {
		i.Price = 0
	}
	if i.StartTime == "" {
		i.StartTime = "2023-01-01T00:00:00Z"
	}
	if i.Description == "" {
		i.Description = "No description"
	}
	if i.Nickname == "" {
		i.Nickname = "No nickname"
	}
	if i.Name == "" {
		i.Name = "No name"
	}

	return i
}
