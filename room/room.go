package salle

import (
	//"Go-Project-ESGI2024-CLI/main"
	time "Go-Project-ESGI2024-CLI/time"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Salle struct {
	capacity int
	name     string
	booked   bool
}

func ShowAvailableRooms(day string, month string, year string, db *sql.DB) {

	startContext := "de début"
	endContext := "de fin"

	startDay, startMonth, startYear := time.GetDate(startContext)
	startHour, startMinut := time.GetTime(startContext)
	
	endDay, endMonth, endYear := time.GetDate(endContext)
	endHour, endMinut := time.GetTime(endContext)

	fmt.Printf("Votre réservation commencera le %02d/%02d/%02d pour %02d:%02d\n", startDay, startMonth, startYear, startHour, startMinut)

	fmt.Printf("Votre réservation se terminera le %02d/%02d/%02d pour %02d:%02d\n", endDay, endMonth, endYear, endHour, endMinut)

	// ---- //
	// VERIF CORRECT DATES //
	// ---- //

	// rows, err := db.Query("Select name FROM room WHERE ?") // //display condition with prepared request

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var name string
	// 	if err := rows.Scan(&name); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(name)
	// }
}
