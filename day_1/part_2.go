package main

import (
	"log"
	"strconv"
	"strings"
)

func Part2() int {
	var left, right []int

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		split := strings.Fields(line)

		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[1])

		left = append(left, l)
		right = append(right, r)
	}

	var sum int

	seen := make(map[int]*int)
	for _, l := range left {
		if seen[l] == nil {
			c := count(l, right)
			seen[l] = &c
		}

		sum += l * *seen[l]
	}

	log.Println(sum)

	return sum
}

func count(l int, right []int) int {
	var c int
	for _, r := range right {
		if r == l {
			c++
		}
	}
	return c
}
