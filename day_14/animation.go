package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand/v2"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/rrichy/advent-of-code-2024/utils"
)

func (h *Headquarters) Draw(screen *ebiten.Image) {
	for _, robot := range h.Robots {
		x := float32(robot.Position.X)
		y := float32(robot.Position.Y)
		vector.DrawFilledRect(screen, x*scale, y*scale, scale, scale, robot.Color, true)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Seconds: %d", h.Seconds))
}

func (h *Headquarters) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return outsideWidth, outsideHeight
}

func Animate() {
	h := Headquarters{
		Robots:  []*Robot{},
		RowMap:  make([][]*Robot, height),
		Seconds: 6150,
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

	ebiten.SetWindowSize(width*int(scale), height*int(scale))
	ebiten.SetWindowTitle("Advent of Code 2024 - Day 14 - Part 2")

	if err := ebiten.RunGame(&h); err != nil {
		log.Fatal(err)
	}
}
