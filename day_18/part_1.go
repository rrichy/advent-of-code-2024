package main

import (
	"log"
	"math"
	"strings"
	"time"

	"github.com/rrichy/advent-of-code-2024/utils"
)

type Coordinate utils.Coordinate

type MovementCost struct {
	Coordinate
	from *MovementCost
	cost int
	done bool
}

type Memory struct {
	corrupted map[Coordinate]bool
	size      int
	queue     []*MovementCost
	appendix  map[Coordinate]*MovementCost
}

func NewMemory(input string, bytes, size int) Memory {
	corrupted := map[Coordinate]bool{}

	for i, line := range strings.Split(input, "\n") {
		if i >= bytes {
			break
		}

		splits := strings.Split(line, ",")
		x := utils.MustAtoi(splits[0])
		y := utils.MustAtoi(splits[1])
		c := Coordinate{x, y}

		corrupted[c] = true
	}

	return Memory{corrupted, size, []*MovementCost{}, map[Coordinate]*MovementCost{}}
}

func (m *Memory) IsTraversable(c *MovementCost) bool {
	if m.corrupted[c.Coordinate] {
		return false
	}

	if c.X < 0 || c.Y < 0 || c.X > m.size || c.Y > m.size {
		return false
	}

	if c.done {
		return false
	}

	return true
}

func (m *Memory) SetCost(mcUpdate, c *MovementCost) {
	if mcCurrent, ok := m.appendix[mcUpdate.Coordinate]; ok {
		if mcUpdate.cost <= mcCurrent.cost {
			mcCurrent.cost = mcUpdate.cost
			mcCurrent.from = c
		}
	} else {
		m.appendix[mcUpdate.Coordinate] = mcUpdate
		m.queue = append(m.queue, m.appendix[mcUpdate.Coordinate])
	}
}

func (m *Memory) GetOptimal(c *MovementCost) {
	cost := c.cost + 1
	left := MovementCost{Coordinate{c.X - 1, c.Y}, c, cost, false}
	right := MovementCost{Coordinate{c.X + 1, c.Y}, c, cost, false}
	up := MovementCost{Coordinate{c.X, c.Y - 1}, c, cost, false}
	down := MovementCost{Coordinate{c.X, c.Y + 1}, c, cost, false}

	if m.IsTraversable(&left) {
		m.SetCost(&left, c)
	}

	if m.IsTraversable(&right) {
		m.SetCost(&right, c)
	}

	if m.IsTraversable(&up) {
		m.SetCost(&up, c)
	}

	if m.IsTraversable(&down) {
		m.SetCost(&down, c)
	}

	c.done = true

	var next *MovementCost
	minimum := math.MaxInt
	for _, movementCost := range m.queue {
		if !movementCost.done && movementCost.cost < minimum {
			next = movementCost
			minimum = movementCost.cost
		}
	}

	if next != nil {
		m.GetOptimal(next)
	}
}

func (m *Memory) DrawPath() {
	path := map[Coordinate]bool{}
	mc := m.appendix[Coordinate{m.size, m.size}]
	for {
		path[mc.Coordinate] = true
		if mc.from == nil {
			break
		}
		mc = mc.from
	}

	for y := 0; y <= m.size; y++ {
		line := ""
		for x := 0; x <= m.size; x++ {
			if m.corrupted[Coordinate{x, y}] {
				line += "#"
			} else if path[Coordinate{x, y}] {
				line += "O"
			} else {
				line += "."
			}
		}
		log.Println(line)
	}
}

func Part1(input string, bytes, size int) int {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	memory := NewMemory(input, bytes, size)
	memory.GetOptimal(&MovementCost{Coordinate{0, 0}, nil, 0, false})

	cost := memory.appendix[Coordinate{size, size}].cost
	memory.DrawPath()

	return cost
}
