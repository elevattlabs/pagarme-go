package models

type EditCustomerResponse struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	Email      string            `json:"email"`
	Document   string            `json:"document"`
	Gender     GendersInput      `json:"gender"`
	Type       CustomerTypeInput `json:"type"`
	Delinquent bool              `json:"delinquent"`
	CreatedAt  string            `json:"created_at"`
	UpdatedAt  string            `json:"updated_at"`
	Phones     PhonesInput       `json:"phones"`
}
