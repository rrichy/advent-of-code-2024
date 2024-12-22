package day11

import (
	"log"
	"strconv"
	"strings"

	"github.com/rrichy/advent-of-code-2024/utils"
)

func Part1() int {
	input := utils.ReadInput("day_11/input")
	stones := strings.Split(input, " ")

	blinkCount := 25
	history := map[string][]string{}

	for i := 0; i < blinkCount; i++ {
		_stones := []string{}
		for _, stone := range stones {
			_stones = append(_stones, Split(&history, stone)...)
		}

		stones = _stones
		// log.Print(stones)
		log.Printf("Count of stones in %d blink: %d", i+1, len(stones))
	}

	return 0
}

func Split(history *map[string][]string, stone string) []string {
	prev := (*history)[stone]
	if prev != nil {
		return prev
	}

	if stone == "0" {
		return []string{"1"}
	}

	length := len(stone)
	if length%2 == 0 {
		var l, r string
		for i, d := range stone {
			if i < length/2 {
				l += string(d)
			} else {
				r += string(d)
			}
		}

		s, _ := strconv.Atoi(r)
		r = strconv.Itoa(s)

		(*history)[stone] = []string{l, r}

		return []string{l, r}
	}

	s, _ := strconv.Atoi(stone)

	return []string{strconv.Itoa(2024 * s)}
}
