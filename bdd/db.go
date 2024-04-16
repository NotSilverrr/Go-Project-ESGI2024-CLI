package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Db_opener() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydatabase")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db
}
