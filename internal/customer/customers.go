package customer

import (
	"github.com/gofiber/fiber/v2"
)

type Service struct {
	client *fiber.Client
	secret string
}

func NewCustomerService(client *fiber.Client, secret string) *Service {
	return &Service{client: client, secret: secret}
}
