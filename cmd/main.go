package main

import (
	"log"

	"github.com/Mapmosphere/mphere_server_test/pkg/books"
	"github.com/Mapmosphere/mphere_server_test/pkg/common/config"
	"github.com/Mapmosphere/mphere_server_test/pkg/common/db"
	"github.com/Mapmosphere/mphere_server_test/pkg/tickets"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	// Default config
	app.Use(cors.New())

	db := db.Init(c.DBUrl)

	books.RegisterRoutes(app, db)
	tickets.RegisterRoutes(app, db)

	app.Listen(c.Port)
}
