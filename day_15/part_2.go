package main

import (
	"log"
	"strings"
	"time"
)

func NewWarehouseTwice(s string) Warehouse {
	w := Warehouse{
		Map: [][]*Object{},
	}
	for y, line := range strings.Split(s, "\n") {
		l := []*Object{}
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				o1 := Object{Coordinate{X: x * 2, Y: y}, false, char, nil}
				o2 := Object{Coordinate{X: x*2 + 1, Y: y}, false, char, nil}
				l = append(l, &o1, &o2)
			} else if char == "." {
				l = append(l, nil, nil)
			} else if char == "O" {
				o1 := Object{Coordinate{X: x * 2, Y: y}, true, "[", nil}
				o2 := Object{Coordinate{X: x*2 + 1, Y: y}, true, "]", nil}
				o1.Couple = &o2
				o2.Couple = &o1
				l = append(l, &o1, &o2)
			} else {
				o := Object{Coordinate{X: x * 2, Y: y}, true, char, nil}
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

func (w *Warehouse) Move2(o *Object, dc *Coordinate) bool {
	if o == nil {
		return true
	}

	c := Coordinate{o.X + dc.X, o.Y + dc.Y}
	if !o.Movable || w.IsOutOfBounds(c) {
		return false
	}

	// Robot roaming
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

	if (w.Map[c.Y][c.X] == nil && w.Map[c2.Y][c2.X] == nil) || (w.Move2(w.Map[c.Y][c.X], dc) && w.Move2(w.Map[c2.Y][c2.X], dc)) {
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

	s := strings.Split(sample3, "\n\n")
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
		log.Print(s)
	}

	log.Print(sum)

	//1517281 too low
	return sum
}
