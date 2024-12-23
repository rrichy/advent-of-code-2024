package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/rrichy/advent-of-code-2024/utils"
)

func Part2() int {
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

		if badCount >= 1 {
			for i := 0; i < len(levels); i++ {
				sublevels := []int{}
				for j, level := range levels {
					if i == j {
						continue
					}
					sublevels = append(sublevels, level)
				}

				if badLevelCount(sublevels) == 0 {
					sum++
					break
				}
			}
		}
	}

	log.Println(sum)

	return sum
}

func badLevelCount(levels []int) int {
	var prev, direction, count int

	for _, level := range levels {
		if prev == 0 {
			prev = level
			continue
		}

		if level == prev {
			count++
			continue
		}

		diff := level - prev
		if utils.AbsInt(diff) > 3 {
			count++
			continue
		}

		if diff < 0 {
			diff = -1
		} else {
			diff = 1
		}

		if direction == 0 {
			direction = diff
			prev = level
			continue
		}

		if direction != diff {
			count++
			continue
		}

		direction = diff
		prev = level
	}

	return count
}
