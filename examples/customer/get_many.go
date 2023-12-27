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

	res, err := client.Customer.GetMany(&models.ListCustomersParams{
		Page: 3,
		Size: 5,
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
