package models

type AddressInput struct {
	Country string  `json:"country"`
	State   string  `json:"state"`
	City    string  `json:"city"`
	ZipCode string  `json:"zip_code"`
	Line1   string  `json:"line_1"`
	Line2   *string `json:"line_2,omitempty"`
}

type Address struct {
	Country   string  `json:"country"`
	State     string  `json:"state"`
	City      string  `json:"city"`
	ZipCode   string  `json:"zip_code"`
	Line1     string  `json:"line_1"`
	Line2     *string `json:"line_2,omitempty"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
