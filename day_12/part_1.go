package main

import (
	_ "embed"
	"log"
	"time"

	"github.com/rrichy/advent-of-code-2024/utils"
)

//go:embed input
var input string

type Coordinate struct {
	X int
	Y int
}

type Region struct {
	Area      int
	Perimeter int
	Sides     int
}

type Garden struct {
	Map     [][]string
	Plants  map[string][]*Region
	Visited map[Coordinate]*Region
	Width   int
	Height  int
}

func NewGarden(input string) Garden {
	i := utils.New2DStringMatrix(input)
	return Garden{
		Map:     i,
		Plants:  map[string][]*Region{},
		Visited: map[Coordinate]*Region{},
		Width:   len(i[0]),
		Height:  len(i),
	}
}

func (g *Garden) GetLabel(c Coordinate) string {
	return g.Map[c.Y][c.X]
}

func (g *Garden) IsOutOfBounds(c Coordinate) bool {
	return c.X < 0 || c.X >= g.Width || c.Y < 0 || c.Y >= g.Height
}

func (g *Garden) CountPerimeter(c Coordinate) (int, int, int, int, int) {
	t, r, b, l := 0, 0, 0, 0
	label := g.Map[c.Y][c.X]
	count := 0

	if c.Y == 0 || (c.Y > 0 && g.Map[c.Y-1][c.X] != label) {
		count++
		t = 1
	}

	if c.Y == g.Height-1 || (c.Y < g.Height-1 && g.Map[c.Y+1][c.X] != label) {
		count++
		b = 1
	}

	if c.X == 0 || (c.X > 0 && g.Map[c.Y][c.X-1] != label) {
		count++
		l = 1
	}

	if c.X == g.Width-1 || (c.X < g.Width-1 && g.Map[c.Y][c.X+1] != label) {
		count++
		r = 1
	}

	return count, t, r, b, l
}

func (g *Garden) CountInnerCorner(c Coordinate, deltaX, deltaY int) int {
	label := g.GetLabel(c)

	p1 := Coordinate{c.X + deltaX, c.Y}
	p2 := Coordinate{c.X, c.Y + deltaY}
	p3 := Coordinate{c.X + deltaX, c.Y + deltaY}

	if g.IsOutOfBounds(p1) || g.IsOutOfBounds(p2) || g.IsOutOfBounds(p3) {
		return 0
	}

	if label == g.GetLabel(p1) && label == g.GetLabel(p2) && label != g.GetLabel(p3) {
		return 1
	}

	return 0
}

func (g *Garden) Spread(region *Region, label string, c Coordinate) {
	if _, ok := g.Visited[c]; ok || g.GetLabel(c) != label {
		return
	}

	perimeter, t, r, b, l := g.CountPerimeter(c)

	(*region).Area += 1
	(*region).Perimeter += perimeter

	if perimeter == 4 {
		(*region).Sides += 4
	} else if perimeter == 3 {
		(*region).Sides += 2
	} else {
		if perimeter == 2 && ((t == 1 && r == 1) || (t == 1 && l == 1) || (b == 1 && r == 1) || (b == 1 && l == 1)) {
			(*region).Sides += 1
		}
		(*region).Sides += g.CountInnerCorner(c, 1, 1)
		(*region).Sides += g.CountInnerCorner(c, 1, -1)
		(*region).Sides += g.CountInnerCorner(c, -1, 1)
		(*region).Sides += g.CountInnerCorner(c, -1, -1)
	}

	g.Visited[c] = region
	if c.X > 0 {
		g.Spread(region, label, Coordinate{c.X - 1, c.Y})
	}
	if c.X < g.Width-1 {
		g.Spread(region, label, Coordinate{c.X + 1, c.Y})
	}
	if c.Y > 0 {
		g.Spread(region, label, Coordinate{c.X, c.Y - 1})
	}
	if c.Y < g.Height-1 {
		g.Spread(region, label, Coordinate{c.X, c.Y + 1})
	}
}

func (g *Garden) FindRegions() {
	for y, horizontal := range g.Map {
		for x := range horizontal {
			c := Coordinate{x, y}

			if _, ok := g.Visited[c]; ok {
				continue
			}

			region := Region{}
			label := g.GetLabel(c)
			g.Spread(&region, label, c)

			g.Plants[label] = append(g.Plants[label], &region)
		}
	}
}

func Part1() int {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	garden := NewGarden(input)
	garden.FindRegions()

	sum := 0
	for _, regions := range garden.Plants {
		for _, region := range regions {
			sum += region.Area * region.Perimeter
		}
	}

	log.Print(sum)

	return sum
}
