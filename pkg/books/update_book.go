package books

import (
	"github.com/Mapmosphere/mphere_server_test/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateBookRequestBody struct {
	Feature     string `json:"feature"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Remarks     string `json:"remarks"`
	Lat         string `json:"lat"`
	Long        string `json:"long"`
	// Lattitude   float64 `json:"lat"`
	// Longitude   float64 `json:"long"`
	Tags  string `json:"tags"`
	Tags2 string `json:"tags1"`
	Tags1 string `json:"tags2"`
	// Tags        map[string]interface{} `json:"tags" sql:"type:jsonb"`
}

func (h handler) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateBookRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	book.Feature = body.Feature
	book.Name = body.Name
	// book.Description = body.Description
	book.Description = body.Description
	book.Remarks = body.Remarks
	book.Lat = body.Lat
	book.Long = body.Long
	// book.Tags = body.Tags
	// book.Tags = postgres.Jsonb{body.Tags}
	book.Tags = body.Tags
	book.Tags1 = body.Tags1
	book.Tags2 = body.Tags2

	// save book
	h.DB.Save(&book)

	return c.Status(fiber.StatusOK).JSON(&book)
}
