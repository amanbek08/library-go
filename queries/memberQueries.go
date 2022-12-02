package queries

import (
	"library-go/database"
	"library-go/models"
)

var lastmID = 0

func GetMembers() ([]models.Member, error) {
	db := database.Connect()
	m := []models.Member{}
	err := db.Select(&m, "SELECT * FROM members")
	if err != nil {
		return nil, err
	}
	return m, nil
}

func GetMember(id int) (models.Member, error) {
	db := database.Connect()
	m := models.Member{}

	err := db.Get(&m, "SELECT * FROM members WHERE ID = $1", id)

	if err != nil {
		return m, err
	}

	return m, nil
}

func getLastMemberID() (int, error) {
	db := database.Connect()
	var id int
	query := `SELECT ID
	FROM members 
	ORDER BY ID DESC 
	LIMIT 1`
	err := db.Get(&id, query)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func AddMember(m *models.Member) error {
	db := database.Connect()

	//id, err := getLastMemberID()
	//
	//if err != nil {
	//	return err
	//}

	id := lastmID

	_, err := db.Exec("INSERT INTO members VALUES ($1, $2)", id, m.Name)

	if err != nil {
		return err
	}

	lastmID = lastmID + 1

	return nil
}

func UpdateMember(id int, m *models.Member) error {
	db := database.Connect()

	if m.Name != "" {
		query := "UPDATE members SET Name = $2 WHERE id = $1"
		_, err := db.Exec(query, id, m.Name)

		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteMember(id int) error {
	db := database.Connect()

	query := `DELETE FROM members WHERE ID = $1`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func AddMemberBook(id int, bookid int) error {
	db := database.Connect()

	query := `INSERT INTO members_books VALUES ($1, $2)`

	_, err := db.Exec(query, id, bookid)
	if err != nil {
		return err
	}

	return nil
}

func GetMemberBooks(id int) ([]string, error) {
	db := database.Connect()
	books := []string{}

	query := `SELECT books.Name 
	FROM members_books
	JOIN books ON books.ID = members_books.BookID
	WHERE members_books.MemberID = $1`

	err := db.Select(&books, query, id)
	if err != nil {
		return nil, err
	}

	return books, nil
}
