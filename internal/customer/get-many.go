package customer

import (
	"errors"
	"github.com/elevattlabs/pagarme-go/configs"
	"github.com/elevattlabs/pagarme-go/pkg/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func (s *Service) GetMany(input *models.ListCustomersParams) (*models.ListCustomersResponse, error) {
	res := new(models.ListCustomersResponse)
	agent := s.client.Get(configs.UrlV5+"/customers?").BasicAuth(s.secret, "")
	name := "name" + "=" + input.Name + "&"
	email := "email" + "=" + input.Email + "&"
	document := "document" + "=" + input.Document + "&"
	page := "page" + "=" + strconv.Itoa(int(input.Page)) + "&"
	size := "size" + "=" + strconv.Itoa(int(input.Size)) + "&"
	cd := "code" + "=" + input.Code

	agent.QueryString(name + email + document + page + size + cd)
	//query["document"] = "document" + "=" + input.Document + "&"
	//query["email"] = "email" + "=" + input.Email + "&"
	//query["page"] = "page" + "=" + string(input.Page) + "&"
	//query["size"] = "size" + "=" + string(input.Size) + "&"
	//query["code"] = "code" + "=" + input.Code
	log.Println(string(agent.Request().URI().String()))
	log.Println(strconv.Itoa(int(input.Page)))
	code, body, _ := agent.Struct(&res)

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
