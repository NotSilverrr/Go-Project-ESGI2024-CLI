package time

import (
	verif "Go-Project-ESGI2024-CLI/verif"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func GetBook() (int, int, int, int, int, int, int, int, int, int) {
  startContext := "de début"
  endContext := "de fin"

  for {
    startDay, startMonth, startYear := getDate(startContext)

    // Error in Day/Month number
    if msg := verif.IsDateLogic(startDay, startMonth, startYear); msg != "ok" {
      fmt.Println(msg)
      continue
    }
        
    // Error if Day in the past
    if msg := verif.IsDayInPast(startDay, startMonth, startYear); msg != "ok" {
      fmt.Println(msg)
      continue
    }
        
    for {
      startHour, startMinute := getTime(startContext)

      // Error in Hour/Minute number
      if msg := verif.IsTimeLogic(startHour, startMinute); msg != "ok" {
        fmt.Println(msg)
        continue
      }

			for{
				endDay, endMonth, endYear := getDate(endContext)
							
				// Error in Day/Month number
				if msg := verif.IsDateLogic(endDay, endMonth, endYear); msg != "ok" {
					fmt.Println(msg)
					continue
				}
	
				// Error if Day in the past
				if msg := verif.IsDayInPast(endDay, endMonth, endYear); msg != "ok" {
					fmt.Println(msg)
					continue
				}
							
				for {
					endHour, endMinute := getTime(endContext)
	
					// Error in Hour/Minute number
					if msg := verif.IsTimeLogic(endHour, endMinute); msg != "ok" {
						fmt.Println(msg)
						continue
					}
	
					// Error if end time is before start time
					if msg := verif.IsBookingLogic(startDay, startMonth, startYear, startHour, startMinute, endDay, endMonth, endYear, endHour, endMinute); msg != "ok" {
						fmt.Println(msg)
						continue
					}
	
					return startDay, startMonth, startYear, startHour, startMinute, endDay, endMonth, endYear, endHour, endMinute
				}
			}
    }
  }
}

func getDate(dayContext string) (int,int,int){
	for {
		date := ""
		fmt.Printf("Quelle est la date %s de réservation (JJ-MM-AAAA) ?\n", dayContext)
		fmt.Scan(&date)
		
		if msg := verif.VerifDate(date); msg != "ok" {
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

func getTime(context string) (int, int) {
	for {
		time := ""
		fmt.Printf("Quelle est l'heure %s de réservation (HH:MM) ?\n", context)
		fmt.Scan(&time)

		if msg := verif.VerifTime(time); msg != "ok" {
			fmt.Println(msg)
			continue
		}

		result := strings.Split(time, ":")

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
