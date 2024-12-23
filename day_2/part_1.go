package main

import (
	"log"
	"strconv"
	"strings"
)

func Part1() int {
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
