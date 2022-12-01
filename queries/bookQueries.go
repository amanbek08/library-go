package queries

import (
	"library-go/database"
	"library-go/models"
)

func GetBooks() ([]models.Book, error) {
	db := database.Connect()
	b := []models.Book{}
	err := db.Select(&b, "SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GetBook(id int) (models.Book, error) {
	db := database.Connect()
	b := models.Book{}

	err := db.Get(&b, "SELECT * FROM books WHERE ID = $1", id)

	if err != nil {
		return b, err
	}

	return b, nil
}

func AddBook(b *models.Book) error {
	db := database.Connect()

	_, err := db.Exec("INSERT INTO books VALUES ($1, $2, $3, $4, $5)", b.ID, b.Name, b.Genre, b.ISBN, b.AuthorID)

	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO authors_books VALUES ($1, $2)", b.AuthorID, b.ID)

	if err != nil {
		return err
	}

	return nil
}

func UpdateBook(id int, b *models.Book) error {
	db := database.Connect()

	if b.Name != "" {
		query := "UPDATE books SET Name = $2 WHERE id = $1"
		_, err := db.Exec(query, id, b.Name)

		if err != nil {
			return err
		}
	}
	if b.Genre != "" {
		query := "UPDATE books SET Genre = $2 WHERE id = $1"
		_, err := db.Exec(query, id, b.Genre)

		if err != nil {
			return err
		}
	}
	if b.ISBN != "" {
		query := "UPDATE books SET ISBN = $2 WHERE id = $1"
		_, err := db.Exec(query, id, b.ISBN)

		if err != nil {
			return err
		}
	}
	if b.AuthorID != 0 {
		query := "UPDATE books SET AuthorID = $2 WHERE id = $1"
		_, err := db.Exec(query, id, b.AuthorID)

		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteBook(id int) error {
	db := database.Connect()

	query := `DELETE FROM books WHERE ID = $1`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
