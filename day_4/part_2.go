package day4

import (
	"fmt"
	"log"
	"strings"

	"github.com/rrichy/advent-of-code-2024/utils"
)

/*
M.S   S.S   M.M   S.M
.A.   .A.   .A.   .A.
M.S   M.M   S.S   S.M
*/

func Part2() int {
	input := utils.ReadInput("./day_4/input")

	lines := [][]string{}

	for _, line := range strings.Split(input, "\n") {
		lines = append(lines, strings.Split(line, ""))
	}

	count := 0
	n := len(lines)

	for i := 0; i <= n-3; i++ {
		for j := 0; j <= n-3; j++ {
			dacross := fmt.Sprintf("%s%s%s", lines[i][j], lines[i+1][j+1], lines[i+2][j+2])
			aacross := fmt.Sprintf("%s%s%s", lines[i][j+2], lines[i+1][j+1], lines[i+2][j])

			if (dacross == "MAS" && aacross == "MAS") || (dacross == "MAS" && aacross == "SAM") || (dacross == "SAM" && aacross == "MAS") || (dacross == "SAM" && aacross == "SAM") {
				count++
			}
		}
	}

	log.Println(count)

	return count
}
