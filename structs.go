package serp

// Results ...
type Results struct {
	OrganicResults    []OrganicResult  `json:"organic_results"`
	SearchInformation Info             `json:"search_information"`
	SearchMetadata    Meta             `json:"search_metadata"`
	SearchParameters  Params           `json:"search_parameters"`
	SerpapiPagination interface{}      `json:"serpapi_pagination"`
	ShoppingResults   []ShoppingResult `json:"shopping_results"`
	ErrorMessage      string           `json:"error"`
}

// Params are search parameters
type Params struct {
	Device            string `json:"device"`
	Engine            string `json:"engine"`
	Filter            string `json:"filter"`
	Gl                string `json:"gl"`
	GoogleDomain      string `json:"google_domain"`
	Hl                string `json:"hl"`
	LocationRequested string `json:"location_requested"`
	LocationUsed      string `json:"location_used"`
	Q                 string `json:"q"`
	Tbm               string `json:"tbm"`
}

// Info is search information
type Info struct {
	QueryDisplayed string `json:"query_displayed"`
}

// Meta ...
type Meta struct {
	CreatedAt      string  `json:"created_at"`
	GoogleURL      string  `json:"google_url"`
	ID             string  `json:"id"`
	JSONEndpoint   string  `json:"json_endpoint"`
	ProcessedAt    string  `json:"processed_at"`
	RawHTMLFile    string  `json:"raw_html_file"`
	Status         string  `json:"status"`
	TotalTimeTaken float64 `json:"total_time_taken"`
}

// ShoppingResult ...
type ShoppingResult struct {
	Extensions        []string `json:"extensions"`
	Link              string   `json:"link"`
	Position          int      `json:"position"`
	ProductID         string   `json:"product_id"`
	Rating            int      `json:"rating"`
	Reviews           int      `json:"reviews"`
	SerpapiProductAPI string   `json:"serpapi_product_api"`
	Source            string   `json:"source"`
	Snippet           string   `json:"snippet"`
	Thumbnail         string   `json:"thumbnail"`
	Title             string   `json:"title"`
	Price             string   `json:"price"`
	ExtractedPrice    float64  `json:"extracted_price"`
}

// OrganicResult ...
type OrganicResult struct {
	Title          string  `json:"title,omitempty"`
	CachedPageLink string  `json:"cached_page_link,omitempty"`
	DisplayedLink  string  `json:"displayed_link,omitempty"`
	Link           string  `json:"link,omitempty"`
	Source         string  `json:"source"`
	Snippet        string  `json:"snippet,omitempty"`
	Price          string  `json:"price,omitempty"`
	ExtractedPrice float64 `json:"extracted_price"`
}
