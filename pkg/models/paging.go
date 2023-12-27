package models

type Paging struct {
	Total int    `json:"total"`
	Next  string `json:"next"`
}
