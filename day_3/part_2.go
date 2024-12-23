package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Part2() int {
	var sum int

	muls := []string{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

		matches := re.FindAllString(line, -1)
		muls = append(muls, matches...)
	}

	do := true
	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, mul := range muls {
		if mul == "do()" {
			do = true
		} else if mul == "don't()" {
			do = false
		} else if do {
			match := mulRe.FindStringSubmatch(mul)
			l, _ := strconv.Atoi(match[1])
			r, _ := strconv.Atoi(match[2])

			sum += l * r
		}
	}

	log.Println(sum)

	return sum
}
