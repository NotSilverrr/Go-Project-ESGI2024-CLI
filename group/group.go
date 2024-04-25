package group

import (
	"Go-Project-ESGI2024-CLI/verif"
	"fmt"
)

func GetGroupSize() int {
	for {
		var groupSize int

		fmt.Println("Combien serez-vous ?")
		fmt.Scan(&groupSize)

		if msg := verif.VerifGroupSize(groupSize); msg != "ok" {
			fmt.Println(msg)
			continue
		}
	return groupSize
	}
}
