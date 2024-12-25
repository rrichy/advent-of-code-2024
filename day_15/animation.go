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

const (
	WARRIOR_FRAME_COUNT = 6
	WARRIOR_WIDTH       = 192
	WARRIOR_HEIGHT      = 192

	TILEMAP_SIZE     = 64
	FOAM_SIZE        = 192
	FOAM_FRAME_COUNT = 8
)

var (
	water       = utils.GetSpriteSheet(resources.Water).SubImage(image.Rect(0, 0, TILEMAP_SIZE, TILEMAP_SIZE)).(*ebiten.Image)
	foamSheet   = utils.GetSpriteSheet(resources.Foam)
	groundSheet = utils.GetSpriteSheet(resources.Tilemap_Flat)
	groundTL    = groundSheet.SubImage(image.Rect(0*TILEMAP_SIZE, 0*TILEMAP_SIZE, 1*TILEMAP_SIZE, 1*TILEMAP_SIZE)).(*ebiten.Image)
	groundT     = groundSheet.SubImage(image.Rect(1*TILEMAP_SIZE, 0*TILEMAP_SIZE, 2*TILEMAP_SIZE, 1*TILEMAP_SIZE)).(*ebiten.Image)
	groundTR    = groundSheet.SubImage(image.Rect(2*TILEMAP_SIZE, 0*TILEMAP_SIZE, 3*TILEMAP_SIZE, 1*TILEMAP_SIZE)).(*ebiten.Image)
	groundL     = groundSheet.SubImage(image.Rect(0*TILEMAP_SIZE, 1*TILEMAP_SIZE, 1*TILEMAP_SIZE, 2*TILEMAP_SIZE)).(*ebiten.Image)
	groundM     = groundSheet.SubImage(image.Rect(1*TILEMAP_SIZE, 1*TILEMAP_SIZE, 2*TILEMAP_SIZE, 2*TILEMAP_SIZE)).(*ebiten.Image)
	groundR     = groundSheet.SubImage(image.Rect(2*TILEMAP_SIZE, 1*TILEMAP_SIZE, 3*TILEMAP_SIZE, 2*TILEMAP_SIZE)).(*ebiten.Image)
	groundBL    = groundSheet.SubImage(image.Rect(0*TILEMAP_SIZE, 2*TILEMAP_SIZE, 1*TILEMAP_SIZE, 3*TILEMAP_SIZE)).(*ebiten.Image)
	groundB     = groundSheet.SubImage(image.Rect(1*TILEMAP_SIZE, 2*TILEMAP_SIZE, 2*TILEMAP_SIZE, 3*TILEMAP_SIZE)).(*ebiten.Image)
	groundBR    = groundSheet.SubImage(image.Rect(2*TILEMAP_SIZE, 2*TILEMAP_SIZE, 3*TILEMAP_SIZE, 3*TILEMAP_SIZE)).(*ebiten.Image)
	// 3x1
	ground3X1A = groundSheet.SubImage(image.Rect(0*TILEMAP_SIZE, 3*TILEMAP_SIZE, 1*TILEMAP_SIZE, 4*TILEMAP_SIZE)).(*ebiten.Image)
	ground3X1B = groundSheet.SubImage(image.Rect(1*TILEMAP_SIZE, 3*TILEMAP_SIZE, 2*TILEMAP_SIZE, 4*TILEMAP_SIZE)).(*ebiten.Image)
	ground3X1C = groundSheet.SubImage(image.Rect(2*TILEMAP_SIZE, 3*TILEMAP_SIZE, 3*TILEMAP_SIZE, 4*TILEMAP_SIZE)).(*ebiten.Image)
	// 1x2
	ground1X3A = groundSheet.SubImage(image.Rect(3*TILEMAP_SIZE, 0*TILEMAP_SIZE, 4*TILEMAP_SIZE, 1*TILEMAP_SIZE)).(*ebiten.Image)
	ground1X3B = groundSheet.SubImage(image.Rect(3*TILEMAP_SIZE, 1*TILEMAP_SIZE, 4*TILEMAP_SIZE, 2*TILEMAP_SIZE)).(*ebiten.Image)
	ground1X3C = groundSheet.SubImage(image.Rect(3*TILEMAP_SIZE, 2*TILEMAP_SIZE, 4*TILEMAP_SIZE, 3*TILEMAP_SIZE)).(*ebiten.Image)
	groundS    = groundSheet.SubImage(image.Rect(3*TILEMAP_SIZE, 3*TILEMAP_SIZE, 4*TILEMAP_SIZE, 4*TILEMAP_SIZE)).(*ebiten.Image)

	ground2TL = groundSheet.SubImage(image.Rect(5*TILEMAP_SIZE, 0*TILEMAP_SIZE, 6*TILEMAP_SIZE, 1*TILEMAP_SIZE)).(*ebiten.Image)
	ground2T  = groundSheet.SubImage(image.Rect(6*TILEMAP_SIZE, 0*TILEMAP_SIZE, 7*TILEMAP_SIZE, 1*TILEMAP_SIZE)).(*ebiten.Image)
	ground2TR = groundSheet.SubImage(image.Rect(7*TILEMAP_SIZE, 0*TILEMAP_SIZE, 8*TILEMAP_SIZE, 1*TILEMAP_SIZE)).(*ebiten.Image)
	ground2L  = groundSheet.SubImage(image.Rect(5*TILEMAP_SIZE, 1*TILEMAP_SIZE, 6*TILEMAP_SIZE, 2*TILEMAP_SIZE)).(*ebiten.Image)
	ground2M  = groundSheet.SubImage(image.Rect(6*TILEMAP_SIZE, 1*TILEMAP_SIZE, 7*TILEMAP_SIZE, 2*TILEMAP_SIZE)).(*ebiten.Image)
	ground2R  = groundSheet.SubImage(image.Rect(7*TILEMAP_SIZE, 1*TILEMAP_SIZE, 8*TILEMAP_SIZE, 2*TILEMAP_SIZE)).(*ebiten.Image)
	ground2BL = groundSheet.SubImage(image.Rect(5*TILEMAP_SIZE, 2*TILEMAP_SIZE, 6*TILEMAP_SIZE, 3*TILEMAP_SIZE)).(*ebiten.Image)
	ground2B  = groundSheet.SubImage(image.Rect(6*TILEMAP_SIZE, 2*TILEMAP_SIZE, 7*TILEMAP_SIZE, 3*TILEMAP_SIZE)).(*ebiten.Image)
	ground2BR = groundSheet.SubImage(image.Rect(7*TILEMAP_SIZE, 2*TILEMAP_SIZE, 8*TILEMAP_SIZE, 3*TILEMAP_SIZE)).(*ebiten.Image)
	// 3x1
	ground23X1A = groundSheet.SubImage(image.Rect(5*TILEMAP_SIZE, 3*TILEMAP_SIZE, 6*TILEMAP_SIZE, 4*TILEMAP_SIZE)).(*ebiten.Image)
	ground23X1B = groundSheet.SubImage(image.Rect(6*TILEMAP_SIZE, 3*TILEMAP_SIZE, 7*TILEMAP_SIZE, 4*TILEMAP_SIZE)).(*ebiten.Image)
	ground23X1C = groundSheet.SubImage(image.Rect(7*TILEMAP_SIZE, 3*TILEMAP_SIZE, 8*TILEMAP_SIZE, 4*TILEMAP_SIZE)).(*ebiten.Image)
	// 1x2
	ground21X3A = groundSheet.SubImage(image.Rect(8*TILEMAP_SIZE, 0*TILEMAP_SIZE, 9*TILEMAP_SIZE, 1*TILEMAP_SIZE)).(*ebiten.Image)
	ground21X3B = groundSheet.SubImage(image.Rect(8*TILEMAP_SIZE, 1*TILEMAP_SIZE, 9*TILEMAP_SIZE, 2*TILEMAP_SIZE)).(*ebiten.Image)
	ground21X3C = groundSheet.SubImage(image.Rect(8*TILEMAP_SIZE, 2*TILEMAP_SIZE, 9*TILEMAP_SIZE, 3*TILEMAP_SIZE)).(*ebiten.Image)
	ground2S    = groundSheet.SubImage(image.Rect(8*TILEMAP_SIZE, 3*TILEMAP_SIZE, 9*TILEMAP_SIZE, 4*TILEMAP_SIZE)).(*ebiten.Image)

	warriorSheet          = utils.GetSpriteSheet(resources.Warrior)
	tilemapElevationSheet = utils.GetSpriteSheet(resources.Tilemap_Elevation)
	// 3x2
	elevationTL = tilemapElevationSheet.SubImage(image.Rect(0*TILEMAP_SIZE, 0*2*TILEMAP_SIZE, 1*TILEMAP_SIZE, 1*2*TILEMAP_SIZE)).(*ebiten.Image)
	elevationT  = tilemapElevationSheet.SubImage(image.Rect(1*TILEMAP_SIZE, 0*2*TILEMAP_SIZE, 2*TILEMAP_SIZE, 1*2*TILEMAP_SIZE)).(*ebiten.Image)
	elevationTR = tilemapElevationSheet.SubImage(image.Rect(2*TILEMAP_SIZE, 0*2*TILEMAP_SIZE, 3*TILEMAP_SIZE, 1*2*TILEMAP_SIZE)).(*ebiten.Image)
	elevationBL = tilemapElevationSheet.SubImage(image.Rect(0*TILEMAP_SIZE, 1*2*TILEMAP_SIZE, 1*TILEMAP_SIZE, 2*2*TILEMAP_SIZE)).(*ebiten.Image)
	elevationB  = tilemapElevationSheet.SubImage(image.Rect(1*TILEMAP_SIZE, 1*2*TILEMAP_SIZE, 2*TILEMAP_SIZE, 2*2*TILEMAP_SIZE)).(*ebiten.Image)
	elevationBR = tilemapElevationSheet.SubImage(image.Rect(2*TILEMAP_SIZE, 1*2*TILEMAP_SIZE, 3*TILEMAP_SIZE, 2*2*TILEMAP_SIZE)).(*ebiten.Image)
	// 1x2
	elevation1X2A = tilemapElevationSheet.SubImage(image.Rect(3*TILEMAP_SIZE, 0*2*TILEMAP_SIZE, 4*TILEMAP_SIZE, 1*2*TILEMAP_SIZE)).(*ebiten.Image)
	elevation1X2B = tilemapElevationSheet.SubImage(image.Rect(3*TILEMAP_SIZE, 1*2*TILEMAP_SIZE, 4*TILEMAP_SIZE, 2*2*TILEMAP_SIZE)).(*ebiten.Image)
	// 3x1
	elevation3X1A = tilemapElevationSheet.SubImage(image.Rect(0*TILEMAP_SIZE, 2*2*TILEMAP_SIZE, 1*TILEMAP_SIZE, 3*2*TILEMAP_SIZE)).(*ebiten.Image)
	elevation3X1B = tilemapElevationSheet.SubImage(image.Rect(1*TILEMAP_SIZE, 2*2*TILEMAP_SIZE, 2*TILEMAP_SIZE, 3*2*TILEMAP_SIZE)).(*ebiten.Image)
	elevation3X1C = tilemapElevationSheet.SubImage(image.Rect(2*TILEMAP_SIZE, 2*2*TILEMAP_SIZE, 3*TILEMAP_SIZE, 3*2*TILEMAP_SIZE)).(*ebiten.Image)
	// 1x1
	elevationS = tilemapElevationSheet.SubImage(image.Rect(3*TILEMAP_SIZE, 2*2*TILEMAP_SIZE, 4*TILEMAP_SIZE, 3*2*TILEMAP_SIZE)).(*ebiten.Image)
)

