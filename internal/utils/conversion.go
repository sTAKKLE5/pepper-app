package utils

import (
	"log"
	"strconv"
)

func ConvertStringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
