package booking

import (
	"Go-Project-ESGI2024-CLI/group"
	room "Go-Project-ESGI2024-CLI/room"
	time "Go-Project-ESGI2024-CLI/time"
	"Go-Project-ESGI2024-CLI/verif"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type resa struct {
	Id         int
	Id_salle   int
	Time_start string
	Time_end   string
	Date_start string
	Date_end   string
}

func FormReservation(db *sql.DB) (int, int, int, int, int, int, int, int, int, int) {
	var ID string
	var roomID int
	var startDay, startMonth, startYear, startHour, startMinute, endDay, endMonth, endYear, endHour, endMinute int
	roomVerif := "pasOK"
	timeVerif := "pasOK"

	groupSize := group.GetGroupSize()
	room.DisplayRoom(groupSize, db)

	for roomVerif != "ok" {
		fmt.Printf("Quelle salle voulez vous réserver?\n")
		fmt.Scan(&ID)
		roomVerif = verif.VerifIDRoom(ID, groupSize, db)
	}
	roomID, err := strconv.Atoi(ID)

	if err != nil {
		log.Fatal(err)
	}

	for timeVerif != "ok" {
		startDay, startMonth, startYear, startHour, startMinute, endDay, endMonth, endYear, endHour, endMinute = time.GetBook()
		timeVerif = verif.VerifResa(roomID, startYear, endYear, startMonth, endMonth, startDay, endDay, startHour, endHour, startMinute, endMinute, db)
		if timeVerif != "ok" {
			println(timeVerif)
		}
	}

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
	fmt.Printf("\033[32mVotre réservation a bien été validée\n\033[0m")
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
	fmt.Printf("\033[32mLa réservation %d a bien été annulée.\033[0m\n", resID)
}

func VisualizeReservationsRoom(db *sql.DB) {
	var ID string
	roomVerif := "pasOK"

	room.DisplayRoom(0, db)

	for roomVerif != "ok" {
		fmt.Println("Quelle salle voulez-vous visualiser ?")
		fmt.Scan(&ID)
		roomVerif = verif.VerifIDRoom(ID, 0, db)
	}

	intID, err := strconv.Atoi(ID)
	if err != nil {
		log.Fatal(err)
	}

	DisplayReservation(intID, db)
}

func DisplayAllReservation(db *sql.DB) {
	println("Réservations :")

	// Get room reservations
	rows, err := db.Query("SELECT id, id_salle,date_start, date_end, time_start, time_end FROM reservation")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, id_salle int
		var datestart, dateend, timestart, timeend string
		err := rows.Scan(&id, &id_salle, &datestart, &dateend, &timestart, &timeend)

		if err != nil {
			log.Fatal(err)
		}

		var roomName string

		err = db.QueryRow("SELECT name FROM room WHERE id = ?", id_salle).Scan(&roomName)

		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			log.Fatal(err)
		}
		println(id, ".", roomName, datestart, timestart, " - ", dateend, timeend)
	}
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

func ExportResaChoice(db *sql.DB) {
	var roomChoice string
	var roomChoiceInt int
	roomVerif := "pasok"
	room.DisplayRoom(0, db)

	for roomVerif != "ok" {
		fmt.Printf("De quelle salle voulez vous exporter les réservations?\n")
		fmt.Scan(&roomChoice)
		roomVerif = verif.VerifIDRoom(roomChoice, 0, db)
	}
	roomChoiceInt, err := strconv.Atoi(roomChoice)

	if err != nil {
		log.Fatal(err)
	}

	ExportRoomResa(roomChoiceInt, db)
}

func ExportRoomResa(roomID int, db *sql.DB) {
	var resaArray []resa
	rows, err := db.Query("SELECT id, id_salle, date_start, date_end, time_start, time_end FROM reservation WHERE id_salle = ?", roomID)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var resaTemp resa
		err := rows.Scan(&resaTemp.Id, &resaTemp.Id_salle, &resaTemp.Date_start, &resaTemp.Date_end, &resaTemp.Time_start, &resaTemp.Time_end)
		if err != nil {
			log.Fatal(err)
		}
		resaArray = append(resaArray, resaTemp)
	}
	resaJSON, err := json.Marshal(resaArray)

	if err != nil {
		log.Fatal(err)
	}

	var roomName string

	err = db.QueryRow("SELECT name FROM room WHERE id = ?", roomID).Scan(&roomName)

	if err != nil {
		log.Fatal(err)
	}

	filename := "./JSONoutput/" + roomName

	err = ioutil.WriteFile(filename, resaJSON, 0644)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Votre export a bien été enregistré a %s !\n", filename)
}
