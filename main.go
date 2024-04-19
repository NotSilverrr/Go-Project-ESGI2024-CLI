package main

import (
	db "Go-Project-ESGI2024-CLI/db"
	room "Go-Project-ESGI2024-CLI/room"
	time "Go-Project-ESGI2024-CLI/time"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	connec := db.Db_opener()

	fmt.Println("Bienvenue dans le Service de Réservation en Ligne")
	for {
		choice := displayMenu()
		switch choice {
		case 1:
			room.ShowAvailableRooms(connec)

		case 2:
			startDay, startMonth, startYear, endDay, endMonth, endYear := time.GetBook()
			//startHour, startMinut := time.GetTime()

			// endDay, endMonth, endYear := time.GetDate(endContext)
			// endHour, endMinut := time.GetTime(endContext)

			fmt.Printf("Votre réservation commencera le %02d/%02d/%02d pour [time]\n", startDay, startMonth, startYear)//startHour, startMinut

			fmt.Printf("Votre réservation se terminera le %02d/%02d/%02d pour [time]\n", endDay, endMonth, endYear)//endHour, endMinut

		case 5:
			fmt.Println("A plus dans le bus !")
			db.Db_closer(connec)
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