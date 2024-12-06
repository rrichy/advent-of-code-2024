package main

import (
	"log"
	"os"

	day1 "github.com/rrichy/advent-of-code-2024/day_1"
	day2 "github.com/rrichy/advent-of-code-2024/day_2"
	day3 "github.com/rrichy/advent-of-code-2024/day_3"
	day4 "github.com/rrichy/advent-of-code-2024/day_4"
	day5 "github.com/rrichy/advent-of-code-2024/day_5"
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
	case "3":
		day3.Part1()
		day3.Part2()
	case "4":
		day4.Part1()
		day4.Part2()
	case "5":
		day5.Part1()
		day5.Part2()
	default:
		log.Println("Invalid day")
	}
}
