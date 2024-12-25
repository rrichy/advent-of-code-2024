package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"
	"slices"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	resources "github.com/rrichy/advent-of-code-2024/assets/topdown/topdown/tiny_swords"
	"github.com/rrichy/advent-of-code-2024/utils"
)

var (
	groundSheet = utils.GetSpriteSheet(resources.GroundFlat)
	groundTL    = groundSheet.SubImage(image.Rect(0*64, 0*64, 1*64, 1*64)).(*ebiten.Image)
	groundT     = groundSheet.SubImage(image.Rect(1*64, 0*64, 2*64, 1*64)).(*ebiten.Image)
	groundTR    = groundSheet.SubImage(image.Rect(2*64, 0*64, 3*64, 1*64)).(*ebiten.Image)
	groundL     = groundSheet.SubImage(image.Rect(0*64, 1*64, 1*64, 2*64)).(*ebiten.Image)
	groundM     = groundSheet.SubImage(image.Rect(1*64, 1*64, 2*64, 2*64)).(*ebiten.Image)
	groundR     = groundSheet.SubImage(image.Rect(2*64, 1*64, 3*64, 2*64)).(*ebiten.Image)
	groundBL    = groundSheet.SubImage(image.Rect(0*64, 2*64, 1*64, 3*64)).(*ebiten.Image)
	groundB     = groundSheet.SubImage(image.Rect(1*64, 2*64, 2*64, 3*64)).(*ebiten.Image)
	groundBR    = groundSheet.SubImage(image.Rect(2*64, 2*64, 3*64, 3*64)).(*ebiten.Image)
)

func NewWarehouseTwiceAnimate(s string) Warehouse {
	sheepSheet := utils.GetSpriteSheet(resources.HappySheepAll)
	rocks3Sheet := utils.GetSpriteSheet(resources.Rocks2)
	warriorSheet := utils.GetSpriteSheet(resources.Warrior)

	sheep := sheepSheet.SubImage(image.Rect(0*128, 0*128, 1*128, 1*128)).(*ebiten.Image)
	rock := rocks3Sheet.SubImage(image.Rect(0*128, 0*128, 1*128, 1*128)).(*ebiten.Image)
	warrior := warriorSheet.SubImage(image.Rect(0*192, 0*192, 1*192, 1*192)).(*ebiten.Image)

	w := Warehouse{
		Map: [][]*Object{},
	}
	for y, line := range strings.Split(s, "\n") {
		l := []*Object{}
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				o1 := NewObject(x*2, y, false, char, rock)
				o2 := NewObject(x*2+1, y, false, char, rock)
				l = append(l, &o1, &o2)
			} else if char == "." {
				l = append(l, nil, nil)
			} else if char == "O" {
				o1 := NewObject(x*2, y, true, "[", sheep)
				o2 := NewObject(x*2+1, y, true, "]", sheep)
				o1.Couple = &o2
				o2.Couple = &o1
				l = append(l, &o1, &o2)
			} else {
				o := NewObject(x*2, y, true, char, warrior)
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

func (w *Warehouse) DrawGround(screen *ebiten.Image, x, y int, op *ebiten.DrawImageOptions) {
	op.GeoM.Reset()
	op.GeoM.Scale(.5, .5)
	op.GeoM.Translate(float64((x-camX)*32), float64((y-camY)*32))

	var sprite *ebiten.Image

	if y == 0 {
		if x == 0 {
			sprite = groundTL
		} else if x == w.Width-1 {
			sprite = groundTR
		} else {
			sprite = groundT
		}
	} else if y == w.Height-1 {
		if x == 0 {
			sprite = groundBL
		} else if x == w.Width-1 {
			sprite = groundBR
		} else {
			sprite = groundB
		}
	} else {
		if x == 0 {
			sprite = groundL
		} else if x == w.Width-1 {
			sprite = groundR
		} else {
			sprite = groundM
		}
	}

	screen.DrawImage(sprite, op)
}

var (
	maxTilesWidth  = 20
	maxTilesHeight = 16
	camX           = 0
	camY           = 0
)

// Draw implements ebiten.Game.
func (w *Warehouse) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for y, row := range w.Map {
		for x, o := range row {
			if x < camX || x >= camX+maxTilesWidth || y < camY || y >= camY+maxTilesHeight {
				continue
			}

			w.DrawGround(screen, x, y, op)

			if o == nil {
				continue
			}

			op.GeoM.Reset()

			if o.Char == "@" {
				op.GeoM.Scale(.33, .33)
				op.GeoM.Translate(float64((o.X-camX)*32-16), float64((o.Y-camY)*32-16))
				screen.DrawImage(o.Sprite, op)
			} else if o.Char == "[" {
				vector.DrawFilledRect(screen, float32((x-camX)*32), float32((y-camY)*32), 32, 32, color.Gray16{}, true)
			} else if o.Char == "]" {
				vector.DrawFilledRect(screen, float32((x-camX)*32), float32((y-camY)*32), 32, 32, color.Gray{}, true)
			} else {
				op.GeoM.Scale(.5, .5)
				op.GeoM.Translate(float64((o.X-camX)*32-16), float64((o.Y-camY)*32-16))
				screen.DrawImage(o.Sprite, op)
			}
		}

	}
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("Seconds: %d", h.Seconds))
}

