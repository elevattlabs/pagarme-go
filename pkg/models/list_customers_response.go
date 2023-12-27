package models

type CustomerInList struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Gender     GendersInput `json:"gender"`
	Email      string       `json:"email"`
	Delinquent bool         `json:"delinquent"`
	CreatedAt  string       `json:"created_at"`
	UpdatedAt  string       `json:"updated_at"`
}
type ListCustomersResponse struct {
	Data   []CustomerInList `json:"data"`
	Paging Paging           `json:"paging"`
}
