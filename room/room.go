package salle

import (
	"log"
	"strconv"
)

  type Salle struct {
  	capacity int
  	name     string
  	booked   bool
  }

  func ShowAvailableRooms(day string, month string, year string) {
			room_day, err := strconv.Atoi(day)

				if err != nil {
 				log.Fatal(err)
 			}

			room_year, err := strconv.Atoi(month) 
			
			if err != nil {
				log.Fatal(err)
			}
			
			//fflush
			println(room_day, month, room_year)
			// ---- //
			// VERIF DATES CORRECTES //
			// ---- //

  	 	rows, err := db.Query("Select name FROM room WHERE booked = 0")

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
