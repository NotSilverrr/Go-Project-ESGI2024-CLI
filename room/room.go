package salle

import (
	"Go-Project-ESGI2024-CLI/group"
	time "Go-Project-ESGI2024-CLI/time"
	"Go-Project-ESGI2024-CLI/verif"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ShowAvailableRooms(db *sql.DB) {
	//user choose start date and hour for is reservation
	startDay, startMonth, startYear,
	startHour, startMinut,
	endDay, endMonth, endYear,
	endHour, endMinut := time.GetBook()

	capacity := group.GetGroupSize()

	var result string

	fmt.Printf("Vous avez choisi une réservation commençant le %02d/%02d/%02d à %02d:%02d et se terminant le %02d/%02d/%02d à %02d:%02d.\nVous souhaitez une salle pouvant accueillir au minimum %d personnes. \nLes salles disponibles sont les suivantes : \n", startDay, startMonth, startYear, startHour, startMinut, endDay, endMonth, endYear, endHour, endMinut, capacity)

	rows, err := db.Query("Select id, name, capacity from room where capacity >= ?", capacity)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var capacity int
		if err := rows.Scan(&id, &name, &capacity); err != nil {
			log.Fatal(err)
		}

		result = verif.VerifResa(id, startYear, endYear, startMonth, endMonth, startDay, endDay, startHour, endHour, startMinut, endMinut, db)

		if result == "ok" {
			fmt.Printf("%s -- %d places\n", name, capacity)
		}
	}
}

func DisplayRooms(db *sql.DB) {
	rows, err := db.Query("Select id,name from room")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var id string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s. %s\n", id, name)
	}
}