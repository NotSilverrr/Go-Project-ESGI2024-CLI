package booking

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type booking struct {
	reference int
	name      string
	bookedOn  string
	duration  int
}

func CreateReservation(idSalle int, dstart int, dend int, hstart int, hend int, db *sql.DB) {
	verif := VerifResa(idSalle, dstart, dend, hstart, hend, db)
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

func VerifResa(roomID int, dstart int, dend int, hstart int, hend int, db *sql.DB) string {
	var datestart string
	var dateend string
	var timestart string
	var timeend string

	datestartint, err := strconv.Atoi(datestart)
	dateendint, err := strconv.Atoi(dateend)
	timestartint, err := strconv.Atoi(timestart)
	timeendint, err := strconv.Atoi(timeend)

	rows, err := db.Query("Select id,date_start,date_end,time_start,time_end FROM reservation WHERE id_salle=?", roomID)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&datestart, &dateend, &timestart, &timeend)

		if err != nil {
			log.Fatal(err)
		}

		if dstart >= datestartint || dstart <= dateendint && dstart >= datestartint {
			dstartstr := strconv.Itoa(dstart)
			return "Cette salle n'est pas disponible le" + dstartstr
		}
		if dstart == datestartint {
			if dstart != dateendint {
				if hstart > timestartint {
					hstartstr := strconv.Itoa(hstart)
					return "Cette salle n'est pas disponible a" + hstartstr
				}
			} else {
				if hstart > timestartint && hstart < timeendint {
					hstartstr := strconv.Itoa(hstart)
					return "Cette salle n'est pas disponible a" + hstartstr
				}
			}
		}

		if dend >= datestartint || dend <= dateendint && dend > datestartint {
			dendstr := strconv.Itoa(dend)
			return "Cette salle n'est pas disponible le" + dendstr
		}
		if dend == datestartint {
			if dend != dateendint {
				if hend > timestartint {
					hendstr := strconv.Itoa(hend)
					return "Cette salle n'est pas disponible a" + hendstr
				}
			} else {
				if hend > timestartint && hend < timeendint {
					hendstr := strconv.Itoa(hend)
					return "Cette salle n'est pas disponible a" + hendstr
				}
			}
		}

	}
	return "ok"
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
