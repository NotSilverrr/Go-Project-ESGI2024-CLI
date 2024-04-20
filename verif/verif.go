package verif

import (
	"regexp"
	"time"
)

func VerifDate(date string) string {
	correctDate := regexp.MustCompile(`^\d{2}-\d{2}-\d{4}$`)
	if !correctDate.MatchString(date) {
		err := "\033[31mFormat de date invalide. Veuillez entrer une date au format JJ-MM-AAAA.\033[0m"
		return err
	}
	return "ok"
}

func IsDateLogic(day, month, year int) string {
	monthDays := [][]int{{1, 31}, {2, 28}, {3, 31}, {4, 30}, {5, 31}, {6, 30}, {7, 31}, {8, 31}, {9, 30}, {10, 31}, {11, 30}, {12, 31}}

	//leap year
	if year%400 == 0 || (year%4 == 0 && year % 100 != 0) {
		monthDays[1][1] = 29
	} 

	if month < 1 || month > 12{
		err := "\033[31mLe mois choisi n'est pas valide !\033[0m"
		return err
	}

	if day < 1 || day > monthDays[month-1][1]{
		err := "\033[31mLe jour choisi n'est pas valide !\033[0m"
		return err
	}
	return "ok"
}

func IsBookingDayInPast(startDay, startMonth, startYear int) string {
	currentDate := time.Now().UTC()
	startDate := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)

	if startDate.Before(currentDate) {
		err := "\033[31mLa date de début de réservation ne peut pas être dans le passé.\033[0m"
		return err
	}
	return "ok"
}

func IsBookingLogic(startDay, startMonth, startYear, endDay, endMonth, endYear int) string {
	startDate := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
  endDate := time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC)

	if endDate.Before(startDate) {
		err := "\033[31mLa date de fin ne peut pas être antérieure à la date de début.\033[0m"
		return err
	}
	return "ok"
}

func VerifTime(hour string) string {
	correctHour := regexp.MustCompile(`^\d{2}:\d{2}$`)
	if !correctHour.MatchString(hour) {
		err := "\033[31mFormat d'heure invalide. Veuillez entrer une heure au format HH:MM.\033[0m"
		return err
	}
	return "ok"
}