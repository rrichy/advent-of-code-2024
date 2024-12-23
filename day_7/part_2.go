package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part2() int {
	total := 0
	for _, _line := range strings.Split(input, "\n") {
		line := strings.Split(_line, ": ")

		sum, _ := strconv.Atoi(line[0])

		values := []int{}
		for _, _v := range strings.Split(line[1], " ") {
			v, _ := strconv.Atoi(_v)
			values = append(values, v)
		}

		operator_count := len(values) - 1
		combinations := int(math.Pow(3, float64(operator_count)))
		_sum := values[0]
		for i := 0; i < combinations; i++ {
			tritwise := strings.Split(fmt.Sprintf("%0*s", operator_count, strconv.FormatInt(int64(i), 3)), "")

			for trit_index, trit := range tritwise {
				if trit == "1" {
					_sum += values[trit_index+1]
				} else if trit == "2" {
					digit_string := fmt.Sprintf("%d%d", _sum, values[trit_index+1])
					_sum, _ = strconv.Atoi(digit_string)
				} else {
					_sum *= values[trit_index+1]
				}
			}

			if _sum == sum {
				total += sum
				break
			}

			_sum = values[0]
		}
	}

	fmt.Println(total)

	return total
}
