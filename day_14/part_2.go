package main

import (
	"image/color"
	"log"
	"math/rand/v2"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rrichy/advent-of-code-2024/utils"
)

type Headquarters struct {
	Robots  []*Robot
	RowMap  [][]*Robot
	Seconds int
	Halt    bool
}

var (
	width  = 101
	height = 103
	scale  = float32(4)
)

func (h *Headquarters) ClearMap() {
	for i := range h.RowMap {
		h.RowMap[i] = make([]*Robot, width)
	}
}

func (h *Headquarters) PotentialEasterEggAppeared() bool {
	minLineLength := 15
	for _, row := range h.RowMap {
		if len(row) > 15 {
			length := 0
			for _, robot := range row {
				if robot != nil {
					length++
				} else {
					if length > minLineLength {
						return true
					}
					length = 0
				}
			}
		}
	}

	return false
}

func (h *Headquarters) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		h.Halt = false
	}

	if h.PotentialEasterEggAppeared() {
		h.Halt = true
	}

	if !h.Halt {
		h.ClearMap()
		h.Seconds++
		for i, robot := range h.Robots {
			x := (robot.OriginalPosition.X + h.Seconds*robot.Velocity.X) % width
			y := (robot.OriginalPosition.Y + h.Seconds*robot.Velocity.Y) % height

			if x >= 0 {
				h.Robots[i].Position.X = x
			} else {
				h.Robots[i].Position.X = x + width
			}

			if y >= 0 {
				h.Robots[i].Position.Y = y
			} else {
				h.Robots[i].Position.Y = y + height
			}

			h.RowMap[robot.Position.Y][robot.Position.X] = robot
		}
	}

	return nil
}

func Part2() int {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	h := Headquarters{
		Robots: []*Robot{},
		RowMap: make([][]*Robot, height),
		// Seconds: 6510,
	}

	h.ClearMap()

	for _, line := range strings.Split(input, "\n") {
		m := re.FindStringSubmatch(line)

		position := Position{utils.MustAtoi(m[1]), utils.MustAtoi(m[2])}
		robot := Robot{
			OriginalPosition: position,
			Position:         position,
			Velocity:         Velocity{utils.MustAtoi(m[3]), utils.MustAtoi(m[4])},
			Color:            color.RGBA{uint8(rand.Uint()), uint8(rand.Uint()), uint8(rand.Uint()), 0xff},
		}

		h.RowMap[robot.Position.Y][robot.Position.X] = &robot

		h.Robots = append(h.Robots, &robot)
	}

	for {
		if h.Halt {
			break
		}

		_ = h.Update()
	}

	log.Print(h.Seconds)

	return h.Seconds
}
