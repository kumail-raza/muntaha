package models

// Dua Dua
type Dua struct {
	Arabic      string `json:"arabic,omitempty" neo:"name"`
	Translation string `json:"translation,omitempty" neo:"translation"`
	Title       string `json:"title,omitempty" neo:"title"`
}
