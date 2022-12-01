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
	if err != nil {
		return err
	}
	_, err = db.Exec(schema)

	if err != nil {
		return err
	}

	schema = `CREATE TABLE authors (
	ID integer CONSTRAINT authors_pk PRIMARY KEY,
	Name text,
	Nickname text,
	Speciality text);`

	_, err = db.Exec("DROP TABLE IF EXISTS authors")
	if err != nil {
		return err
	}
	_, err = db.Exec(schema)
	if err != nil {
		return err
	}

	schema = `CREATE TABLE members (
	ID integer CONSTRAINT members_pk PRIMARY KEY,
	Name text);`

	_, err = db.Exec("DROP TABLE IF EXISTS members")
	if err != nil {
		return err
	}
	_, err = db.Exec(schema)
	if err != nil {
		return err
	}

	schema = `CREATE TABLE members_books (
	MemberID integer,
	BookID integer,
	CONSTRAINT FK_author FOREIGN KEY(MemberID)
        REFERENCES members(ID),
    CONSTRAINT FK_author FOREIGN KEY(BookID)
        REFERENCES books(ID));`

	_, err = db.Exec("DROP TABLE IF EXISTS members_books")
	if err != nil {
		return err
	}
	_, err = db.Exec(schema)
	if err != nil {
		return err
	}

	return err
}
