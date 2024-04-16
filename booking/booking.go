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

func CreateReservation(idSalle int, d_start string, d_end string, h_start string, h_end string, db *sql.DB) {
	VerifResa(idSalle, d_start, d_end, h_start, h_end, db)

	fmt.Println("Quelle salle souhaitez vous réserver ?")

	res, err := db.Exec("INSERT INTO reservation (salle_id,date_start,date_end,time_start,time_end) VALUES (?,?,?,?,?)", idSalle, d_start, d_end, h_start, h_end)

	if err != nil {
		log.Fatal(err)
	}
	println("Votre réservation a bien été validé")
}

func VerifResa(roomID int, d_start string, d_end string, h_start string, h_end string, db *sql.DB) string {
	var date_start string
	var date_end string
	var time_start string
	var time_end string

	rows, err := db.Query("Select id,date_start,date_end,time_start,time_end FROM reservation WHERE id_salle=?", roomID)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&date_start, &date_end, &time_start, &time_end)

		if err != nil {
			log.Fatal(err)
		}

		if d_start >= date_start || d_start <= date_end && d_start >= date_start {
			return "Cette salle n'est pas disponible le" + d_start
		}
		if d_start == date_start {
			if d_start != date_end {
				if h_start > time_start {
					return "Cette salle n'est pas disponible a" + h_start
				}
			} else {
				if h_start > time_start && h_start < time_end {
					return "Cette salle n'est pas disponible a" + h_start
				}
			}
		}

		if d_end >= date_start || d_end <= date_end && d_end > date_start {
			return "Cette salle n'est pas disponible le" + d_end
		}
		if d_end == date_start {
			if d_end != date_end {
				if h_end > time_start {
					return "Cette salle n'est pas disponible a" + h_end
				}
			} else {
				if h_end > time_start && h_end < time_end {
					return "Cette salle n'est pas disponible a" + h_end
				}
			}
		}

	}
}

func CancelReservation(id int, db *sql.DB) {
	res, err := db.Exec("DELETE FROM reservation WHERE id_salle=?", id)

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
		err = rowsRoom.Scan(&id, &date_start, &date_end, &time_start, &time_end)

		if err != nil {
			log.Fatal(err)
		}

		println(id, ".", date_start, time_start, " - ", date_end, time_end)
	}

}
