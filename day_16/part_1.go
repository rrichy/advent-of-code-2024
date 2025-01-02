package main

import (
	"log"
	"math"
	"strings"
	"time"

	"github.com/rrichy/advent-of-code-2024/utils"
)

type Coordinate utils.Coordinate

type Movement struct {
	Coordinate
	Direction utils.Direction
}

type Reindeer Movement

type Maze struct {
	Map      [][]string
	Reindeer *Reindeer
	Exit     Coordinate
	Score    int
	Appendix map[Coordinate]*MovementCost
}

func NewMaze(input string) (*Maze, Reindeer) {
	m := Maze{
		Map:      [][]string{},
		Appendix: map[Coordinate]*MovementCost{},
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
	Movement
	From *MovementCost
	Cost int
	Done bool
}

func (m *Maze) GetOptimalRoute(queue *[]*MovementCost, current *MovementCost) {
	front, left, right, _ := m.GetNeighbourCoordinates(&current.Movement)
	score := current.Cost

	if m.Map[front.Y][front.X] == "." {
		cost := score + 1
		if mc, ok := m.Appendix[front]; ok {
			if cost <= mc.Cost {
				mc.Cost = cost
				mc.From = current
			}
		} else {
			destination := Movement{Coordinate: front, Direction: current.Movement.Direction}
			m.Appendix[front] = &MovementCost{Movement: destination, From: current, Cost: cost}
			*queue = append(*queue, m.Appendix[front])
		}
	}
	if m.Map[left.Y][left.X] == "." {
		cost := score + 1001
		direction := current.Movement.Direction
		direction.RotateLeft()
		from := MovementCost{Movement: Movement{Coordinate: current.Coordinate, Direction: direction}, From: current.From, Cost: current.Cost}
		if mc, ok := m.Appendix[left]; ok {
			if cost <= mc.Cost {
				mc.Cost = cost
				mc.From = &from
			}
		} else {
			destination := Movement{Coordinate: left, Direction: direction}
			m.Appendix[left] = &MovementCost{Movement: destination, From: &from, Cost: cost}
			*queue = append(*queue, m.Appendix[left])
		}
	}
	if m.Map[right.Y][right.X] == "." {
		cost := score + 1001
		direction := current.Movement.Direction
		direction.RotateRight()
		from := MovementCost{Movement: Movement{Coordinate: current.Coordinate, Direction: direction}, From: current.From, Cost: current.Cost}
		if movementCost, ok := m.Appendix[right]; ok {
			if cost <= movementCost.Cost {
				movementCost.Cost = cost
				movementCost.From = &from
			}
		} else {
			destination := Movement{Coordinate: right, Direction: direction}
			m.Appendix[right] = &MovementCost{Movement: destination, From: &from, Cost: cost}
			*queue = append(*queue, m.Appendix[right])
		}
	}

	current.Done = true

	var next *MovementCost
	minimum := math.MaxInt
	for _, movementCost := range *queue {
		if !movementCost.Done && movementCost.Cost < minimum {
			next = movementCost
			minimum = movementCost.Cost
		}
	}

	if next != nil {
		m.GetOptimalRoute(queue, next)
	}
}

func Part1(input string) int {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	m, r := NewMaze(input)

	movement := Movement(r)
	movementCost := MovementCost{Movement: movement, Cost: 0}
	m.Appendix[r.Coordinate] = &movementCost

	queue := []*MovementCost{}

	m.GetOptimalRoute(&queue, &movementCost)

	return m.Appendix[m.Exit].Cost
}
