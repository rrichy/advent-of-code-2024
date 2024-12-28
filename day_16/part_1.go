package main

import (
	"log"
	"math"
	"strings"
	"time"

	"github.com/rrichy/advent-of-code-2024/utils"
)

type Coordinate utils.Coordinate

type Reindeer struct {
	Coordinate
	Direction utils.Direction
}

type Movement struct {
	Coordinate
	Direction utils.Direction
	Score     int
}

func (r *Reindeer) RotateLeft() {
	switch r.Direction {
	case utils.Up:
		r.Direction = utils.Left
	case utils.Down:
		r.Direction = utils.Right
	case utils.Left:
		r.Direction = utils.Down
	case utils.Right:
		r.Direction = utils.Up
	}
}

func (r *Reindeer) RotateRight() {
	switch r.Direction {
	case utils.Up:
		r.Direction = utils.Right
	case utils.Down:
		r.Direction = utils.Left
	case utils.Left:
		r.Direction = utils.Up
	case utils.Right:
		r.Direction = utils.Down
	}
}

func (r *Reindeer) Rotate180() {
	switch r.Direction {
	case utils.Up:
		r.Direction = utils.Down
	case utils.Down:
		r.Direction = utils.Up
	case utils.Left:
		r.Direction = utils.Right
	case utils.Right:
		r.Direction = utils.Left
	}
}

func (r *Reindeer) Move() {
	switch r.Direction {
	case utils.Up:
		r.Y--
	case utils.Down:
		r.Y++
	case utils.Left:
		r.X--
	case utils.Right:
		r.X++
	}
}

func (r *Reindeer) Undo() {
	switch r.Direction {
	case utils.Up:
		r.Y++
	case utils.Down:
		r.Y--
	case utils.Left:
		r.X++
	case utils.Right:
		r.X--
	}
}

func (r *Reindeer) GetFrontCoordinate() Coordinate {
	r.Move()
	c := r.Coordinate
	r.Undo()
	return c
}

func (r *Reindeer) GetLeftCoordinate() Coordinate {
	r.RotateLeft()
	r.Move()
	c := r.Coordinate
	r.Undo()
	r.RotateRight()
	return c
}

func (r *Reindeer) GetRightCoordinate() Coordinate {
	r.RotateRight()
	r.Move()
	c := r.Coordinate
	r.Undo()
	r.RotateLeft()
	return c
}

func (r *Reindeer) GetBackCoordinate() Coordinate {
	r.Rotate180()
	r.Move()
	c := r.Coordinate
	r.Undo()
	r.Rotate180()
	return c
}

type Maze struct {
	Map      [][]string
	Reindeer *Reindeer
	Exit     Coordinate
	Score    int
}

func NewMaze(input string) (*Maze, Reindeer) {
	m := Maze{
		Map: [][]string{},
	}
	r := Reindeer{}
	for y, line := range strings.Split(input, "\n") {
		row := []string{}
		for x, char := range strings.Split(line, "") {
			coordinate := Coordinate{x, y}
			if char == "S" {
				r.Coordinate = coordinate
				r.Direction = utils.Right
				row = append(row, ".")
			} else if char == "E" {
				m.Exit = Coordinate{x, y}
				row = append(row, ".")
			} else {
				row = append(row, char)
			}
		}
		m.Map = append(m.Map, row)
	}

	m.Reindeer = &r

	return &m, r
}

func (m *Maze) GetNeighbourCoordinates(r *Movement) (Coordinate, Coordinate, Coordinate, Coordinate) {
	switch r.Direction {
	// front, left, right, back
	case utils.Up:
		return Coordinate{r.X, r.Y - 1}, Coordinate{r.X - 1, r.Y}, Coordinate{r.X + 1, r.Y}, Coordinate{r.X, r.Y + 1}
	case utils.Down:
		return Coordinate{r.X, r.Y + 1}, Coordinate{r.X + 1, r.Y}, Coordinate{r.X - 1, r.Y}, Coordinate{r.X, r.Y - 1}
	case utils.Left:
		return Coordinate{r.X - 1, r.Y}, Coordinate{r.X, r.Y + 1}, Coordinate{r.X, r.Y - 1}, Coordinate{r.X + 1, r.Y}
	default:
		return Coordinate{r.X + 1, r.Y}, Coordinate{r.X, r.Y - 1}, Coordinate{r.X, r.Y + 1}, Coordinate{r.X - 1, r.Y}
	}
}

type MovementCost struct {
	Coordinate *Coordinate
	From       *Movement
	Cost       int
	Done       bool
}

func (m *Maze) GetOptimalRoute(appendix *map[Coordinate]*MovementCost, queue *[]*MovementCost, current *Movement, score int) {
	front, left, right, _ := m.GetNeighbourCoordinates(current)

	if m.Map[front.Y][front.X] == "." {
		cost := score + 1
		if movementCost, ok := (*appendix)[front]; ok {
			if cost < movementCost.Cost {
				movementCost.Cost = cost
				movementCost.From = current
			}
		} else {
			(*appendix)[front] = &MovementCost{Coordinate: &front, From: current, Cost: cost}
			*queue = append(*queue, (*appendix)[front])
		}
	}
	if m.Map[left.Y][left.X] == "." {
		cost := score + 1001
		direction := current.Direction
		direction.RotateLeft()
		from := Movement{Coordinate: left, Direction: direction}
		if movementCost, ok := (*appendix)[left]; ok {
			if cost < movementCost.Cost {
				movementCost.Cost = cost
				movementCost.From = &from
			}
		} else {
			(*appendix)[left] = &MovementCost{Coordinate: &left, From: &from, Cost: cost}
			*queue = append(*queue, (*appendix)[left])
		}
	}
	if m.Map[right.Y][right.X] == "." {
		cost := score + 1001
		direction := current.Direction
		direction.RotateRight()
		from := Movement{Coordinate: left, Direction: direction}
		if movementCost, ok := (*appendix)[right]; ok {
			if cost < movementCost.Cost {
				movementCost.Cost = cost
				movementCost.From = &from
			}
		} else {
			(*appendix)[right] = &MovementCost{Coordinate: &right, From: &from, Cost: cost}
			*queue = append(*queue, (*appendix)[right])
		}
	}

	(*appendix)[current.Coordinate].Done = true

	var next *MovementCost
	minimum := math.MaxInt
	for _, movementCost := range *queue {
		if !movementCost.Done && movementCost.Cost < minimum {
			next = movementCost
			minimum = movementCost.Cost
		}
	}

	if next != nil {
		m.GetOptimalRoute(appendix, queue, &Movement{Coordinate: *next.Coordinate, Direction: next.From.Direction}, next.Cost)
	}
}

func Part1() int {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	m, r := NewMaze(input)

	movement := Movement{Coordinate: r.Coordinate, Direction: r.Direction, Score: 0}
	movementCost := MovementCost{Coordinate: &r.Coordinate, From: &movement, Cost: 0}

	appendix := map[Coordinate]*MovementCost{}
	queue := []*MovementCost{}
	queue = append(queue, &movementCost)
	appendix[r.Coordinate] = &movementCost

	m.GetOptimalRoute(&appendix, &queue, &movement, 0)

	log.Print(appendix[m.Exit].Cost)

	log.Println("score", appendix[m.Exit].Cost)

	return appendix[m.Exit].Cost
}
