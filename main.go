package main

import (
	"log"
	"os"

	day1 "github.com/rrichy/advent-of-code-2024/day_1"
	day2 "github.com/rrichy/advent-of-code-2024/day_2"
)

func main() {
	args := os.Args[1:]

	switch args[0] {
	case "1":
		day1.Part1()
		day1.Part2()
	case "2":
		day2.Part1()
		day2.Part2()
	default:
		log.Println("Invalid day")
	}
}
