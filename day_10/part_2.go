package day10

import (
	"log"

	"github.com/rrichy/advent-of-code-2024/utils"
)

func (t *Topography) RateTrailHeadsPart2() {
	for _, trailHead := range t.TrailHeads {
		trailHead.Score = t.TraversePart2(trailHead.Coordinate)
	}
}

func (t *Topography) TraversePart2(c Coordinate) int {
	if t.Tiles[c.Y][c.X].Elevation == 9 {
		return 1
	}

	r, l, u, d := 0, 0, 0, 0

	// Go right
	right := Coordinate{X: c.X + 1, Y: c.Y}
	if !t.IsOutOfBounds(right) && t.IsTraversable(c, right) {
		r = t.TraversePart2(right)
	}

	// Go left
	left := Coordinate{X: c.X - 1, Y: c.Y}
	if !t.IsOutOfBounds(left) && t.IsTraversable(c, left) {
		l = t.TraversePart2(left)
	}

	// Go up
	up := Coordinate{X: c.X, Y: c.Y - 1}
	if !t.IsOutOfBounds(up) && t.IsTraversable(c, up) {
		u = t.TraversePart2(up)
	}

	// Go down
	down := Coordinate{X: c.X, Y: c.Y + 1}
	if !t.IsOutOfBounds(down) && t.IsTraversable(c, down) {
		d = t.TraversePart2(down)
	}

	return r + l + u + d
}

func Part2() int {
	input := utils.ReadInput("day_10/input")

	topography := NewTopography(input)
	topography.RateTrailHeadsPart2()

	total := topography.GetTrailHeadsTotalScore()

	log.Print(total)

	return total
}
