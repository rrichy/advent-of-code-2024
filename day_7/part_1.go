package day7

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/rrichy/advent-of-code-2024/utils"
)

func Part1() int {
	input := utils.ReadInput("day_7/input")

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
		combinations := int(math.Pow(2, float64(operator_count)))
		_sum := values[0]
		for i := 0; i < combinations; i++ {
			bitwise := strings.Split(fmt.Sprintf("%0*b", operator_count, i), "")

			for bit_index, bit := range bitwise {
				if bit == "1" {
					_sum += values[bit_index+1]
				} else {
					_sum *= values[bit_index+1]
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
