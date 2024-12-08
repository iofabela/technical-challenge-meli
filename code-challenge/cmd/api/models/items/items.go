package items

type Data struct {
	Items []SaveItem `json:"save_item"`
}

type Item struct {
	Price      int    `json:"price" binding:"required"`
	StartTime  string `json:"date_created" binding:"required"`
	CategoryID string `json:"category_id" binding:"required"`
	CurrencyID string `json:"currency_id" binding:"required"`
	SellerID   int    `json:"seller_id" binding:"required"`
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
	SiteID      int    `json:"site_id"`
	ID          int    `json:"id"`
	StartTime   string `json:"start_time"`
	Price       int    `json:"price"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Nickname    string `json:"nickname"`
}
