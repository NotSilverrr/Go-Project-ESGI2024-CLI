package verif

import "regexp"

func VerifDate(date string) string {
	correctDate := regexp.MustCompile(`^\d{2}-\d{2}-\d{4}$`)
	if !correctDate.MatchString(date) {
		err := "Format de date invalide. Veuillez entrer une date au format JJ-MM-AAAA."
		return err
	}
	return "ok"
}

func VerifTime(hour string) string {
	correctHour := regexp.MustCompile(`^\d{2}:\d{2}$`)
	if !correctHour.MatchString(hour) {
		err := "Format d'heure invalide. Veuillez entrer une heure au format HH:MM."
		return err
	}
	return "ok"
}