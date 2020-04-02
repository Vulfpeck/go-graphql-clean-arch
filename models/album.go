package models

type Album struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title"`
	Year  string `json:"year"`
	Genre string `json:"genre"`
}
