package models

type Artist struct {
	ID     string  `json:"id,omitempty"`
	Name   string  `json:"name"`
	Albums []Album `json:"albums"`
}
