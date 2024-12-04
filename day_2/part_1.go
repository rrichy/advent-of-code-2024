package day2

import (
	"log"
	"strconv"
	"strings"

	"github.com/rrichy/advent-of-code-2024/utils"
)

func Part1() int {
	input := utils.ReadInput("./day_2/input")

	var sum int

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		split := strings.Fields(line)
		levels := []int{}

		for _, s := range split {
			cur, _ := strconv.Atoi(s)
			levels = append(levels, cur)
		}

		badCount := badLevelCount(levels)
		if badCount == 0 {
			sum++
		}
	}

	log.Println(sum)

	return sum
}
