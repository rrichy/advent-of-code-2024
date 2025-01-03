package main

import (
	"log"
	"time"

	"github.com/rrichy/advent-of-code-2024/utils"
)

func (m *Maze) TraverseToOrigin(targetCoordinate, currentCoordinate Coordinate, v *map[Coordinate]bool) {
	(*v)[currentCoordinate] = true

	current := m.Appendix[currentCoordinate]
	target := m.Appendix[targetCoordinate]

	if currentCoordinate == m.Reindeer.Coordinate {
		return
	}

	f, l, r, b := m.GetNeighbourCoordinates(&current.Movement)
	left, right, back, front := m.Appendix[l], m.Appendix[r], m.Appendix[b], m.Appendix[f]

	if current.Movement.Coordinate.X == 11 && current.Movement.Coordinate.Y == 123 {
		log.Println("current", current.Movement.Coordinate, current.Cost)
	}
	if current.Movement.Coordinate.X == 9 && current.Movement.Coordinate.Y == 127 {
		log.Println("current", current.Movement.Coordinate, current.Cost)
	}
	if current.Movement.Coordinate.X == 9 && current.Movement.Coordinate.Y == 129 {
		log.Println("current", current.Movement.Coordinate, current.Cost)
	}
	if current.Movement.Coordinate.X == 3 && current.Movement.Coordinate.Y == 9 {
		log.Println("current", current.Movement.Coordinate, current.Cost)
	}
	if current.Movement.Coordinate.X == m.Exit.X && current.Movement.Coordinate.Y == m.Exit.Y {
		log.Println("current", current.Movement.Coordinate, current.Cost)
	}

	if m.ShouldTraverse(target, current, left, v) {
		m.TraverseToOrigin(currentCoordinate, left.Coordinate, v)
	}

	if m.ShouldTraverse(target, current, right, v) {
		m.TraverseToOrigin(currentCoordinate, right.Coordinate, v)
	}

	if m.ShouldTraverse(target, current, back, v) {
		m.TraverseToOrigin(currentCoordinate, back.Coordinate, v)
	}

	if m.ShouldTraverse(target, current, front, v) {
		m.TraverseToOrigin(currentCoordinate, front.Coordinate, v)
	}
}

func (m *Movement) FaceDirectionCost(direction utils.Direction) int {
	if m.Direction == direction {
		return 0
	}

	if m.Direction == utils.Up && direction == utils.Down ||
		m.Direction == utils.Down && direction == utils.Up ||
		m.Direction == utils.Left && direction == utils.Right ||
		m.Direction == utils.Right && direction == utils.Left {
		return 2000
	}

	return 1000
}

func (m *Maze) ShouldTraverse(target, current, possibleFrom *MovementCost, v *map[Coordinate]bool) bool {
	if possibleFrom == nil ||
		(*v)[possibleFrom.Coordinate] ||
		(possibleFrom.From != nil && possibleFrom.From.Coordinate == current.Coordinate) ||
		target.Coordinate == possibleFrom.Coordinate {
		return false
	}

	if current.From.Coordinate == possibleFrom.Coordinate {
		return true
	}

	cost := possibleFrom.FaceDirectionCost(target.Direction) + possibleFrom.Cost + utils.AbsInt(possibleFrom.X-target.X) + utils.AbsInt(possibleFrom.Y-target.Y)
	if cost == target.Cost || (possibleFrom.FaceDirectionCost(target.Direction) == 2000 && (utils.AbsInt(possibleFrom.X-target.X) == 2 || utils.AbsInt(possibleFrom.Y-target.Y) == 2)) {
		return true
	}

	return false
}

// TODO: Redo. Not a very satisfying solution.
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

	visited := map[Coordinate]bool{}
	visited[m.Exit] = true
	m.TraverseToOrigin(m.Exit, m.Appendix[m.Exit].From.Coordinate, &visited)

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

	return len(visited)
}
