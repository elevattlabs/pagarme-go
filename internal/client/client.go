package client

import (
	"github.com/elevattlabs/pagarme-go/internal/customer"
	"github.com/elevattlabs/pagarme-go/internal/order"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Client struct {
	secret   string
	Customer *customer.Service
	Order    *order.Service
}

func New(secret string) *Client {
	if secret == "" {
		panic("secret is required")
	}

	client := &fiber.Client{
		NoDefaultUserAgentHeader: false,
		JSONEncoder:              json.Marshal,
		JSONDecoder:              json.Unmarshal,
	}

	customersService := customer.NewCustomerService(client, secret)
	ordersService := order.NewOrdersService(client, secret)

	return &Client{
		secret:   secret,
		Customer: customersService,
		Order:    ordersService,
	}
}
