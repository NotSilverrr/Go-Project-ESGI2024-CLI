package main

import (
	db "Go-Project-ESGI2024-CLI/bdd"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	connec := db.Db_opener()
	println(connec)

	fmt.Println("Bienvenue dans le Service de Réservation en Ligne")
	for {
		choice := displayMenu()
		switch choice {
		case 1:
			startContext := "de début"
			endContext := "de fin"

			startDay, startMonth, startYear := GetDate(startContext)
			startHour, startMinut := getTime(startContext)
			
			endDay, endMonth, endYear := GetDate(endContext)
			endHour, endMinut := getTime(endContext)

			fmt.Printf("Votre réservation commencera le %02d/%02d/%02d pour %02d:%02d\n", startDay, startMonth, startYear, startHour, startMinut)

			fmt.Printf("Votre réservation se terminera le %02d/%02d/%02d pour %02d:%02d\n", endDay, endMonth, endYear, endHour, endMinut)

		case 2:
			startContext := "de début"
			day, month, year := GetDate(startContext)
			hour, minut := getTime(startContext)
			fmt.Printf("Vous avez réservé le %02d/%02d/%02d pour %02d:%02d\n", day, month, year, hour, minut)

		case 5:
			fmt.Println("A plus dans le bus !")
			return
		}
	}
}

func displayMenu() int {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("1. Lister les salles disponibles")
	fmt.Println("2. Créer une réservation")
	fmt.Println("3. Annuler une réservation")
	fmt.Println("4. Visualiser les réservations")
	fmt.Println("5. Quitter")
	fmt.Println("")
	fmt.Println("Choisissez une option : ")
	var choice int
	fmt.Scan(&choice)
	return choice
}

func GetDate(context string) (int, int, int){
	for{
		date := ""
		fmt.Printf("Quelle est la date %s de réservation (JJ-MM-AAAA) ?\n", context)
		fmt.Scan(&date)

		if msg := verifDate(date); msg != "ok" {
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

		return day, month, year
	}
}

func getTime(context string) (int, int){
	for{
		startHour := ""
		fmt.Printf("Quelle est l'heure %s de réservation (HH:MM) ?\n", context)
		fmt.Scan(&startHour)

		if msg := verifTime(startHour); msg != "ok" {
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

func verifDate(date string) string {
	correctDate := regexp.MustCompile(`^\d{2}-\d{2}-\d{4}$`)
	if !correctDate.MatchString(date) {
		err := "Format de date invalide. Veuillez entrer une date au format JJ-MM-AAAA."
		return err
	}
	return "ok"
}

func verifTime(startHour string) string {
	correctHour := regexp.MustCompile(`^\d{2}:\d{2}$`)
	if !correctHour.MatchString(startHour) {
		err := "Format d'heure invalide. Veuillez entrer une heure au format HH:MM."
		return err
	}
	return "ok"
}
