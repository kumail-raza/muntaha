package models

// Origin ...
type Origin struct {
	Type       string      `json:"type,omitempty" neo:"type"`
	References []Reference `json:"references,omitempty" neo:"references"`
}
