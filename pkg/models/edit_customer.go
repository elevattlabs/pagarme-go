package models

type EditCustomerParamsID string
type EditCustomerInput struct {
	CustomerID string `json:"customer_id,omitempty"`
	Body       EditCustomerBodyInput
}

type EditCustomerBodyInput struct {
	Name         string            `json:"name,omitempty"`
	Email        string            `json:"email,omitempty"`
	Code         string            `json:"code,omitempty"`
	Document     string            `json:"document,omitempty"`
	DocumentType DocumentTypeInput `json:"document_type,omitempty"`
	Type         CustomerTypeInput `json:"type,omitempty"`
	Gender       GendersInput      `json:"gender,omitempty"`
	BirthDate    string            `json:"birthdate,omitempty"`
}
