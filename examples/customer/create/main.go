package main

import (
	"fmt"
	"github.com/elevattlabs/pagarme-go"
	"github.com/elevattlabs/pagarme-go/pkg/models"
	"github.com/goccy/go-json"
	"os"
)

func main() {
	client := pagarme.NewClient(os.Getenv("PAGARME_API_KEY"))

	res, err := client.Customer.Create(models.CreateCustomerInput{
		Name:         "John Doe",
		Email:        "johndoe@example.com",
		Code:         "a419902e-996d-409d-ba7e-983d21e8028f",
		Document:     "12345678900",
		DocumentType: models.DocumentTypeCPF,
		Type:         models.CustomerTypeIndividual,
		Gender:       models.GenderMale,
		Address: models.AddressInput{
			Country: "BR",
			State:   "RJ",
			City:    "Niterói",
			ZipCode: "00000000",
			Line1:   "Rua XXX",
		},
		Phones: models.PhonesInput{
			Home: models.GenericPhoneInput{
				CountryCode: "55",
				AreaCode:    "21",
				Number:      "9xxxxxxxx",
			},
			Mobile: models.GenericPhoneInput{
				CountryCode: "55",
				AreaCode:    "21",
				Number:      "9xxxxxxxx",
			},
		},
		BirthDate: "08/03/2001",
	})
	if err != nil {
		fmt.Println(err)
	}
	ident, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(ident))
}
