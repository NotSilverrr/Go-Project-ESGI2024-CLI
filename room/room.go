package salle

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Salle struct {
	capacity int
	name     string
	booked   bool
}

func ShowAvailableRooms(day string, month string, year string, db *sql.DB) {
	room_day, err := strconv.Atoi(day)

	if err != nil {
		log.Fatal(err)
	}

	room_year, err := strconv.Atoi(month)

	if err != nil {
		log.Fatal(err)
	}

	// ---- //
	// VERIF CORRECT DATES //
	// ---- //

	rows, err := db.Query("Select name FROM room WHERE ?") // //display condition with prepared request

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
	}
}
