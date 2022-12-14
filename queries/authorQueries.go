package queries

import (
	"library-go/database"
	"library-go/models"
)

func GetAuthors() ([]models.Author, error) {
	db := database.Connect()
	a := []models.Author{}
	err := db.Select(&a, "SELECT * FROM authors")
	if err != nil {
		return nil, err
	}
	return a, nil
}

func GetAuthor(id int) (models.Author, error) {
	db := database.Connect()
	a := models.Author{}

	err := db.Get(&a, "SELECT * FROM authors WHERE ID = $1", id)

	if err != nil {
		return a, err
	}

	return a, nil
}

func AddAuthor(a *models.Author) error {
	db := database.Connect()

	_, err := db.Exec("INSERT INTO authors (name, nickname, speciality) VALUES ($1, $2, $3)", a.Name, a.Nickname, a.Speciality)

	if err != nil {
		return err
	}

	return nil
}

func UpdateAuthor(id int, a *models.Author) error {
	db := database.Connect()

	if a.Name != "" {
		query := "UPDATE authors SET Name = $2 WHERE id = $1"
		_, err := db.Exec(query, id, a.Name)

		if err != nil {
			return err
		}
	}
	if a.Nickname != "" {
		query := "UPDATE authors SET Nickname = $2 WHERE id = $1"
		_, err := db.Exec(query, id, a.Nickname)

		if err != nil {
			return err
		}
	}
	if a.Speciality != "" {
		query := "UPDATE authors SET Speciality = $2 WHERE id = $1"
		_, err := db.Exec(query, id, a.Speciality)

		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteAuthor(id int) error {
	db := database.Connect()

	query := `DELETE FROM authors WHERE ID = $1`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func GetAuthorsBooks(id int) ([]string, error) {
	db := database.Connect()
	books := []string{}
	query := `SELECT Name FROM books WHERE AuthorID = $1`

	err := db.Select(&books, query, id)
	if err != nil {
		return nil, err
	}

	return books, nil
}
