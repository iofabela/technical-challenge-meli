package items

type Data struct {
	Items []Item `json:"items"`
}

type Item struct {
	ID          int    `json:"id"`
	SiteID      int    `json:"site_id"`
	Price       int    `json:"price"`
	StartTime   string `json:"start_time"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Nickname    string `json:"nickname"`
}
