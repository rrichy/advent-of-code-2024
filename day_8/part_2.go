package main

import (
	"fmt"
	"strings"
)

func Part2() int {
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
			antinodes[p1] = true
			for _, p2 := range positions {
				if p1 == p2 {
					continue
				}

				rise := p2.Y - p1.Y
				run := p2.X - p1.X

				a1 := Position{p1.X - run, p1.Y - rise}
				for {
					if a1.OutOfBounds(width, height) {
						break
					}

					antinodes[a1] = true
					a1 = Position{a1.X - run, a1.Y - rise}
				}

				a2 := Position{p2.X + run, p2.Y + rise}
				for {
					if a2.OutOfBounds(width, height) {
						break
					}

					antinodes[a2] = true
					a2 = Position{a2.X + run, a2.Y + rise}
				}
			}
		}
	}

	fmt.Printf("Antinodes: %d\n", len(antinodes))

	return len(antinodes)
}