func (w *Warehouse) NeighbourIsHill(x, y, dx, dy int) bool {
	c := Coordinate{x + dx, y + dy}
	if w.IsOutOfBounds(c) {
		return false
	}
	n := w.Map[c.Y][c.X]
	if n == nil {
		return false
	}

	return n.Char == "#" && !n.IsWater
}

func (w *Warehouse) NeighbourIsPond(x, y, dx, dy int) bool {
	c := Coordinate{x + dx, y + dy}
	if w.IsOutOfBounds(c) {
		return false
	}
	n := w.Map[c.Y][c.X]
	if n == nil {
		return false
	}

	return n.Char == "#" && n.IsWater
}

func (w *Warehouse) SetupBlockadeSprite(x, y int) {
	t := w.NeighbourIsHill(x, y, 0, -1)
	l := w.NeighbourIsHill(x, y, -1, 0)
	r := w.NeighbourIsHill(x, y, 1, 0)
	b := w.NeighbourIsHill(x, y, 0, 1)

	// setups a hill else pond will be drawn
	if t && l && r && b {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevationB, groundM)
	} else if t && r && b {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevationBL, groundL)
	} else if t && l && b {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevationBR, groundR)
	} else if l && r && b {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevationT, groundT, groundM)
	} else if t && l && r {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevationB, groundB)
	} else if t && l {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevationBR, groundBR)
	} else if t && r {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevationBL, groundBL)
	} else if b && l {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevationTR, groundTR, groundR)
	} else if b && r {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevationTL, groundTL, groundR)
	} else if l && r {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevation3X1B, ground3X1B)
	} else if t && b {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevation1X2A, ground1X3B)
	} else if t {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevation1X2B, ground1X3C)
	} else if b {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevation1X2B, ground1X3A)
	} else if l {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevation3X1C, ground3X1C)
	} else if r {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevation3X1A, ground3X1A)
	} else {
		w.Map[y][x].Sprites = append(w.Map[y][x].Sprites, elevationS, groundS)
	}
}

