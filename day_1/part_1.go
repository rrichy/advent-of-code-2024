package day1

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/rrichy/advent-of-code-2024/utils"
)

func Part1() int {
	input := utils.ReadInput("input")

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

	sort.Ints(left)
	sort.Ints(right)

	var sum int

	for i, l := range left {
		r := right[i]
		sum += utils.AbsInt(r - l)
	}

	log.Println(sum)

	return sum
}
