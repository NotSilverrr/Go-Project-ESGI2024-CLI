package booking

import (
	"fmt"
	"log"
)

type booking struct {
	reference int
	name      string
	bookedOn  string
	duration  int
}

func CreateReservation() {
	fmt.Println("Quelle est la date du début de votre séjour ? (JJ-MM-AAAA)")
	date := ""
	fmt.Scan(&date)

	//implement date verification

	fmt.Println("Quelle sera la durée de votre séjour ?")
	duration := 0
	fmt.Scan(&duration)
	
	//verif duration of the reservation isn't going on another book

	fmt.Println("Quelle salle souhaitez vous réserver ?")

  rows, err := db.Query("Select name FROM room WHERE ?") //display condition with prepared request

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

	room := "" //number or room name?
	fmt.Scan(&room)
	
	res, err := db.Exec("INSERT INTO ? (?) VALUES ($?)", "?") //booking process

	if err != nil {
		log.Fatal(err)
	}
}

func CancelReservation() {}

func DisplayReservation() {}