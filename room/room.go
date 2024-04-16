package salle

import (
	//"Go-Project-ESGI2024-CLI/main"
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

	startContext := "de début"
	endContext := "de fin"

	//user choose start date and hour for is reservation
	startDay, startMonth, startYear := time.GetDate(startContext) 
	startHour, startMinut := time.GetTime(startContext)
	
	//user choose end date and hour for is reservation
	endDay, endMonth, endYear := time.GetDate(endContext)
	endHour, endMinut := time.GetTime(endContext)

	fmt.Printf("Vous avez choisi une réservation commençant le %02d/%02d/%02d à %02d:%02d et se terminant le %02d/%02d/%02d à %02d:%02d.\nLes salles disponibles sont les suivantes : \n", startDay, startMonth, startYear, startHour, startMinut, endDay, endMonth, endYear, endHour, endMinut)

	rows, err := db.Query("Select * from room")

	if err != nil {
		log.Fatal(err)
	 }
	 
	 for rows.Next() {
	 	var name string
	 	if err := rows.Scan(&name); err != nil {
	 		log.Fatal(err)
	 	}
	 	fmt.Println(name)
	 }
}
