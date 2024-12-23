package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return i
}

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

func New2DStringMatrix(s string) [][]string {
	temp := [][]string{}
	for _, line := range strings.Split(s, "\n") {
		temp = append(temp, strings.Split(line, ""))
	}

	return temp
}
