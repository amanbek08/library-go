package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"library-go/data"
	"log"
)

func setupRoutes(app *fiber.App) {
	app.Get("/books", data.GetBooks)
	app.Get("/book/:id", data.GetBook)
	app.Post("/book", data.AddBook)
	app.Patch("/book/:id", data.UpdateBook)
}

func main() {

	app := fiber.New()

	app.Use(
		logger.New(logger.Config{
			Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		}), // add Logger middleware
	)

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
