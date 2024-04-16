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

	for {
		choice := displayMenu()
		switch choice {
		case 1:
			GetDate()

			//room.ShowAvailableRooms(day, month, year) //change to room
		case 5:
			fmt.Println("A plus dans le bus !")
			return
		}
	}
}

func displayMenu() int {
	fmt.Println("Bienvenue dans le Service de Réservation en Ligne")
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

func GetDate(){
	for{
		date := ""
		fmt.Println("Quelle est la date de réservation (JJ-MM-AAAA) ?")
		fmt.Scan(&date)
		
		correctDate := regexp.MustCompile(`^\d{2}-\d{2}-\d{4}$`)
		if !correctDate.MatchString(date) {
			fmt.Println("Format de date invalide. Veuillez entrer une date au format JJ-MM-AAAA.")
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

		fmt.Println(day, month, year)
		break //need return later
	}
}