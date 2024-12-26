package main

import (
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"slices"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
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

	SHEEP_SIZE         = 128
	SHEEP1_FRAME_COUNT = 7
	SHEEP2_FRAME_COUNT = 6
)

var (
	sheepSheet  = utils.GetSpriteSheet(resources.HappySheepAll)
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

	shadows   = utils.GetSpriteSheet(resources.Shadows)
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
	// blend
	ground2Blend = groundSheet.SubImage(image.Rect(9*TILEMAP_SIZE, 0*TILEMAP_SIZE, 10*TILEMAP_SIZE, 1*TILEMAP_SIZE)).(*ebiten.Image)

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

	deco01 = utils.GetSpriteSheet(resources.Deco01)
	deco02 = utils.GetSpriteSheet(resources.Deco02)
	deco03 = utils.GetSpriteSheet(resources.Deco03)
	deco04 = utils.GetSpriteSheet(resources.Deco04)
	deco05 = utils.GetSpriteSheet(resources.Deco05)
	deco06 = utils.GetSpriteSheet(resources.Deco06)
	deco07 = utils.GetSpriteSheet(resources.Deco07)
	deco08 = utils.GetSpriteSheet(resources.Deco08)
	deco09 = utils.GetSpriteSheet(resources.Deco09)
	deco10 = utils.GetSpriteSheet(resources.Deco10)
	deco11 = utils.GetSpriteSheet(resources.Deco11)
	deco12 = utils.GetSpriteSheet(resources.Deco12)
	deco13 = utils.GetSpriteSheet(resources.Deco13)
	deco14 = utils.GetSpriteSheet(resources.Deco14)
	deco15 = utils.GetSpriteSheet(resources.Deco15)
	deco16 = utils.GetSpriteSheet(resources.Deco16)
	deco17 = utils.GetSpriteSheet(resources.Deco17)
	deco18 = utils.GetSpriteSheet(resources.Deco18)

	decorations = []*ebiten.Image{deco01, deco02, deco03, deco04, deco05, deco06, deco07, deco08, deco09, deco10, deco11, deco12, deco13, deco14, deco15, deco16, deco17, deco18}

	ArrowUp    = utils.GetSpriteSheet(resources.ArrowUp)
	ArrowDown  = utils.GetSpriteSheet(resources.ArrowDown)
	ArrowLeft  = utils.GetSpriteSheet(resources.ArrowLeft)
	ArrowRight = utils.GetSpriteSheet(resources.ArrowRight)

	BlueButton        = utils.GetSpriteSheet(resources.BlueButton)
	BlueButtonPressed = utils.GetSpriteSheet(resources.BlueButtonPressed)

	buttons = []*ebiten.Image{BlueButton, BlueButtonPressed}
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

	if rand.Intn(8) == 0 {
		index := rand.Intn(18)
		w.BlockadeDecors[y][x] = &index
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

	if o != nil && o.Char == "#" {
		return
	}

	if rand.Intn(8) == 0 {
		index := rand.Intn(15)
		w.GroundDecors[y][x] = &index
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
		w.GroundDecors = append(w.GroundDecors, make([]*int, w.Width))
		w.BlockadeDecors = append(w.BlockadeDecors, make([]*int, w.Width))
		for x, o := range line {
			w.SetupGroundSprite(x, y, o)

			if o != nil {
				if o.Char == "#" && !o.IsWater {
					w.SetupBlockadeSprite(x, y)
				}
				if o.Char == "[" || o.Char == "]" {
					if utils.RandBool() {
						o.Sheep1Sprite = true
					} else {
						o.Sheep2Sprite = true
					}
				}
			}
		}
	}

	return w
}

var (
	maxTilesWidth  = 21
	maxTilesHeight = 17
	camX           = 0
	camY           = 0
	animation      = 0
	animationState = 0
	reverse        = false
	offsetX        = 48
	offsetY        = 48
	direction      = ""
)

func (w *Warehouse) GetDrawDimensions() (int, int, int, int) {
	minY := max(0, camY)
	if minY < 0 {
		minY = 0
	}
	maxY := min(minY+maxTilesHeight, w.Height)
	minX := max(0, camX)
	if minX < 0 {
		minX = 0
	}
	maxX := min(minX+maxTilesWidth, w.Width)
	return minX, minY, maxX, maxY
}

func (w *Warehouse) DrawWaters(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	op.GeoM.Reset()
	op.GeoM.Scale(float64(maxTilesWidth/2), float64(maxTilesHeight/2+1))

	screen.DrawImage(water, op)
}

func (w *Warehouse) DrawFoams(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	minX, minY, maxX, maxY := w.GetDrawDimensions()
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			o := w.Map[y][x]
			if o == nil || (o != nil && o.Couple == nil && !o.IsWater) {
				op.GeoM.Reset()
				op.GeoM.Scale(.5, .5)
				op.GeoM.Translate(float64((x-camX)*32-offsetX), float64((y-camY)*32-offsetY))

				i := (animation / 5) % FOAM_FRAME_COUNT
				screen.DrawImage(foamSheet.SubImage(image.Rect(i*FOAM_SIZE, 0, (i+1)*FOAM_SIZE, FOAM_SIZE)).(*ebiten.Image), op)
			}
		}
	}
}

