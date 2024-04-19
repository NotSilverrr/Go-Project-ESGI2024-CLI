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

	//startContext := "de début"
	//endContext := "de fin"

	//user choose start date and hour for is reservation
	startDay, startMonth, startYear, endDay, endMonth, endYear:= time.GetDate() 
	//startHour, startMinut := time.GetTime(startContext)
	
	//user choose end date and hour for is reservation
	//endHour, endMinut := time.GetTime(endContext)

	fmt.Printf("Vous avez choisi une réservation commençant le %02d/%02d/%02d à [time] et se terminant le %02d/%02d/%02d à [time].\nLes salles disponibles sont les suivantes : \n", startDay, startMonth, startYear, endDay, endMonth, endYear)

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
