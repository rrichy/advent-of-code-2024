package day8

import (
	"fmt"
	"strings"

	"github.com/rrichy/advent-of-code-2024/utils"
)

type Position struct {
	X int
	Y int
}

func (p *Position) OutOfBounds(width, height int) bool {
	return p.X < 0 || p.Y < 0 || p.X >= width || p.Y >= height
}

func Part1() int {
	input := utils.ReadInput("day_8/input")
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(strings.Split(lines[0], ""))

	antennas := map[string][]Position{}

	for y, line := range lines {
		row := strings.Split(line, "")

		for x, cell := range row {
			if cell != "." {
				positions := antennas[cell]
				if positions == nil {
					antennas[cell] = []Position{{x, y}}
				} else {
					antennas[cell] = append(positions, Position{x, y})
				}
			}
		}
	}

	antinodes := map[Position]bool{}
	for _, positions := range antennas {
		for _, p1 := range positions {
			for _, p2 := range positions {
				if p1 == p2 {
					continue
				}

				rise := p2.Y - p1.Y
				run := p2.X - p1.X

				a1 := Position{p1.X - run, p1.Y - rise}
				a2 := Position{p2.X + run, p2.Y + rise}

				if !a1.OutOfBounds(width, height) {
					antinodes[a1] = true
				}

				if !a2.OutOfBounds(width, height) {
					antinodes[a2] = true
				}
			}
		}
	}

	fmt.Printf("Antinodes: %d\n", len(antinodes))

	return len(antinodes)
}
