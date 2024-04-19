package date

import (
	verif "Go-Project-ESGI2024-CLI/verif"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func GetDate(context string) (int, int, int) {
	for {
		date := ""
		fmt.Printf("Quelle est la date %s de réservation (JJ-MM-AAAA) ?\n", context)
		fmt.Scan(&date)

		if msg := verif.VerifDate(date); msg != "ok" {
			fmt.Println(msg)
			continue
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

		if msg := verif.IsDateLogic(day, month, year); msg != "ok" {
			fmt.Println(msg)
			continue
		}
		
		return day, month, year
	}
}

func GetTime(context string) (int, int) {
	for {
		startHour := ""
		fmt.Printf("Quelle est l'heure %s de réservation (HH:MM) ?\n", context)
		fmt.Scan(&startHour)

		if msg := verif.VerifTime(startHour); msg != "ok" {
			fmt.Println(msg)
			continue
		}

		result := strings.Split(startHour, ":")

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

//todo
//verification of correct duration beetween the two dates
