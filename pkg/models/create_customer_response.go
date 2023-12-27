package models

type CreateCustomerResponse struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Email        string            `json:"email"`
	Code         string            `json:"code"`
	Document     string            `json:"document"`
	DocumentType DocumentTypeInput `json:"document_type"`
	Type         CustomerTypeInput `json:"type"`
	Delinquent   bool              `json:"delinquent"`
	Gender       GendersInput      `json:"gender"`
	Address      Address           `json:"address"`
	Phones       PhonesInput       `json:"phones"`
	BirthDate    string            `json:"birthdate"`
	CreatedAt    string            `json:"created_at"`
	UpdatedAt    string            `json:"updated_at"`
}