func (w *Warehouse) DrawGround(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	minX, minY, maxX, maxY := w.GetDrawDimensions()
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			sprite := w.GroundSprites[y][x]
			if sprite != nil {
				op.GeoM.Reset()
				op.GeoM.Scale(.5, .5)
				op.GeoM.Translate(float64((x-camX)*32-(offsetX-32)), float64((y-camY)*32-(offsetY-32)))
				screen.DrawImage(sprite, op)
			}

			index := w.GroundDecors[y][x]
			if index != nil {
				op.GeoM.Reset()
				op.GeoM.Scale(.5, .5)
				op.GeoM.Translate(float64((x-camX)*32-(offsetX-32)), float64((y-camY)*32-(offsetY-32)))

				screen.DrawImage(decorations[*index], op)
			}
		}
	}
}

func (w *Warehouse) DrawBlockade(screen *ebiten.Image, object *Object, op *ebiten.DrawImageOptions) {
	if len(object.Sprites) > 0 {
		op.GeoM.Reset()
		op.GeoM.Scale(.5, .5)
		op.GeoM.Translate(float64((object.X-camX-1)*32-(offsetX-32)), float64((object.Y-camY)*32-offsetY))
		screen.DrawImage(shadows, op)

		op.GeoM.Reset()
		op.GeoM.Scale(.5, .5)
		op.GeoM.Translate(float64((object.X-camX)*32-(offsetX-32)), float64((object.Y-camY)*32-offsetY))
		screen.DrawImage(object.Sprites[0], op)

		op.GeoM.Reset()
		op.GeoM.Scale(.5, .5)
		op.GeoM.Translate(float64((object.X-camX)*32-(offsetX-32)), float64((object.Y-camY)*32-offsetY))
		screen.DrawImage(object.Sprites[1], op)

		op.GeoM.Reset()
		op.GeoM.Scale(.5, .5)
		op.GeoM.Translate(float64((object.X-camX)*32-(offsetX-32)), float64((object.Y-camY+1)*32-offsetY))
		screen.DrawImage(ground2Blend, op)

	}

	index := w.BlockadeDecors[object.Y][object.X]
	if index != nil {
		op.GeoM.Reset()
		op.GeoM.Scale(.5, .5)
		if *index < 15 {
			op.GeoM.Translate(float64((object.X-camX)*32-(offsetX-32)), float64((object.Y-camY)*32-offsetY))
		} else if *index < 17 {
			op.GeoM.Translate(float64((object.X-camX)*32-(offsetX-32)), float64((object.Y-camY-1)*32-offsetY))
		} else {
			op.GeoM.Translate(float64((object.X-camX-1)*32-(offsetX-32)), float64((object.Y-camY-2)*32-offsetY))
		}
		screen.DrawImage(decorations[*index], op)
	}
}

