package main

import (
	"fmt"
	"strings"
)

func Part2() int {
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

	offgrid_y := len(grid)
	offgrid_x := len(grid[0])
	visited := map[string]string{current_pos.Coordinates(): current_pos.direction}
	loops := map[string]bool{}

	for {
		x := current_pos.x
		y := current_pos.y

		if current_pos.direction == "^" {
			y = y - 1
		} else if current_pos.direction == ">" {
			x = x + 1
		} else if current_pos.direction == "v" {
			y = y + 1
		} else if current_pos.direction == "<" {
			x = x - 1
		}

		if y < 0 || x >= offgrid_x || y >= offgrid_y || x < 0 {
			break
		}

		cell := grid[y][x]
		if cell == "#" {
			current_pos.TurnRight()
			continue
		}

		if boundToLoop(&grid, &visited, current_pos) {
			loops[fmt.Sprintf("%d,%d", x, y)] = true
		}

		current_pos.SetCoordinates(x, y)
		visited[current_pos.Coordinates()] = current_pos.direction
	}

	fmt.Println(len(loops))

	return len(loops)
}

func boundToLoop(grid *[][]string, visited *map[string]string, current_position Position) bool {
	block_x, block_y := current_position.GetNextPossibleCoordinates()
	if _, ok := (*visited)[fmt.Sprintf("%d,%d", block_x, block_y)]; ok {
		return false
	}

	localVisited := map[string]bool{}

	for k, v := range *visited {
		localVisited[fmt.Sprintf("%s,%s", k, v)] = true
	}

	current_position.TurnRight()
	offgrid_y := len(*grid)
	offgrid_x := len((*grid)[0])

	for {
		x, y := current_position.GetNextPossibleCoordinates()

		if y < 0 || x >= offgrid_x || y >= offgrid_y || x < 0 {
			return false
		}

		cell := (*grid)[y][x]
		if cell == "#" || (x == block_x && y == block_y) {
			current_position.TurnRight()
			continue
		}

		current_position.SetCoordinates(x, y)
		if _, ok := localVisited[current_position.Key()]; ok {
			return true
		}

		localVisited[current_position.Key()] = true
	}
}
