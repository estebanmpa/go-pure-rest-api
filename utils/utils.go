package utils

import (
	"strconv"
)

func ToInt(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0 //TODO should return error 
	}
	return i
}
