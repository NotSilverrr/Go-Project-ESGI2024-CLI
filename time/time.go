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
		fmt.Printf("Quelle est la date de début de réservation (JJ-MM-AAAA) ?\n")
		startDay, startMonth, startYear := verif.ConvertStringToInt()

		if msg := verif.IsDateLogic(startDay, startMonth, startYear); msg != "ok" {
			fmt.Println(msg)
			continue
		}

		for {
			fmt.Printf("Quelle est la date de fin de réservation (JJ-MM-AAAA) ?\n")
			endDay, endMonth, endYear := verif.ConvertStringToInt()

			if msg := verif.IsDateLogic(endDay, endMonth, endYear); msg != "ok" {
				fmt.Println(msg)
				continue
			}

			return startDay, startMonth, startYear, endDay, endMonth, endYear
		}
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
