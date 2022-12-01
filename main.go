package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
	"library-go/data"
	"library-go/database"
	"log"
)

func setupRoutes(app *fiber.App) {
	app.Get("/books", data.GetBooks)
	app.Get("/book/:id", data.GetBook)
	app.Post("/book", data.AddBook)
	app.Patch("/book/:id", data.UpdateBook)
	app.Delete("/book/:id", data.DeleteBook)
	app.Get("")

	app.Get("/authors", data.GetAuthors)
	app.Get("/author/:id", data.GetAuthor)
	app.Post("/author", data.AddAuthor)
	app.Patch("/author/:id", data.UpdateAuthor)
	app.Delete("/author/:id", data.DeleteAuthor)
	app.Get("/author")
}

func main() {
	database.CreateTables()

	app := fiber.New()

	app.Use(
		logger.New(logger.Config{
			Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		}), // add Logger middleware
	)

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
