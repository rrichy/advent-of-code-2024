package main

import (
	"log"
	"strings"
	"time"
)

func Part2(input string, byte, size int) string {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	for {
		m := NewMemory(input, byte, size)
		m.GetOptimal(&MovementCost{Coordinate{0, 0}, nil, 0, false})

		if _, ok := m.appendix[Coordinate{size, size}]; !ok {
			break
		}

		byte++
	}

	for i, line := range strings.Split(input, "\n") {
		if i == byte-1 {
			return line
		}
	}

	return ""
}
