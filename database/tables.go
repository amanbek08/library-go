package database

func CreateTables() error {

	db := Connect()

	schema := `CREATE TABLE books (
	ID integer CONSTRAINT books_pk PRIMARY KEY,
	Name text,
	Genre text,
	ISBN text,
	AuthorID integer);`

	_, err := db.Exec("DROP TABLE IF EXISTS books")
	_, err = db.Exec(schema)

	schema = `CREATE TABLE authors (
	ID integer CONSTRAINT authors_pk PRIMARY KEY,
	Name text,
	Nickname text,
	Speciality text);`

	_, err = db.Exec("DROP TABLE IF EXISTS authors")
	_, err = db.Exec(schema)

	return err
}
