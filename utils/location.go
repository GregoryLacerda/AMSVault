package utils

import "time"

func GetDefaultLocation() *time.Location {

	loc, _ := time.LoadLocation("America/Sao_Paulo")

	return loc
}
