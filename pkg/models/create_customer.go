package models

type CreateCustomerInput struct {
	Name         string            `json:"name"`
	Email        string            `json:"email"`
	Code         string            `json:"code"`
	Document     string            `json:"document"`
	DocumentType DocumentTypeInput `json:"document_type"`
	Type         CustomerTypeInput `json:"type"`
	Gender       GendersInput      `json:"gender"`
	Address      AddressInput      `json:"address"`
	Phones       PhonesInput       `json:"phones"`
	BirthDate    string            `json:"birthdate"`
}