func (w *Warehouse) SetupGroundSprite(x, y int, o *Object) {
	if o != nil && o.IsWater {
		return
	}

	t := w.NeighbourIsPond(x, y, 0, -1)
	l := w.NeighbourIsPond(x, y, -1, 0)
	r := w.NeighbourIsPond(x, y, 1, 0)
	b := w.NeighbourIsPond(x, y, 0, 1)

	if t && l && r && b {
		w.GroundSprites[y][x] = ground2S
	} else if t && l && r {
		w.GroundSprites[y][x] = ground21X3A
	} else if t && l && b {
		w.GroundSprites[y][x] = ground23X1A
	} else if t && r && b {
		w.GroundSprites[y][x] = ground23X1C
	} else if l && r && b {
		w.GroundSprites[y][x] = ground21X3C
	} else if t && l {
		w.GroundSprites[y][x] = ground2TL
	} else if t && r {
		w.GroundSprites[y][x] = ground2TR
	} else if b && l {
		w.GroundSprites[y][x] = ground2BL
	} else if b && r {
		w.GroundSprites[y][x] = ground2BR
	} else if l && r {
		w.GroundSprites[y][x] = ground21X3B
	} else if t && b {
		w.GroundSprites[y][x] = ground23X1B
	} else if t {
		w.GroundSprites[y][x] = ground2T
	} else if b {
		w.GroundSprites[y][x] = ground2B
	} else if l {
		w.GroundSprites[y][x] = ground2L
	} else if r {
		w.GroundSprites[y][x] = ground2R
	} else {
		w.GroundSprites[y][x] = ground2M
	}
}

