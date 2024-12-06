package day5

import (
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/rrichy/advent-of-code-2024/utils"
)

func Part2() int {
	input := utils.ReadInput("./day_5/input")

	sections := strings.Split(input, "\n\n")

	orders := make(map[string]Order, 0)

	for _, line := range strings.Split(sections[0], "\n") {
		temp := strings.Split(line, "|")
		l := temp[0]
		r := temp[1]

		if _, ok := orders[l]; !ok {
			orders[l] = Order{left: []string{}, right: []string{r}}
		} else {
			right := orders[l].right
			right = append(right, r)
			orders[l] = Order{left: orders[l].left, right: right}
		}

		if _, ok := orders[r]; !ok {
			orders[r] = Order{left: []string{l}, right: []string{}}
		} else {
			left := orders[r].left
			left = append(left, l)
			orders[r] = Order{left: left, right: orders[r].right}
		}
	}

	sum := 0
	for _, line := range strings.Split(sections[1], "\n") {
		temp := strings.Split(line, ",")

		for i := 0; i < len(temp)-1; i++ {
			l := temp[i]
			r := temp[i+1]

			if !utils.SliceContains(orders[l].right, r) {
				slices.SortFunc(temp, func(a, b string) int {
					if a == b {
						return 0
					}

					if utils.SliceContains(orders[a].right, b) {
						return -1
					}

					return 1
				})

				n, _ := strconv.Atoi(temp[len(temp)/2])
				sum += n
			}
		}
	}

	log.Println(sum)

	return sum
}
