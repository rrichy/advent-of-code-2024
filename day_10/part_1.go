package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rrichy/advent-of-code-2024/assets/isometric"
)

type Coordinate struct {
	X int
	Y int
}

type TrailHead struct {
	Coordinate
	Score int
}

const tileSizeW = 256
const tileSizeH = 192
const trailHeadSizeW = 256
const trailHeadSizeH = 128

type Tile struct {
	Elevation int
	Sprite    *ebiten.Image
}

func NewTile(elevation int, sheet *ebiten.Image) Tile {
	x := rand.Intn(8)
	y := 0

	sprite := sheet.SubImage(image.Rect(x*tileSizeW, y*tileSizeH, (x+1)*tileSizeW, (y+1)*tileSizeH)).(*ebiten.Image)
	return Tile{Elevation: elevation, Sprite: sprite}
}

type Topography struct {
	Tiles           [][]Tile
	TrailHeads      []*TrailHead
	Width           int
	Height          int
	WorldWidth      float64
	WorldHeight     float64
	FlagSprite      *ebiten.Image
	TrailHeadSprite *ebiten.Image

	// Animation
	CX float64
	CY float64
	SX float64
	SY float64
}

func NewTopography(input string) Topography {
	img1, _, err := image.Decode(bytes.NewReader(isometric.Tiles))
	if err != nil {
		log.Fatal(err)
	}

	tileSheet := ebiten.NewImageFromImage(img1)

	img2, _, err := image.Decode(bytes.NewReader(isometric.TileOverlays))
	if err != nil {
		log.Fatal(err)
	}

	tileOverlaySheet := ebiten.NewImageFromImage(img2)
	trailHeadSprite := tileOverlaySheet.SubImage(image.Rect(1*trailHeadSizeW, 5*trailHeadSizeH, 2*trailHeadSizeW, 6*trailHeadSizeH)).(*ebiten.Image)

	img3, _, err := image.Decode(bytes.NewReader(isometric.Objects))
	if err != nil {
		log.Fatal(err)
	}

	objectSheet := ebiten.NewImageFromImage(img3)
	flagSprite := objectSheet.SubImage(image.Rect(0, 256, 256, 512)).(*ebiten.Image)

	tiles := [][]Tile{}
	trailHeads := []*TrailHead{}
	for y, line := range strings.Split(input, "\n") {
		row := []Tile{}
		for x, char := range strings.Split(line, "") {
			c, _ := strconv.Atoi(char)
			row = append(row, NewTile(c, tileSheet))

			if c == 0 {
				trailHeads = append(trailHeads, &TrailHead{Coordinate: Coordinate{X: x, Y: y}})
			}
		}

		tiles = append(tiles, row)
	}

	width := len(tiles[0])
	height := len(tiles)
	worldWidth := float64((width + height) * tileSizeW / 2)
	worldHeight := float64((width + height) * tileSizeH / 2)

	return Topography{
		Tiles:           tiles,
		TrailHeads:      trailHeads,
		Width:           width,
		Height:          height,
		WorldWidth:      worldWidth,
		WorldHeight:     worldHeight,
		FlagSprite:      flagSprite,
		TrailHeadSprite: trailHeadSprite,
	}
}

func (t *Topography) RateTrailHeadsPart1() {
	for _, trailHead := range t.TrailHeads {
		trailHead.Score = t.TraversePart1(trailHead.Coordinate, &map[Coordinate]bool{})
	}
}

func (t *Topography) GetTrailHeadsTotalScore() int {
	totalScore := 0
	for _, trailHead := range t.TrailHeads {
		totalScore += trailHead.Score
	}

	return totalScore
}

func (t *Topography) IsOutOfBounds(c Coordinate) bool {
	return c.X < 0 || c.X >= t.Width || c.Y < 0 || c.Y >= t.Height
}

func (t *Topography) IsTraversable(c1, c2 Coordinate) bool {
	currentElevation := t.Tiles[c1.Y][c1.X].Elevation
	return t.Tiles[c2.Y][c2.X].Elevation-currentElevation == 1
}

func (t *Topography) TraversePart1(c Coordinate, m *map[Coordinate]bool) int {
	if t.Tiles[c.Y][c.X].Elevation == 9 {
		if (*m)[c] {
			return 0
		}
		(*m)[c] = true

		return 1
	}

	r, l, u, d := 0, 0, 0, 0

	// Go right
	right := Coordinate{X: c.X + 1, Y: c.Y}
	if !t.IsOutOfBounds(right) && t.IsTraversable(c, right) {
		r = t.TraversePart1(right, m)
	}

	// Go left
	left := Coordinate{X: c.X - 1, Y: c.Y}
	if !t.IsOutOfBounds(left) && t.IsTraversable(c, left) {
		l = t.TraversePart1(left, m)
	}

	// Go up
	up := Coordinate{X: c.X, Y: c.Y - 1}
	if !t.IsOutOfBounds(up) && t.IsTraversable(c, up) {
		u = t.TraversePart1(up, m)
	}

	// Go down
	down := Coordinate{X: c.X, Y: c.Y + 1}
	if !t.IsOutOfBounds(down) && t.IsTraversable(c, down) {
		d = t.TraversePart1(down, m)
	}

	return r + l + u + d
}

func Part1() int {
	topography := NewTopography(input)
	topography.RateTrailHeadsPart1()

	total := topography.GetTrailHeadsTotalScore()

	log.Print(total)

	return total
}
