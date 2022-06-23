package books

import (
	"github.com/Mapmosphere/mphere_server_test/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type AddBookRequestBody struct {
	Feature     string `json:"feature"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Remarks     string `json:"remarks"`
	// Lat         string `json:"lat"`
	// Long        string `json:"long"`
	Tags  string `json:"tags"`
	Tags2 string `json:"tags1"`
	Tags1 string `json:"tags2"`
	// Lattitude   float64 `json:"lat"`
	// Longitude   float64 `json:"long"`
	// Tags        map[string]interface{} `json:"tags" sql:"type:jsonb"`
}

func (h handler) AddBook(c *fiber.Ctx) error {
	body := AddBookRequestBody{}

	// parse body, attach to AddBookRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book models.Book

	book.Feature = body.Feature
	book.Name = body.Name
	book.Description = body.Description
	book.Remarks = body.Remarks
	// book.Lat = body.Lat
	// book.Long = body.Long
	// book.Tags = body.Tags
	// book.Tags = postgres.Jsonb{body.Tags}
	// book.Tags = postgres.Hstore{"properties": &body.Tags}
	book.Tags = body.Tags
	book.Tags1 = body.Tags1
	book.Tags2 = body.Tags2

	// json.Unmarshal([]byte, )

	// insert new db entry
	if result := h.DB.Create(&book); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&book)
}
