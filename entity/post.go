package entity

// Post data model
type Post struct {
	ID    int64  `json:"ID"`
	Title string `json:"Title"`
	Text  string `json:"Text"`
}
