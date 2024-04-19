package verif

import (
	"regexp"
	"time"
)

func VerifDate(date string) string {
	correctDate := regexp.MustCompile(`^\d{2}-\d{2}-\d{4}$`)
	if !correctDate.MatchString(date) {
		err := "Format de date invalide. Veuillez entrer une date au format JJ-MM-AAAA."
		return err
	}
	return "ok"
}

func IsDateLogic(day int, month int, year int) string {
	monthDays := [][]int{{1, 31}, {2, 28}, {3, 31}, {4, 30}, {5, 31}, {6, 30}, {7, 31}, {8, 31}, {9, 30}, {10, 31}, {11, 30}, {12, 31}}

	//leap year
	if year%400 == 0 || (year%4 == 0 && year % 100 != 0) {
		monthDays[1][1] = 29
	} 

	if month < 1 || month > 12{
		err := "Le mois choisi n'est pas valide !"
		return err
	}

	if day < 1 || day > monthDays[month-1][1]{
		err := "Le jour choisi n'est pas valide !"
		return err
	}
	return "ok"
}

func IsBookLogic(startDay int, startMonth int, startYear int, endDay int, endMonth int, endYear int) string {
	startDate := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
  endDate := time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC)

	if endDate.Before(startDate) {
		err := "La date de fin ne peut pas être antérieure à la date de début."
		return err
	}
	return "ok"
}

func VerifTime(hour string) string {
	correctHour := regexp.MustCompile(`^\d{2}:\d{2}$`)
	if !correctHour.MatchString(hour) {
		err := "Format d'heure invalide. Veuillez entrer une heure au format HH:MM."
		return err
	}
	return "ok"
}