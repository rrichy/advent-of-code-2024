package main

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rrichy/advent-of-code-2024/utils"
)

func (d *DiskMap) DefragPaint(s tcell.Screen) {
	freeStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlack)
	for i := range d.Blocks {
		block := d.Blocks[len(d.Blocks)-1-i]
		for _, free := range d.Frees {
			if free.Size == 0 || free.Index > block.Index {
				continue
			}

			if free.Size >= block.Size {
				x, y := indexToCoord(block.Index)
				_, _, blockStyle, _ := s.GetContent(x, y)

				for j := 0; j < block.Size; j++ {
					d.Cells[free.Index+j] = &block.Id
					d.Cells[block.Index+j] = nil

					x1, y1 := indexToCoord(free.Index + j)
					x2, y2 := indexToCoord(block.Index + j)
					updateCell(s, x1, y1, blockStyle)
					updateCell(s, x2, y2, freeStyle)
				}

				s.Show()

				block.Index = free.Index
				free.Size -= block.Size
				free.Index += block.Size

				break
			}
		}
	}
}
func (d *DiskMap) Defrag() {
	for i := range d.Blocks {
		block := d.Blocks[len(d.Blocks)-1-i]
		for _, free := range d.Frees {
			if free.Size == 0 || free.Index > block.Index {
				continue
			}

			if free.Size >= block.Size {
				for j := 0; j < block.Size; j++ {
					d.Cells[free.Index+j] = &block.Id
					d.Cells[block.Index+j] = nil
				}

				block.Index = free.Index
				free.Size -= block.Size
				free.Index += block.Size

				break
			}
		}
	}
}

var width = 300

func indexToCoord(index int) (int, int) {
	return index % width, index / width
}

func (d *DiskMap) Paint(s tcell.Screen) {
	for _, block := range d.Blocks {
		color := generateRandomColor()
		for j := 0; j < block.Size; j++ {
			x, y := indexToCoord(block.Index + j)
			s.SetContent(x, y, ' ', nil, tcell.StyleDefault.Background(color))
		}
	}
}

func Part2() int {
	diskMap := DiskMap{}
	currentId := 0
	diskMapPointer := 0
	for i, block := range strings.Split(input, "") {
		size, _ := strconv.Atoi(block)
		if i%2 == 0 {
			diskMap.AddBlocks(currentId, size)
			currentId++
		} else {
			diskMap.Expand(size)
		}
		diskMapPointer += size
	}

	diskMap.Defrag()

	sum := 0
	for i, block := range diskMap.Cells {
		if block != nil {
			sum += i * *block
		}
	}

	log.Printf("Sum: %d\n", sum)

	return sum
}

func Part2Animate() {
	input := utils.ReadInput("day_9/input2")

	diskMap := DiskMap{}
	currentId := 0
	diskMapPointer := 0
	for i, block := range strings.Split(input, "") {
		size, _ := strconv.Atoi(block)
		if i%2 == 0 {
			diskMap.AddBlocks(currentId, size)
			currentId++
		} else {
			diskMap.Expand(size)
		}
		diskMapPointer += size
	}

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlack))
	s.Clear()

	for i := range diskMap.Cells {
		x, y := indexToCoord(i)
		s.SetContent(x, y, ' ', nil, tcell.StyleDefault)
	}

	diskMap.Paint(s)

	quit := func() {
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	diskMap.DefragPaint(s)

	for {
		s.Show()
		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				s.Clear()
			}
		}
	}
}

const (
	animationDuration = 1 * time.Millisecond
	animationFrames   = 5
)

var (
	animationState = 0
)

func updateCell(s tcell.Screen, x, y int, style tcell.Style) {
	for i := 0; i < animationFrames; i++ {
		animationState = (animationState + 1) % animationFrames

		time.Sleep(animationDuration)
	}

	s.SetContent(x, y, ' ', nil, style)
}

func generateRandomColor() tcell.Color {
	r := int32(rand.Intn(256))
	g := int32(rand.Intn(256))
	b := int32(rand.Intn(256))

	color := tcell.NewRGBColor(r, g, b)

	return color
}
