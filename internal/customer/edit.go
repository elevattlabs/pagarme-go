package customer

import (
	"errors"
	"github.com/elevattlabs/pagarme-go/configs"
	"github.com/elevattlabs/pagarme-go/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func (s *Service) Edit(input models.EditCustomerInput) (*models.EditCustomerResponse, error) {
	res := new(models.EditCustomerResponse)
	code, body, _ := s.client.Put(configs.UrlV5+"/customers/"+input.CustomerID).BasicAuth(s.secret, "").JSON(input.Body).Struct(&res)

	switch code {
	case 200:
		return res, nil
	case 400:
		return nil, fiber.ErrBadRequest
	case 401:
		return nil, fiber.ErrUnauthorized
	case 404:
		return nil, fiber.ErrNotFound
	case 500:
		return nil, fiber.ErrInternalServerError
	case 422:
		return nil, errors.New(string(body))
	case 412:
		return nil, errors.New(string(body))
	case 429:
		return nil, errors.New(string(body))
	}

	return res, nil

}
