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

	app.Get("/authors", data.GetAuthors)
	app.Get("/author/:id", data.GetAuthor)
	app.Post("/author", data.AddAuthor)
	app.Patch("/author/:id", data.UpdateAuthor)
	app.Delete("/author/:id", data.DeleteAuthor)
	app.Get("/author/:id/books", data.GetAuthorsBooks)

	app.Get("/members", data.GetMembers)
	app.Get("/member/:id", data.GetMember)
	app.Post("/member", data.AddMember)
	app.Patch("/member/:id", data.UpdateMember)
	app.Delete("/member/:id", data.DeleteMember)
	app.Post("/members/:id/books/:id2", data.AddMemberBook)
	app.Get("/members/:id/books", data.GetMemberBooks)
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
