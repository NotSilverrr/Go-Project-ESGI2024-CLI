package main

import (
	book "Go-Project-ESGI2024-CLI/booking"
	db "Go-Project-ESGI2024-CLI/db"
	room "Go-Project-ESGI2024-CLI/room"

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
			startDay, startMonth, startYear,
				startHour, startMinut,
				endDay, endMonth, endYear,
				endHour, endMinut := book.FormReservation(connec)

			fmt.Printf("Votre réservation commencera le %02d/%02d/%02d pour %02d:%02d\n", startDay, startMonth, startYear, startHour, startMinut)

			fmt.Printf("Votre réservation se terminera le %02d/%02d/%02d pour %02d:%02d\n", endDay, endMonth, endYear, endHour, endMinut)

		case 3:
			book.CancelReservation(connec)

		case 4:
			fmt.Println("Quelle salle souhaitez vous visualiser ?")
			fmt.Scan(&choice)
			book.DisplayReservation(choice, connec)
		case 5:
			fmt.Println("A plus dans le bus !")
			db.Db_closer(connec)
			return
		}
	}
}

func displayMenu() int {

	for {
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

		if msg := correctChoice(choice); msg != "ok" {
			fmt.Println(msg)
			continue
		}
		return choice
	}
}

func correctChoice(choice int) string {
	if choice < 1 || choice > 5 {
		return "\033[31mVeuillez choisir une option valide (entre 1 et 5).\033[0m"
	}
	return "ok"
}
