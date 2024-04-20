package verif

import (
	"fmt"
	"regexp"
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
	if year%400 == 0 || (year%4 == 0 && year % 100 != 0) {
		monthDays[1][1] = 29
	} 

	if month < 1 || month > 12{
		return "\033[31mLe mois choisi n'est pas valide !\033[0m"
	}

	if day < 1 || day > monthDays[month-1][1]{
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

func IsBookingLogic(startDay, startMonth, startYear, startHour, startMinut, endDay, endMonth, endYear, endHour, endMinut int) string {
	startDate := time.Date(startYear, time.Month(startMonth), startDay, startHour, startMinut, 0, 0, time.UTC)
  endDate := time.Date(endYear, time.Month(endMonth), endDay, endHour, endMinut, 0, 0, time.UTC)

	if endDate.Before(startDate) {
		return "\033[31mLa date de fin ne peut pas être antérieure à la date de début.\033[0m"
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