package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/rrichy/advent-of-code-2024/utils"
)

func Part1() int {
	lines := utils.New2DStringMatrix(input)

	count := 0
	n := len(lines)

	for i, chars := range lines {
		line := strings.Join(chars, "")

		count += strings.Count(line, "XMAS")
		count += strings.Count(line, "SAMX")

		for j := 0; j < n; j++ {
			if i <= n-4 {
				down := fmt.Sprintf("%s%s%s%s", lines[i][j], lines[i+1][j], lines[i+2][j], lines[i+3][j])

				if down == "XMAS" || down == "SAMX" {
					count++
				}

				if j <= n-4 {
					dacross := fmt.Sprintf("%s%s%s%s", lines[i][j], lines[i+1][j+1], lines[i+2][j+2], lines[i+3][j+3])
					if dacross == "XMAS" || dacross == "SAMX" {
						count++
					}
				}
			}

			if i >= 3 && j <= n-4 {
				uacross := fmt.Sprintf("%s%s%s%s", lines[i][j], lines[i-1][j+1], lines[i-2][j+2], lines[i-3][j+3])
				if uacross == "XMAS" || uacross == "SAMX" {
					count++
				}
			}
		}
	}

	log.Println(count)

	return count
}
