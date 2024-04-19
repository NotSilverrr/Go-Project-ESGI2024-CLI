package date

import (
	verif "Go-Project-ESGI2024-CLI/verif"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func GetDate() (int, int, int, int, int, int) {
	for {
		startDate := ""
		fmt.Printf("Quelle est la date de début de réservation (JJ-MM-AAAA) ?\n")
		fmt.Scan(&startDate)

		if msg := verif.VerifDate(startDate); msg != "ok" {
			fmt.Println(msg)
			continue
		}

		startResult := strings.Split(startDate, "-")

		startDay, err := strconv.Atoi(startResult[0])
		if err != nil {
			log.Fatal(err)
		}

		startMonth, err := strconv.Atoi(startResult[1])
		if err != nil {
			log.Fatal(err)
		}

		startYear, err := strconv.Atoi(startResult[2])
		if err != nil {
			log.Fatal(err)
		}

		if msg := verif.IsDateLogic(startDay, startMonth, startYear); msg != "ok" {
			fmt.Println(msg)
			continue
		}

		endDate := ""
		fmt.Printf("Quelle est la date de fin de réservation (JJ-MM-AAAA) ?\n")
		fmt.Scan(&endDate)

		if msg := verif.VerifDate(startDate); msg != "ok" {
			fmt.Println(msg)
			continue
		}

		endResult := strings.Split(endDate, "-")

		endDay, err := strconv.Atoi(endResult[0])
		if err != nil {
			log.Fatal(err)
		}

		endMonth, err := strconv.Atoi(endResult[1])
		if err != nil {
			log.Fatal(err)
		}

		endYear, err := strconv.Atoi(endResult[2])
		if err != nil {
			log.Fatal(err)
		}

		if msg := verif.IsDateLogic(endDay, endMonth, endYear); msg != "ok" {
			fmt.Println(msg)
			continue
		}
		
		return startDay, startMonth, startYear, endDay, endMonth, endYear
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
