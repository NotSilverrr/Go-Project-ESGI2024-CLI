package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db()

	for {
		choice := displayMenu()
		switch choice {
		case 5:
			break
		}
		fmt.Println("A plus dans le bus !")
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


func db() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydatabase")

	if err != nil {
		log.Fatal(err)
	}
	//rows, err := db.Query("SELECT id,name FROM test")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	err = rows.Scan(&id, &name)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
		//fmt.Println("Ligne :", id, name)
	//}

	defer db.Close()
}