package day10

import (
	"log"
	"strconv"
	"strings"

	"github.com/rrichy/advent-of-code-2024/utils"
)

type Coordinate struct {
	X int
	Y int
}

type TrailHead struct {
	Coordinate
	Score int
}

type Topography struct {
	Map        [][]rune
	TrailHeads []*TrailHead
	MaxX       int
	MaxY       int
}

func NewTopography(input string) Topography {
	topography := [][]rune{}
	trailHeads := []*TrailHead{}
	for y, line := range strings.Split(input, "\n") {
		runes := []rune{}
		for x, char := range strings.Split(line, "") {
			c, _ := strconv.Atoi(char)
			runes = append(runes, rune(c))

			if c == 0 {
				trailHeads = append(trailHeads, &TrailHead{Coordinate: Coordinate{X: x, Y: y}})
			}
		}

		topography = append(topography, runes)
	}

	return Topography{Map: topography, TrailHeads: trailHeads, MaxX: len(topography[0]) - 1, MaxY: len(topography) - 1}
}

func (t *Topography) RateTrailHeadsPart1() {
	for _, trailHead := range t.TrailHeads {
		trailHead.Score = t.TraversePart1(trailHead.Coordinate, &map[Coordinate]bool{})
	}
}

func (t *Topography) GetTrailHeadsTotalScore() int {
	totalScore := 0
	for _, trailHead := range t.TrailHeads {
		totalScore += trailHead.Score
	}

	return totalScore
}

func (t *Topography) IsOutOfBounds(c Coordinate) bool {
	return c.X < 0 || c.X > t.MaxX || c.Y < 0 || c.Y > t.MaxY
}

func (t *Topography) IsTraversable(c1, c2 Coordinate) bool {
	currentElevation := t.Map[c1.Y][c1.X]
	return t.Map[c2.Y][c2.X]-currentElevation == 1
}

func (t *Topography) TraversePart1(c Coordinate, m *map[Coordinate]bool) int {
	if t.Map[c.Y][c.X] == 9 {
		if (*m)[c] {
			return 0
		}
		(*m)[c] = true

		return 1
	}

	r, l, u, d := 0, 0, 0, 0

	// Go right
	right := Coordinate{X: c.X + 1, Y: c.Y}
	if !t.IsOutOfBounds(right) && t.IsTraversable(c, right) {
		r = t.TraversePart1(right, m)
	}

	// Go left
	left := Coordinate{X: c.X - 1, Y: c.Y}
	if !t.IsOutOfBounds(left) && t.IsTraversable(c, left) {
		l = t.TraversePart1(left, m)
	}

	// Go up
	up := Coordinate{X: c.X, Y: c.Y - 1}
	if !t.IsOutOfBounds(up) && t.IsTraversable(c, up) {
		u = t.TraversePart1(up, m)
	}

	// Go down
	down := Coordinate{X: c.X, Y: c.Y + 1}
	if !t.IsOutOfBounds(down) && t.IsTraversable(c, down) {
		d = t.TraversePart1(down, m)
	}

	return r + l + u + d
}

func Part1() int {
	input := utils.ReadInput("day_10/input")

	topography := NewTopography(input)
	topography.RateTrailHeadsPart1()

	total := topography.GetTrailHeadsTotalScore()

	log.Print(total)

	return total
}
