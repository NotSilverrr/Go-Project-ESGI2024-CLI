package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Db_opener() *sql.DB {
	db, err := sql.Open("mysql", "root:Respons11@tcp(localhost:3306)/booking")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Db_closer(db *sql.DB) {
	defer db.Close()
}
