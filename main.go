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

			fmt.Printf("Elle commencera le %02d/%02d/%02d pour %02d:%02d\n", startDay, startMonth, startYear, startHour, startMinut)

			fmt.Printf("Elle se terminera le %02d/%02d/%02d pour %02d:%02d\n", endDay, endMonth, endYear, endHour, endMinut)

		case 3:
			book.CancelReservation(connec)

		case 4:
			choiceResa := displayMenuViewResa()
			if choiceResa == 1 {
				book.VisualizeReservationsRoom(connec)
			}
			if choiceResa == 2 {
				book.DisplayAllReservation(connec)
			}

		case 5:
			room.AddRoom(connec)

		case 6:
			book.ExportResaChoice(connec)

		case 7:
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
		fmt.Println("5. Créer une salle")
		fmt.Println("6. Exporter des réservations en JSON")
		fmt.Println("7. Quitter")
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

func displayMenuViewResa() int {

	for {
		fmt.Println("-----------------------------------------------------------")
		fmt.Println("1. Lister réservations par rapport a une salle")
		fmt.Println("2. Lister toute les réservations")
		fmt.Println("3. Retour")
		fmt.Println("")
		fmt.Println("Choisissez une option : ")
		var choice int
		fmt.Scan(&choice)

		if msg := correctChoiceViewResa(choice); msg != "ok" {
			fmt.Println(msg)
			continue
		}
		return choice
	}
}

func correctChoice(choice int) string {
	if choice < 1 || choice > 7 {
		return "\033[31mVeuillez choisir une option valide (entre 1 et 7).\033[0m"
	}
	return "ok"
}

func correctChoiceViewResa(choice int) string {
	if choice < 1 || choice > 3 {
		return "\033[31mVeuillez choisir une option valide (entre 1 et 3).\033[0m"
	}
	return "ok"
}
