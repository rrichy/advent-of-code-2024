package main

import (
	"log"
	"time"
)

func (m *Maze) TraverseToOrigin(current *MovementCost, v *map[Coordinate]bool) {
	(*v)[current.Movement.Coordinate] = true

	if current.Movement.Coordinate == m.Reindeer.Coordinate {
		return
	}

	f, l, r, b := m.GetNeighbourCoordinates(&current.Movement)
	left, right, back, front := m.Appendix[l], m.Appendix[r], m.Appendix[b], m.Appendix[f]
	from := current.From

	m.TraverseToOrigin(from, v)

	if left != nil && (left.Coordinate != from.Coordinate && left.Cost == from.Cost) {
		m.TraverseToOrigin(left, v)
	}

	if right != nil && (right.Coordinate != from.Coordinate && right.Cost == from.Cost) {
		m.TraverseToOrigin(right, v)
	}

	if back != nil && (back.Coordinate != from.Coordinate && back.Cost == from.Cost) {
		m.TraverseToOrigin(back, v)
	}

	if front != nil && (front.Coordinate != from.Coordinate && front.Cost == from.Cost) {
		m.TraverseToOrigin(front, v)
	}
	// if left != nil && (left.Cost-from.Cost == 1000 || (left.Coordinate != from.Coordinate && left.Cost == from.Cost)) {
	// 	m.TraverseToOrigin(left, v)
	// }

	// if right != nil && (right.Cost-from.Cost == 1000 || (right.Coordinate != from.Coordinate && right.Cost == from.Cost)) {
	// 	m.TraverseToOrigin(right, v)
	// }

	// if back != nil && (back.Cost-from.Cost == 1000 || (back.Coordinate != from.Coordinate && back.Cost == from.Cost)) {
	// 	m.TraverseToOrigin(back, v)
	// }

	// if front != nil && (front.Cost-from.Cost == 1000 || (front.Coordinate != from.Coordinate && front.Cost == from.Cost)) {
	// 	m.TraverseToOrigin(front, v)
	// }
}

func Part2(input string) int {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	m, r := NewMaze(input)

	movement := Movement(r)
	movementCost := MovementCost{Movement: movement, Cost: 0}
	m.Appendix[r.Coordinate] = &movementCost

	queue := []*MovementCost{}

	m.GetOptimalRoute(&queue, &movementCost)

	optimalRoute := m.Appendix[m.Exit]
	visited := map[Coordinate]bool{}
	m.TraverseToOrigin(optimalRoute, &visited)

	for y := range m.Map {
		line := ""
		for x := range m.Map[y] {
			if visited[Coordinate{X: x, Y: y}] {
				line += "O"
			} else {
				line += m.Map[y][x]
			}
		}
		log.Print(line)
	}

	log.Print(len(visited))

	log.Println("score", m.Appendix[m.Exit].Cost)

	// 520 too higg

	return len(visited)
}
