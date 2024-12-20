package utils

import (
	"log"
	"os"
)

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ReadInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	input := string(data)

	return input
}

func SliceContains[T comparable](slice []T, s T) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
