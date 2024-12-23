package main

import (
	"log"
	"strings"
	"time"

	"github.com/rrichy/advent-of-code-2024/utils"
)

func NewModMachine(s string) Machine {
	ss := strings.Split(s, "\n")

	buttonA := re1.FindAllString(ss[0], -1)
	buttonB := re1.FindAllString(ss[1], -1)
	prize := re2.FindAllString(ss[2], -1)

	BA := Disposition{utils.MustAtoi(buttonA[0]), utils.MustAtoi(buttonA[1])}
	BB := Disposition{utils.MustAtoi(buttonB[0]), utils.MustAtoi(buttonB[1])}
	P := Coordinate{utils.MustAtoi(prize[0]) + 10000000000000, utils.MustAtoi(prize[1]) + 10000000000000}

	a_count := (float64(P.X)*float64(BB.Y) - float64(P.Y)*float64(BB.X)) / (float64(BA.X)*float64(BB.Y) - float64(BA.Y)*float64(BB.X))
	b_count := (float64(P.X)*float64(BA.Y) - float64(P.Y)*float64(BA.X)) / (float64(BA.Y)*float64(BB.X) - float64(BA.X)*float64(BB.Y))

	m := Machine{
		ButtonA: BA,
		ButtonB: BB,
		Prize:   P,
	}

	if int(a_count)*BA.X+int(b_count)*BB.X == P.X && int(a_count)*BA.Y+int(b_count)*BB.Y == P.Y {
		m.ACount = int(a_count)
		m.BCount = int(b_count)
	}

	return m
}

func Part2() int {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	sum := 0
	for _, raw := range strings.Split(input, "\n\n") {
		machine := NewModMachine(raw)

		sum += machine.ACount*3 + machine.BCount
	}

	log.Print(sum)

	return sum
}
