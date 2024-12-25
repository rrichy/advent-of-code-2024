package main

import (
	"log"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rrichy/advent-of-code-2024/utils"
)

type Coordinate utils.Coordinate

type Object struct {
	Coordinate
	Movable            bool
	Char               string
	Couple             *Object
	Sprites            []*ebiten.Image
	PreviousCoordinate *Coordinate
	IsWater            bool
}

type Warehouse struct {
	Map           [][]*Object
	GroundSprites [][]*ebiten.Image
	Robot         *Object
	Width         int
	Height        int
	Commands      []string
}

func NewWarehouse(s string) Warehouse {
	w := Warehouse{
		Map: [][]*Object{},
	}
	for y, line := range strings.Split(s, "\n") {
		l := []*Object{}
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				l = append(l, &Object{Coordinate{X: x, Y: y}, false, char, nil, nil, nil, false})
			} else if char == "." {
				l = append(l, nil)
			} else {
				l = append(l, &Object{Coordinate{X: x, Y: y}, true, char, nil, nil, nil, false})
				if char == "@" {
					w.Robot = l[x]
				}
			}
		}

		w.Map = append(w.Map, l)
	}
	w.Width = len(w.Map[0])
	w.Height = len(w.Map)

	return w
}

func (w *Warehouse) PrintMap() {
	s := ""
	for _, line := range w.Map {
		for _, o := range line {
			if o == nil {
				s += "."
			} else {
				s += o.Char
			}
		}
		s += "\n"
	}
	log.Print("\n" + s)
	log.Print(w.Robot)
}

func (w *Warehouse) IsOutOfBounds(c Coordinate) bool {
	return c.X < 0 || c.X >= w.Width || c.Y < 0 || c.Y >= w.Height
}

func (w *Warehouse) Move(o *Object, dc *Coordinate) bool {
	c := Coordinate{o.X + dc.X, o.Y + dc.Y}
	if !o.Movable || w.IsOutOfBounds(c) {
		return false
	}

	if w.Map[c.Y][c.X] == nil || w.Move(w.Map[c.Y][c.X], dc) {
		w.Map[o.Y][o.X] = nil
		w.Map[c.Y][c.X] = o
		o.X = c.X
		o.Y = c.Y
		return true
	}

	return false
}

func GetDisplacement(dir string) Coordinate {
	if dir == "<" {
		return Coordinate{-1, 0}
	}
	if dir == ">" {
		return Coordinate{1, 0}
	}
	if dir == "^" {
		return Coordinate{0, -1}
	}
	if dir == "v" {
		return Coordinate{0, 1}
	}

	log.Fatal("Invalid direction")
	return Coordinate{0, 0}
}

func Part1() int {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	s := strings.Split(input, "\n\n")
	w := NewWarehouse(s[0])

	for _, line := range strings.Split(s[1], "\n") {
		for _, dir := range strings.Split(line, "") {
			dc := GetDisplacement(dir)
			w.Move(w.Robot, &dc)
		}
	}

	w.PrintMap()

	sum := 0
	for y, line := range w.Map {
		for x, o := range line {
			if o != nil && o.Char == "O" {
				sum += 100*y + x
			}
		}
	}

	log.Print(sum)

	return sum
}
