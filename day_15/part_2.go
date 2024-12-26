package main

import (
	"log"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewObject(x, y int, movable bool, char string) Object {
	return Object{Coordinate{x, y}, movable, char, nil, []*ebiten.Image{}, nil, false, false, false}
}

func NewWarehouseTwice(s string) Warehouse {
	w := Warehouse{
		Map: [][]*Object{},
	}
	for y, line := range strings.Split(s, "\n") {
		l := []*Object{}
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				o1 := NewObject(x*2, y, false, char)
				o2 := NewObject(x*2+1, y, false, char)
				l = append(l, &o1, &o2)
			} else if char == "." {
				l = append(l, nil, nil)
			} else if char == "O" {
				o1 := NewObject(x*2, y, true, "[")
				o2 := NewObject(x*2+1, y, true, "]")
				o1.Couple = &o2
				o2.Couple = &o1
				l = append(l, &o1, &o2)
			} else {
				o := NewObject(x*2, y, true, char)
				l = append(l, &o, nil)
				w.Robot = &o
			}
		}

		w.Map = append(w.Map, l)
	}
	w.Width = len(w.Map[0])
	w.Height = len(w.Map)

	return w
}

func (w *Warehouse) IsBoxMoveable(o1 *Object, dc *Coordinate) bool {
	if o1 == nil {
		return true
	}

	if !o1.Movable {
		return false
	}

	o2 := o1.Couple

	c1 := Coordinate{o1.X + dc.X, o1.Y + dc.Y}
	c2 := Coordinate{o2.X + dc.X, o2.Y + dc.Y}

	if w.Map[c1.Y][c1.X] == nil && w.Map[c2.Y][c2.X] == nil {
		return true
	}

	return w.IsBoxMoveable(w.Map[c1.Y][c1.X], dc) && w.IsBoxMoveable(w.Map[c2.Y][c2.X], dc)
}

func (w *Warehouse) Move2(o *Object, dc *Coordinate) bool {
	if o == nil || !o.Movable {
		return false
	}

	c := Coordinate{o.X + dc.X, o.Y + dc.Y}

	if o.Couple == nil || dc.Y == 0 {
		if w.Map[c.Y][c.X] == nil || w.Move2(w.Map[c.Y][c.X], dc) {
			w.Map[o.Y][o.X] = nil
			w.Map[c.Y][c.X] = o
			o.X = c.X
			o.Y = c.Y
			return true
		}

		return false
	}

	o2 := o.Couple
	c2 := Coordinate{o.Couple.X + dc.X, o.Couple.Y + dc.Y}

	if w.Map[c.Y][c.X] == nil && w.Map[c2.Y][c2.X] == nil {
		w.Map[o.Y][o.X] = nil
		w.Map[o2.Y][o2.X] = nil
		w.Map[c.Y][c.X] = o
		w.Map[c2.Y][c2.X] = o2
		o.X = c.X
		o.Y = c.Y
		o2.X = c2.X
		o2.Y = c2.Y
		return true
	}

	if w.IsBoxMoveable(o, dc) {
		w.Move2(w.Map[c.Y][c.X], dc)
		w.Move2(w.Map[c2.Y][c2.X], dc)

		w.Map[o.Y][o.X] = nil
		w.Map[o2.Y][o2.X] = nil
		w.Map[c.Y][c.X] = o
		w.Map[c2.Y][c2.X] = o2
		o.X = c.X
		o.Y = c.Y
		o2.X = c2.X
		o2.Y = c2.Y
		return true
	}

	return false
}

func Part2() int {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	s := strings.Split(input, "\n\n")
	w := NewWarehouseTwice(s[0])

	for _, line := range strings.Split(s[1], "\n") {
		for _, dir := range strings.Split(line, "") {
			dc := GetDisplacement(dir)
			w.Move2(w.Robot, &dc)
		}
	}

	w.PrintMap()

	sum := 0
	for y, line := range w.Map {
		s := ""
		for x, o := range line {
			if o != nil {
				s += o.Char
			} else {
				s += " "
			}

			if o != nil && o.Char == "[" {
				sum += 100*y + x
			}
		}
	}

	log.Print(sum)

	return sum
}
