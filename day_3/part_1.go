package day3

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/rrichy/advent-of-code-2024/utils"
)

func Part1() int {
	input := utils.ReadInput("./day_3/input")

	var sum int

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		re, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)

		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			l, _ := strconv.Atoi(match[1])
			r, _ := strconv.Atoi(match[2])

			sum += l * r
		}
	}

	log.Println(sum)

	return sum
}