func NewWarehouseWideAnimate(s string) Warehouse {
	w := Warehouse{
		Map: [][]*Object{},
	}
	for y, line := range strings.Split(s, "\n") {
		l := []*Object{}
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				o1 := NewObject(x*2, y, false, char)

				if utils.RandBool() {
					o1.IsWater = true
				}

				o2 := NewObject(x*2+1, y, false, char)
				if utils.RandBool() {
					o2.IsWater = true
				}

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

	for y, line := range w.Map {
		w.GroundSprites = append(w.GroundSprites, make([]*ebiten.Image, w.Width))
		for x, o := range line {
			w.SetupGroundSprite(x, y, o)
			if o != nil && o.Char == "#" && !o.IsWater {
				w.SetupBlockadeSprite(x, y)
			}
		}
	}

	return w
}

var (
	maxTilesWidth  = 20
	maxTilesHeight = 16
	camX           = 0
	camY           = 0
	animation      = 0
	animationState = 0
	reverse        = false
)

func (w *Warehouse) DrawWaters(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	for y, row := range w.Map {
		for x := range row {
			if x < camX || x >= camX+maxTilesWidth || y < camY || y >= camY+maxTilesHeight {
				continue
			}

			op.GeoM.Reset()
			op.GeoM.Scale(.5, .5)
			op.GeoM.Translate(float64((x-camX)*32), float64((y-camY)*32))

			screen.DrawImage(water, op)
		}
	}
}

func (w *Warehouse) DrawFoams(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	for y, row := range w.Map {
		for x, o := range row {
			if x < camX || x >= camX+maxTilesWidth || y < camY || y >= camY+maxTilesHeight {
				continue
			}

			if o == nil || (o != nil && o.Couple == nil && !o.IsWater) {
				op.GeoM.Reset()
				op.GeoM.Scale(.5, .5)
				op.GeoM.Translate(float64((x-camX)*32-32), float64((y-camY)*32-32))

				i := (animation / 5) % FOAM_FRAME_COUNT
				screen.DrawImage(foamSheet.SubImage(image.Rect(i*FOAM_SIZE, 0, (i+1)*FOAM_SIZE, FOAM_SIZE)).(*ebiten.Image), op)
			}
		}
	}
}

func (w *Warehouse) DrawGround(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	for y, row := range w.GroundSprites {
		for x, sprite := range row {
			if x < camX || x >= camX+maxTilesWidth || y < camY || y >= camY+maxTilesHeight {
				continue
			}

			if sprite != nil {
				op.GeoM.Reset()
				op.GeoM.Scale(.5, .5)
				op.GeoM.Translate(float64((x-camX)*32), float64((y-camY)*32))
				screen.DrawImage(sprite, op)
			}
		}
	}

}

func (w *Warehouse) DrawBlockade(screen *ebiten.Image, object *Object, op *ebiten.DrawImageOptions) {
	if len(object.Sprites) > 0 {
		op.GeoM.Reset()
		op.GeoM.Scale(.5, .5)
		op.GeoM.Translate(float64((object.X-camX)*32), float64((object.Y-camY)*32-32))
		screen.DrawImage(object.Sprites[0], op)

		op.GeoM.Reset()
		op.GeoM.Scale(.5, .5)
		op.GeoM.Translate(float64((object.X-camX)*32), float64((object.Y-camY)*32-32))
		screen.DrawImage(object.Sprites[1], op)
	}
}

func (w *Warehouse) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w.DrawWaters(screen, op)
	w.DrawFoams(screen, op)
	w.DrawGround(screen, op)

	for y, row := range w.Map {
		for x, o := range row {
			if x < camX || x >= camX+maxTilesWidth || y < camY || y >= camY+maxTilesHeight {
				continue
			}

			if o == nil {
				continue
			}

			op.GeoM.Reset()

			if o.Char == "@" {
				if reverse {
					op.GeoM.Scale(-1, 1)
					op.GeoM.Translate(WARRIOR_WIDTH, 0)
				}
				op.GeoM.Scale(.5, .5)
				op.GeoM.Translate(float64((o.X-camX)*32-32), float64((o.Y-camY)*32-32))
				i := (animation / 5) % WARRIOR_FRAME_COUNT
				screen.DrawImage(warriorSheet.SubImage(image.Rect(i*WARRIOR_WIDTH, WARRIOR_HEIGHT, (i+1)*WARRIOR_WIDTH, 2*WARRIOR_HEIGHT)).(*ebiten.Image), op)
			} else if o.Char == "[" {
				vector.DrawFilledRect(screen, float32((x-camX)*32), float32((y-camY)*32), 32, 32, color.Gray16{}, true)
			} else if o.Char == "]" {
				vector.DrawFilledRect(screen, float32((x-camX)*32), float32((y-camY)*32), 32, 32, color.Gray{}, true)
			} else {
				w.DrawBlockade(screen, o, op)
			}
		}

	}
}

func (w *Warehouse) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (w *Warehouse) Update() error {
	animation++
	animationState = animation % 20
	if len(w.Commands) == 0 || animationState != 0 {
		return nil
	}

	dir, slice := w.Commands[len(w.Commands)-1], w.Commands[:len(w.Commands)-1]
	w.Commands = slice

	if dir == ">" {
		reverse = false
	} else if dir == "<" {
		reverse = true
	}
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

	return nil
}

func Animate() {
	s := strings.Split(input, "\n\n")
	w := NewWarehouseWideAnimate(s[0])
	w.Commands = strings.Split(strings.ReplaceAll(s[1], "\n", ""), "")
	slices.Reverse(w.Commands)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Advent of Code 2024 - Day 15 - Part 2")

	if err := ebiten.RunGame(&w); err != nil {
		log.Fatal(err)
	}
}
