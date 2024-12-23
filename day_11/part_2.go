package main

import (
	"log"
	"strings"
)

func Count(m map[string]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}

	return sum
}

func Part2() int {
	stones := strings.Split(input, " ")

	blinkCount := 75
	history := map[string][]string{}
	count := map[string]int{}

	for _, stone := range stones {
		if count[stone] == 0 {
			count[stone] = 1
		} else {
			count[stone] += 1
		}
	}

	for i := 0; i < blinkCount; i++ {
		_stones := []string{}
		_count := map[string]int{}

		for _, stone := range stones {
			c := count[stone]

			splitted := Split(&history, stone)

			for _, split := range splitted {
				currentCount := _count[split]
				if currentCount == 0 {
					_count[split] = c
					_stones = append(_stones, split)
				} else {
					_count[split] += c
				}
			}
		}

		stones = _stones

		count = _count

		log.Printf("Count of stones in %d blink: %d", i+1, Count(count))
	}

	return Count(count)
}
