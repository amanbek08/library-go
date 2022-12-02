package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func Connect() *sqlx.DB {
	host := "db"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbname := "test_db"

	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user, password, host, port, dbname)
	db, err := sqlx.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	return db
}
