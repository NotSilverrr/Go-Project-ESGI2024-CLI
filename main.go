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
			day, month, year := GetDate()
			hour, minut := getTime()
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

func GetDate() (int, int, int){
	for{
		date := ""
		fmt.Println("Quelle est la date de réservation (JJ-MM-AAAA) ?")
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

func getTime() (int, int){
	for{
		startHour := ""
		fmt.Println("Quelle est l'heure du début de réservation (HH:MM) ?")
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

func verifTime(startHour string) string{
	correctHour := regexp.MustCompile(`^\d{2}:\d{2}$`)
	if !correctHour.MatchString(startHour) {
		err := "Format d'heure invalide. Veuillez entrer une heure au format HH:MM."
		return err
	}
	return "ok"
}