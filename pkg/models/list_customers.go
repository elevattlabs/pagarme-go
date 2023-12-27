package models

type ListCustomersParams struct {
	Name     string       `json:"name,omitempty"`
	Document string       `json:"document,omitempty"`
	Email    string       `json:"email,omitempty"`
	Gender   GendersInput `json:"gender,omitempty"`
	Page     int32        `json:"page,omitempty"`
	Size     int32        `json:"size,omitempty"`
	Code     string       `json:"code,omitempty"`
}
