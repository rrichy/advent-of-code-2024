package main

import (
	"image/color"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/rrichy/advent-of-code-2024/utils"
)

var re = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

type Position utils.Coordinate
type Velocity utils.Coordinate

type Robot struct {
	OriginalPosition Position
	Position         Position
	Velocity         Velocity
	Color            color.RGBA
}

func Part1() int {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	robots := []Robot{}
	for _, line := range strings.Split(input, "\n") {
		m := re.FindStringSubmatch(line)

		position := Position{utils.MustAtoi(m[1]), utils.MustAtoi(m[2])}

		robots = append(robots, Robot{
			OriginalPosition: position,
			Position:         position,
			Velocity:         Velocity{utils.MustAtoi(m[3]), utils.MustAtoi(m[4])},
		})
	}

	w, h := 101, 103
	seconds := 100
	quadrants := make([]int, 4)

	for i, robot := range robots {
		x := (robot.OriginalPosition.X + seconds*robot.Velocity.X) % w
		y := (robot.OriginalPosition.Y + seconds*robot.Velocity.Y) % h

		if x >= 0 {
			robots[i].Position.X = x
		} else {
			robots[i].Position.X = x + w
		}

		if y >= 0 {
			robots[i].Position.Y = y
		} else {
			robots[i].Position.Y = y + h
		}

		if robots[i].Position.X < w/2 && robots[i].Position.Y < h/2 {
			quadrants[0]++
		} else if robots[i].Position.X > w/2 && robots[i].Position.Y < h/2 {
			quadrants[1]++
		} else if robots[i].Position.X < w/2 && robots[i].Position.Y > h/2 {
			quadrants[2]++
		} else if robots[i].Position.X > w/2 && robots[i].Position.Y > h/2 {
			quadrants[3]++
		}
	}

	log.Print(quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3])

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}
