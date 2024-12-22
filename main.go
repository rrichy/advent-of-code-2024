package main

import (
	"log"
	"os"

	day1 "github.com/rrichy/advent-of-code-2024/day_1"
	day10 "github.com/rrichy/advent-of-code-2024/day_10"
	day11 "github.com/rrichy/advent-of-code-2024/day_11"
	day2 "github.com/rrichy/advent-of-code-2024/day_2"
	day3 "github.com/rrichy/advent-of-code-2024/day_3"
	day4 "github.com/rrichy/advent-of-code-2024/day_4"
	day5 "github.com/rrichy/advent-of-code-2024/day_5"
	day6 "github.com/rrichy/advent-of-code-2024/day_6"
	day7 "github.com/rrichy/advent-of-code-2024/day_7"
	day8 "github.com/rrichy/advent-of-code-2024/day_8"
	day9 "github.com/rrichy/advent-of-code-2024/day_9"
	"github.com/rrichy/advent-of-code-2024/utils"
)

func main() {
	args := os.Args[1:]

	switch args[0] {
	// Execute the code for the specified day
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
	case "6":
		day6.Part1()
		day6.Part2()
	case "7":
		day7.Part1()
		day7.Part2()
	case "8":
		day8.Part1()
		day8.Part2()
	case "9":
		day9.Part1()
		day9.Part2()
	case "10":
		day10.Part1()
		day10.Part2()
	case "11":
		day11.Part1()
		day11.Part2()
	case "10-animate":
		day10.Animate2()

	// Unrelated commands
	case "check":
		utils.Checker()
	default:
		log.Println("Invalid day")
	}
}
