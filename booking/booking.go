package booking

import (
	time "Go-Project-ESGI2024-CLI/time"
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
	var roomID int
	fmt.Printf("Quelle salle voulez vous réserver??\n")
	fmt.Scan(&roomID)
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

	if err != nil {
		log.Fatal(err)
	}

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
