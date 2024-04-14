package main

import (
	db "Go-Project-ESGI2024-CLI/bdd"
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db.Db_opener()

	for {
		choice := displayMenu()
		switch choice {
		case 1:
			date := ""
			fmt.Println("Quelle est la date de réservation (JJ-MM-AAAA) ?")
			fmt.Scan(&date)

			result := strings.Split(date, "-")
			day := result[0]
			month := result[1]
			year := result[2]
			println(day, month, year)

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

func Db_open() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydatabase")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