func (w *Warehouse) DrawDecorations(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	for y, row := range w.GroundDecors {
		for x, index := range row {
			if x < camX || x >= camX+maxTilesWidth || y < camY || y >= camY+maxTilesHeight || index == nil {
				continue
			}

			op.GeoM.Reset()
			op.GeoM.Scale(.5, .5)
			op.GeoM.Translate(float64((x-camX)*32), float64((y-camY)*32-32))

			screen.DrawImage(decorations[*index], op)
		}
	}
}

func (w *Warehouse) DrawMovables(screen *ebiten.Image, x, y int, op *ebiten.DrawImageOptions) {
	op.GeoM.Reset()
	op.GeoM.Scale(.5, .5)
	op.GeoM.Translate(float64((x-camX-1)*32), float64((y-camY-1)*32))

	o := w.Map[y][x]
	if o == nil {
		return
	}
	if o.Sheep1Sprite {
		i := (animation / 5) % SHEEP1_FRAME_COUNT
		screen.DrawImage(sheepSheet.SubImage(image.Rect(i*SHEEP_SIZE, 0, (i+1)*SHEEP_SIZE, SHEEP_SIZE)).(*ebiten.Image), op)
	}

	if o.Sheep2Sprite {
		i := (animation / 5) % SHEEP2_FRAME_COUNT
		screen.DrawImage(sheepSheet.SubImage(image.Rect(i*SHEEP_SIZE, SHEEP_SIZE, (i+1)*SHEEP_SIZE, 2*SHEEP_SIZE)).(*ebiten.Image), op)
	}
}

func (w *Warehouse) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w.DrawWaters(screen, op)
	w.DrawFoams(screen, op)
	w.DrawGround(screen, op)
	// w.DrawDecorations(screen, op)

	minX, minY, maxX, maxY := w.GetDrawDimensions()
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			o := w.Map[y][x]

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
				op.GeoM.Translate(float64((o.X-camX)*32-48), float64((o.Y-camY)*32-48))
				i := (animation / 5) % WARRIOR_FRAME_COUNT
				screen.DrawImage(warriorSheet.SubImage(image.Rect(i*WARRIOR_WIDTH, WARRIOR_HEIGHT, (i+1)*WARRIOR_WIDTH, 2*WARRIOR_HEIGHT)).(*ebiten.Image), op)
			} else if o.Char == "[" || o.Char == "]" {
				w.DrawMovables(screen, x, y, op)
			} else {
				w.DrawBlockade(screen, o, op)
			}
		}
	}

	op.GeoM.Reset()
	op.GeoM.Translate(float64(maxTilesWidth-3)*32-float64(offsetX-32), float64(maxTilesHeight-3)*32)
	i := (animation / 10) % 2
	screen.DrawImage(buttons[i], op)

	op.GeoM.Reset()
	op.GeoM.Scale(3, 3)
	op.GeoM.Translate(float64(maxTilesWidth-3)*32, float64(maxTilesHeight-3)*32)
	if direction == ">" {
		screen.DrawImage(ArrowRight, op)
	} else if direction == "<" {
		screen.DrawImage(ArrowLeft, op)
	} else if direction == "^" {
		screen.DrawImage(ArrowUp, op)
	} else if direction == "v" {
		screen.DrawImage(ArrowDown, op)
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

	direction = dir
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

	ebiten.SetWindowSize(640, 518)
	ebiten.SetWindowTitle("Advent of Code 2024 - Day 15 - Part 2")

	if err := ebiten.RunGame(&w); err != nil {
		log.Fatal(err)
	}
}
