package models

type DocumentTypeInput string

const (
	DocumentTypeCPF       DocumentTypeInput = "CPF"
	DocumentTypeCNPJ      DocumentTypeInput = "CNPJ"
	DocumentTypePASSAPORT DocumentTypeInput = "PASSAPORT"
)
