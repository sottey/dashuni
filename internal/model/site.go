package model

// Root is the top-level container for your JSON file.
type Root struct {
	Site Site `json:"site"`
}

// Site represents the top-level site object.
type Site struct {
	Metadata    Metadata `json:"metadata"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Favicon     string   `json:"favicon"`
	Theme       string   `json:"theme"`
	BaseURL     string   `json:"baseURL"`
	Version     string   `json:"version"`
	Pages       []Page   `json:"pages"`
}

// Metadata provides export information for auditing.
type Metadata struct {
	ExportedBy string `json:"exportedBy"`
	ExportedAt string `json:"exportedAt"`
	Owner      string `json:"owner"`
}

// Page represents a single dashboard page.
type Page struct {
	Order       int       `json:"order"`
	Title       string    `json:"title"`
	Logo        string    `json:"logo"`
	Description string    `json:"description"`
	Sections    []Section `json:"sections"`
}

// Section groups related Items on a page.
type Section struct {
	Order       int    `json:"order"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Items       []Item `json:"items"`
}

// Item is a single service/bookmark/tile.
type Item struct {
	Order       int      `json:"order"`
	Title       string   `json:"title"`
	ShortTitle  string   `json:"shortTitle"`
	URL         string   `json:"url"`
	PingURL     string   `json:"pingURL"`
	Icon        string   `json:"icon"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Target      string   `json:"target"`
}
