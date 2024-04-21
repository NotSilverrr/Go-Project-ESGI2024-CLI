package salle

import (
	time "Go-Project-ESGI2024-CLI/time"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Salle struct {
	capacity int
	name     string
	booked   bool
}

func ShowAvailableRooms(db *sql.DB) {
	//user choose start date and hour for is reservation
	startDay, startMonth, startYear,
		startHour, startMinut,
		endDay, endMonth, endYear,
		endHour, endMinut := time.GetBook()

	fmt.Printf("Vous avez choisi une réservation commençant le %02d/%02d/%02d à %02d:%02d et se terminant le %02d/%02d/%02d à %02d:%02d.\nLes salles disponibles sont les suivantes : \n", startDay, startMonth, startYear, startHour, startMinut, endDay, endMonth, endYear, endHour, endMinut)

	rows, err := db.Query("Select name from room") //booked condition

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

func DisplayRooms(db *sql.DB) {
	rows, err := db.Query("Select id,name from room")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var id string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s. %s\n", id, name)
	}
}
