package booking

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type booking struct {
	reference int
	name      string
	bookedOn  string
	duration  int
}

func CreateReservation(db *sql.DB) {
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

func CancelReservation(id int, db *sql.DB) {
	res, err := db.Exec("DELETE FROM réservation WHERE id_salle=?", id)
	println(res)
	if err != nil {
		log.Fatal(err)
	}
}

func DisplayReservation(roomID int, db *sql.DB) {
	var name string
	var id string
	var date_start string
	var date_end string
	var time_start string
	var time_end string
	rows, err := db.Query("Select id,date_start,date_end,time_start,time_end FROM réservation WHERE id_salle=?", roomID)
	rowsRoom, err := db.Query("Select name FROM room WHERE id=?", roomID)

	if err != nil {
		log.Fatal(err)
	}

	err = rows.Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	println("Réservations pour ", name)
	for rows.Next() {
		err = rowsRoom.Scan(&id, &date_start, &date_end, &time_start, &time_end)

		if err != nil {
			log.Fatal(err)
		}

		println(id, ".", date_start, time_start, " - ", date_end, time_end)
	}

}
