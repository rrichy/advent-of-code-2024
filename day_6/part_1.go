package main

import (
	"fmt"
	"strings"
)

type Position struct {
	x         int
	y         int
	direction string
}

func (p *Position) Coordinates() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func (p *Position) TurnRight() {
	if p.direction == "^" {
		p.direction = ">"
	} else if p.direction == ">" {
		p.direction = "v"
	} else if p.direction == "v" {
		p.direction = "<"
	} else if p.direction == "<" {
		p.direction = "^"
	}
}

func (p *Position) SetCoordinates(x int, y int) {
	p.x = x
	p.y = y
}

func (p *Position) Key() string {
	return fmt.Sprintf("%d,%d,%s", p.x, p.y, p.direction)
}

func (p *Position) GetNextPossibleCoordinates() (x, y int) {
	x = p.x
	y = p.y

	if p.direction == "^" {
		y = y - 1
	} else if p.direction == ">" {
		x = x + 1
	} else if p.direction == "v" {
		y = y + 1
	} else if p.direction == "<" {
		x = x - 1
	}

	return x, y
}

func Part1() int {
	grid := [][]string{}
	current_pos := Position{}

	for y, line := range strings.Split(input, "\n") {
		horizontal := strings.Split(line, "")
		grid = append(grid, horizontal)

		x := strings.Index(line, "^")
		if x != -1 {
			current_pos = Position{x, y, "^"}
		}
	}

	offgrid_index := len(grid)
	visited := map[string]bool{current_pos.Coordinates(): true}

	for {
		x, y := current_pos.GetNextPossibleCoordinates()

		if y < 0 || x >= offgrid_index || y >= offgrid_index || x < 0 {
			break
		}

		cell := grid[y][x]
		if cell == "#" {
			current_pos.TurnRight()
			continue
		}

		current_pos.SetCoordinates(x, y)
		visited[current_pos.Coordinates()] = true
	}

	fmt.Println(len(visited))

	return len(visited)
}
