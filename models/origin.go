package models

// Origin ...
type Origin struct {
	Type       string      `json:"type,omitempty"`
	References []Reference `json:"references,omitempty"`
}
