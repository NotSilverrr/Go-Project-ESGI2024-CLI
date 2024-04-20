package booking

import (
	verif "Go-Project-ESGI2024-CLI/verif"
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

func CreateReservation(idSalle int, dstart int, dend int, hstart int, hend int, db *sql.DB) {
	verif := verif.VerifResa(idSalle, dstart, dend, hstart, hend, db)
	if verif != "ok" {
		println(verif)
	}

	fmt.Println("Quelle salle souhaitez vous réserver ?")

	res, err := db.Exec("INSERT INTO reservation (salle_id,date_start,date_end,time_start,time_end) VALUES (?,?,?,?,?)", idSalle, dstart, dend, hstart, hend)

	if err != nil {
		println(res)
		log.Fatal(err)
	}
	println("Votre réservation a bien été validé")
}

func CancelReservation(id int, db *sql.DB) {
	res, err := db.Exec("DELETE FROM reservation WHERE id_salle=?", id)

	if err != nil {
		println(res)
		log.Fatal(err)
	}
}

func DisplayReservation(roomID int, db *sql.DB) {
	var name string
	var id string
	var datestart string
	var dateend string
	var timestart string
	var timeend string
	rows, err := db.Query("Select id,date_start,date_end,time_start,time_end FROM reservation WHERE id_salle=?", roomID)
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
		err = rowsRoom.Scan(&id, &datestart, &dateend, &timestart, &timeend)

		if err != nil {
			log.Fatal(err)
		}

		println(id, ".", datestart, timestart, " - ", dateend, timeend)
	}

}
