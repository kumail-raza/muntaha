package models

// Reference Reference
type Reference struct {
	ID        int64  `json:"id,omitempty" neo:"id"`
	Name      string `json:"name,omitempty" neo:"name,primary"`
	RefNumber string `json:"refNumber,omitempty" neo:"refNumber"`
}
