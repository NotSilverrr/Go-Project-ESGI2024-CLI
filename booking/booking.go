package booking

import (
	room "Go-Project-ESGI2024-CLI/room"
	time "Go-Project-ESGI2024-CLI/time"
	"Go-Project-ESGI2024-CLI/verif"
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

func FormReservation(db *sql.DB) (int, int, int, int, int, int, int, int, int, int) {
	var ID string
	var roomID int
	roomVerif := "pasOK"
	room.DisplayRooms(db)
	for roomVerif != "ok" {
		fmt.Printf("Quelle salle voulez vous réserver?\n")
		fmt.Scan(&ID)
		roomVerif = verif.VerifIDRoom(ID, db)
	}
	roomID, err := strconv.Atoi(ID)

	if err != nil {
		log.Fatal(err)
	}

	startDay, startMonth, startYear, startHour, startMinute, endDay, endMonth, endYear, endHour, endMinute := time.GetBook()

	startDayStr := strconv.Itoa(startDay)
	startMonthStr := strconv.Itoa(startMonth)
	startYearStr := strconv.Itoa(startYear)
	startHourStr := strconv.Itoa(startHour)
	startMinuteStr := strconv.Itoa(startMinute)
	endDayStr := strconv.Itoa(endDay)
	endMonthStr := strconv.Itoa(endMonth)
	endYearStr := strconv.Itoa(endYear)
	endHourStr := strconv.Itoa(endHour)
	endMinuteStr := strconv.Itoa(endMinute)

	dStart := startDayStr + "-" + startMonthStr + "-" + startYearStr
	hStart := startHourStr + ":" + startMinuteStr
	dEnd := endDayStr + "-" + endMonthStr + "-" + endYearStr
	hEnd := endHourStr + ":" + endMinuteStr

	CreateReservation(roomID, dStart, dEnd, hStart, hEnd, db)

	return startDay, startMonth, startYear, startHour, startMinute, endDay, endMonth, endYear, endHour, endMinute
}

func CreateReservation(idSalle int, dstart string, dend string, hstart string, hend string, db *sql.DB) {

	res, err := db.Exec("INSERT INTO reservation (id_salle,date_start,date_end,time_start,time_end) VALUES (?,?,?,?,?)", idSalle, dstart, dend, hstart, hend)

	if err != nil {
		println(res)
		log.Fatal(err)
	}
	println("Votre réservation a bien été validé")
}

func CancelReservation(db *sql.DB) {
	var resID int
	fmt.Printf("Quelle réservation voulez vous annuler?\n")
	fmt.Scan(&resID)

	res, err := db.Exec("DELETE FROM reservation WHERE id=?", resID)

	if err != nil {
		println(res)
		log.Fatal(err)
	}
	fmt.Printf("La réservation %d a bien été annulé.\n", resID)
}


func VisualizeReservations(db *sql.DB) {
				var ID string
			roomVerif := "pasOK"

			room.DisplayRooms(db)

			for roomVerif != "ok" {
				fmt.Println("Quelle salle voulez-vous visualiser ?")
				fmt.Scan(&ID)
				roomVerif = verif.VerifIDRoom(ID, db)
			}

			intID, err := strconv.Atoi(ID)
			if err != nil {
				log.Fatal(err)
			}
		
			DisplayReservation(intID, db)
}

func DisplayReservation(roomID int, db *sql.DB) {
	// Get room name
	var roomName string

	err := db.QueryRow("SELECT name FROM room WHERE id = ?", roomID).Scan(&roomName)
	if err != nil {
		log.Fatal(err)
	}

	println("Réservations pour", roomName)

	// Get room reservations
	rows, err := db.Query("SELECT id, date_start, date_end, time_start, time_end FROM reservation WHERE id_salle = ?", roomID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var datestart, dateend, timestart, timeend string
		err := rows.Scan(&id, &datestart, &dateend, &timestart, &timeend)
		if err != nil {
			log.Fatal(err)
		}
		println(id, ".", datestart, timestart, " - ", dateend, timeend)
	}
}
