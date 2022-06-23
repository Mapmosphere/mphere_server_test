package tickets

import (
	"github.com/Mapmosphere/mphere_server_test/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type AddTicketRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) AddTicket(c *fiber.Ctx) error {
	body := AddTicketRequestBody{}

	// parse body, attach to AddBookRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var ticket models.Ticket

	ticket.Title = body.Title
	ticket.Author = body.Author
	// ticket.Title = ticket.Title

	return nil
}
