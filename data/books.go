package data

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type Book struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	ISBN     string `json:"isbn"`
	AuthorID int    `json:"authorId"`
}

func GetBooks(c *fiber.Ctx) error {

	return c.JSON(bookList)

}

func findBook(id int) (*Book, error) {
	for _, book := range bookList {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fiber.ErrBadRequest
}

func GetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		panic(err)
	}
	book, err := findBook(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(book)
}

func getLastId() int {
	return bookList[len(bookList)-1].ID
}

func AddBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return err
	}

	book.ID = getLastId() + 1
	bookList = append(bookList, book)
	log.Output(1, "201")
	return c.JSON(bookList)
}

func UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Panic(err)
	}

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		log.Panic(err)
	}

	bookToChange, err := findBook(id)
	if err != nil {
		log.Panic(err)
	}

	if book.Name != "" {
		bookToChange.Name = book.Name
	}
	if book.Genre != "" {
		bookToChange.Genre = book.Genre
	}
	if book.ISBN != "" {
		bookToChange.ISBN = book.ISBN
	}
	if book.AuthorID != 0 {
		bookToChange.AuthorID = book.AuthorID
	}

	return c.JSON(bookList)

}

var bookList = []*Book{
	&Book{
		ID:       1,
		Name:     "Potter",
		Genre:    "Fantasy",
		ISBN:     "0-1234-1234-0",
		AuthorID: 3,
	},
	&Book{
		ID:       2,
		Name:     "Interstellar",
		Genre:    "Sci-Fi",
		ISBN:     "4-3244-1568-7",
		AuthorID: 7,
	},
}
