package books

import (
	"github.com/Mapmosphere/mphere_server_test/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetBooks(c *fiber.Ctx) error {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&books)
}

func (h handler) GetBooksGeo(c *fiber.Ctx) error {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	// var geoData map[string]interface{}

	// return c.Status(fiber.StatusOK).JSON(&books)
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"type": "geojson",
		"data": books,
	})
}
