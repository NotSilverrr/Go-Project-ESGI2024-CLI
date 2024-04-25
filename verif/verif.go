package verif

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func VerifDate(date string) string {
	correctDate := regexp.MustCompile(`^\d{2}-\d{2}-\d{4}$`)
	if !correctDate.MatchString(date) {
		return "\033[31mFormat de date invalide. Veuillez entrer une date au format JJ-MM-AAAA.\033[0m"
	}
	return "ok"
}

func IsDateLogic(day, month, year int) string {
	monthDays := [][]int{{1, 31}, {2, 28}, {3, 31}, {4, 30}, {5, 31}, {6, 30}, {7, 31}, {8, 31}, {9, 30}, {10, 31}, {11, 30}, {12, 31}}

	//leap year
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		monthDays[1][1] = 29
	}

	if month < 1 || month > 12 {
		return "\033[31mLe mois choisi n'est pas valide !\033[0m"
	}

	if day < 1 || day > monthDays[month-1][1] {
		return "\033[31mLe jour choisi n'est pas valide !\033[0m"
	}
	return "ok"
}

func IsDayInPast(startDay, startMonth, startYear int) string {
	currentDate := time.Now().UTC()
	startDate := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
	nextDate := currentDate.AddDate(0, 0, 1)

	if startDate.Before(currentDate) {
		err := fmt.Sprintf("\033[31mLes réservations sont ouvertes à partir du %s.\033[0m", nextDate.Format("02-01-2006"))
		return err
	}
	return "ok"
}

func IsEndDayBeforeStart(startDay, startMonth, startYear, endDay, endMonth, endYear int) string {
	startDate := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC)

	if endDate.Before(startDate) {
		return "\033[31mLa date de fin ne peut pas être antérieure à la date de début.\033[0m"
	}
	return "ok"
}

func IsBookingLogic(startDay, startMonth, startYear, startHour, startMinut, endDay, endMonth, endYear, endHour, endMinut int) string {
	startDate := time.Date(startYear, time.Month(startMonth), startDay, startHour, startMinut, 0, 0, time.UTC)
	endDate := time.Date(endYear, time.Month(endMonth), endDay, endHour, endMinut, 0, 0, time.UTC)

	if endDate.Before(startDate) {
		return "\033[31mLa date de fin ne peut pas être antérieure à la date de début.\033[0m"
	}

	if endDate.Equal(startDate) && endHour == startHour && endMinut == startMinut {
		return "\033[31mT'es bête ?\033[0m"
	}

	return "ok"
}

func VerifTime(hour string) string {
	correctHour := regexp.MustCompile(`^\d{2}:\d{2}$`)
	if !correctHour.MatchString(hour) {
		return "\033[31mFormat d'heure invalide. Veuillez entrer une heure au format HH:MM.\033[0m"
	}
	return "ok"
}

func IsTimeLogic(hour, minute int) string {
	if hour < 0 || hour > 23 {
		return "\033[31mL'heure doit être comprise entre 0 et 23.\033[0m"
	}
	if minute < 0 || minute > 59 {
		return "\033[31mLes minutes doivent être comprises entre 0 et 59.\033[0m"
	}
	return "ok"
}

func IsBookingTimeInPast(hour, minute int) string {
	currentTime := time.Now().UTC()
	bookingTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), hour, minute, 0, 0, time.UTC)

	if bookingTime.Before(currentTime) {
		return "\033[31mL'heure de réservation ne peut pas être dans le passé.\033[0m"
	}
	return "ok"
}

func VerifResa(roomID int, ystart int, yend int, mstart int, mend int, dstart int, dend int, hstart int, hend int, minstart int, minend int, db *sql.DB) string {
	var datestart string
	var dateend string
	var timestart string
	var timeend string

	datestartDATE := time.Date(ystart, time.Month(mstart), dstart, hstart, mstart, 0, 0, time.UTC)
	dateendDATE := time.Date(yend, time.Month(mend), dend, hend, minend, 0, 0, time.UTC)

	rows, err := db.Query("Select date_start,date_end,time_start,time_end FROM reservation WHERE id_salle=?", roomID)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&datestart, &dateend, &timestart, &timeend)

		if err != nil {
			log.Fatal(err)
		}

		dateStartTab := strings.Split(datestart, "-")
		dateEndTab := strings.Split(dateend, "-")
		timeStartTab := strings.Split(timestart, ":")
		timeEndTab := strings.Split(timeend, ":")

		yStartDB, _ := strconv.Atoi(dateStartTab[2])
		mStartDB, _ := strconv.Atoi(dateStartTab[1])
		dStartDB, _ := strconv.Atoi(dateStartTab[0])

		yEndDB, _ := strconv.Atoi(dateEndTab[2])
		mEndDB, _ := strconv.Atoi(dateEndTab[1])
		dEndDB, _ := strconv.Atoi(dateEndTab[0])

		hStartDB, _ := strconv.Atoi(timeStartTab[0])
		minStartDB, _ := strconv.Atoi(timeStartTab[1])

		hEndDB, _ := strconv.Atoi(timeEndTab[0])
		minEndDB, _ := strconv.Atoi(timeEndTab[1])

		datestartDB := time.Date(yStartDB, time.Month(mStartDB), dStartDB, hStartDB, minStartDB, 0, 0, time.UTC)
		dateendDB := time.Date(yEndDB, time.Month(mEndDB), dEndDB, hEndDB, minEndDB, 0, 0, time.UTC)

		if dateendDATE.After(datestartDB) && dateendDATE.Before(dateendDB) || datestartDATE.Before(dateendDB) && dateendDATE.After(dateendDB) || datestartDATE.After(datestartDB) && dateendDATE.Before(dateendDB) {
			return "cette réservation n'est pas disponible !"
		}

	}
	return "ok"
}

func VerifIDRoom(id string, groupSize int, db *sql.DB) string {
	idRoom, err := strconv.Atoi(id)

	if err != nil {
		fmt.Printf("\033[31mL'id de la salle doit être un chiffre\n\033[0m")
		return "pasOK"
	}

	var verif int = 0
	rows, err := db.Query("Select id from room where capacity >= ?", groupSize)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var idDB int
		if err := rows.Scan(&idDB); err != nil {
			log.Fatal(err)
		}
		if idDB == idRoom {
			verif = 1
		}
	}

	if verif == 1 {
		return "ok"
	} else {
		fmt.Printf("\033[31mLa salle %d n'existe pas !\n\033[0m", idRoom)
		return "pasOK"
	}
}

func VerifGroupSize(groupSize int) string {
	if groupSize < 1 {
		return "\033[31mVous ne pouvez pas être si peu !\033[0m"
		} 
		
		if groupSize > 1000 {
		return "\033[31mAucune de nos salles n'a une telle capacité, louez un château.\033[0m"
	}
	return "ok"
}