func (w *Warehouse) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return outsideWidth, outsideHeight
}

var animationState = 0

// var keyPressThreshold = 0

func (w *Warehouse) Update() error {

	// if ebiten.IsKeyPressed(ebiten.KeyUp) {
	// 	keyPressThreshold = (keyPressThreshold + 1) % 10
	// 	if keyPressThreshold == 0 {
	// 		w.Commands = append(w.Commands, "^")
	// 	}
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyDown) {
	// 	keyPressThreshold = (keyPressThreshold + 1) % 10
	// 	if keyPressThreshold == 0 {
	// 		w.Commands = append(w.Commands, "v")
	// 	}
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyLeft) {
	// 	keyPressThreshold = (keyPressThreshold + 1) % 10
	// 	if keyPressThreshold == 0 {
	// 		w.Commands = append(w.Commands, "<")
	// 	}
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyRight) {
	// 	keyPressThreshold = (keyPressThreshold + 1) % 10
	// 	if keyPressThreshold == 0 {
	// 		w.Commands = append(w.Commands, ">")
	// 	}
	// }

	animationState = (animationState + 1) % 20
	if len(w.Commands) == 0 || animationState != 0 {
		return nil
	}

	dir, slice := w.Commands[len(w.Commands)-1], w.Commands[:len(w.Commands)-1]
	w.Commands = slice

	dc := GetDisplacement(dir)
	w.Move2(w.Robot, &dc)

	camX = w.Robot.X - maxTilesWidth/2
	camY = w.Robot.Y - maxTilesHeight/2

	if camX < 0 {
		camX = 0
	}
	if camY < 0 {
		camY = 0
	}
	if camX+maxTilesWidth > w.Width {
		camX = w.Width - maxTilesWidth
	}
	if camY+maxTilesHeight > w.Height {
		camY = w.Height - maxTilesHeight
	}

	// if camX+maxTilesWidth > w.Width {
	// 	camX = w.Width - maxTilesWidth
	// }
	// if camY+maxTilesHeight > w.Height {
	// 	camY = w.Height - maxTilesHeight
	// }

	return nil
}

func Animate() {
	s := strings.Split(sample1, "\n\n")
	w := NewWarehouseTwiceAnimate(s[0])
	// w.Commands = []string{">", "<"}
	w.Commands = strings.Split(strings.ReplaceAll(s[1], "\n", ""), "")
	slices.Reverse(w.Commands)

	// w.Commands = strings.Split(s[1], "\n")

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Advent of Code 2024 - Day 15 - Part 2")

	if err := ebiten.RunGame(&w); err != nil {
		log.Fatal(err)
	}
}
