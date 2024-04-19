package date

import (
	verif "Go-Project-ESGI2024-CLI/verif"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func GetBook() (int, int, int, int, int, int, int, int, int, int) {
	startContext := "de début"
	endContext := "de fin"

	for {
		startDay, startMonth, startYear := GetDate(startContext)

		if msg := verif.IsDateLogic(startDay, startMonth, startYear); msg != "ok" {
			fmt.Println(msg)
			continue
		}

		startHour, startMinut := GetTime(startContext)

		for {
			endDay, endMonth, endYear := GetDate(endContext)

			if msg := verif.IsDateLogic(endDay, endMonth, endYear); msg != "ok" {
				fmt.Println(msg)
				continue
			}

			endHour, endMinut := GetTime(endContext)

			return startDay, startMonth, startYear, startHour, startMinut, endDay, endMonth, endYear, endHour, endMinut
		}
	}
}

func GetDate(dayContext string) (int,int,int){
	date := ""
	fmt.Printf("Quelle est la date %s de réservation (JJ-MM-AAAA) ?\n", dayContext)
	fmt.Scan(&date)
	
	if msg := verif.VerifDate(date); msg != "ok" {
		fmt.Println(msg)
	}
	
	result := strings.Split(date, "-")
	
	day, err := strconv.Atoi(result[0])
	if err != nil {
		log.Fatal(err)
	}
	
	month, err := strconv.Atoi(result[1])
	if err != nil {
		log.Fatal(err)
	}
	
	year, err := strconv.Atoi(result[2])
	if err != nil {
		log.Fatal(err)
	}
	return day, month, year
}

func GetTime(context string) (int, int) {
	for {
		time := ""
		fmt.Printf("Quelle est l'heure %s de réservation (HH:MM) ?\n", context)
		fmt.Scan(&time)

		if msg := verif.VerifTime(time); msg != "ok" {
			fmt.Println(msg)
			continue
		}

		result := strings.Split(time, ":")

		hour, err := strconv.Atoi(result[0])
		if err != nil {
			log.Fatal(err)
		}

		minut, err := strconv.Atoi(result[1])
		if err != nil {
			log.Fatal(err)
		}

		return hour, minut
	}
}