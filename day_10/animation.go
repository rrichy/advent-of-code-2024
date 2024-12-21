package day10

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rrichy/advent-of-code-2024/utils"
)

func (t *Tile) Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions) {
	screen.DrawImage(t.Sprite, options)
}

func (t *Topography) cartesianToIso(x, y float64) (float64, float64) {
	xi := float64((x - y) * tileSizeW / 2)
	yi := float64((x + y) * tileSizeH / 3)
	return xi, yi
}

const ScreenWidth = 1280
const ScreenHeight = 720

var (
	maxX           = 1
	maxY           = 1
	animationState = 0
	mapperX        = map[int]int{}
)

func (t Topography) Draw(screen *ebiten.Image) {
	animationState = (animationState + 1) % 2
	op := &ebiten.DrawImageOptions{}

	for y := 0; y < maxY; y++ {
		_maxX := mapperX[y]
		if _maxX == 0 {
			_maxX = 1
		}

		for x := 0; x < _maxX; x++ {
			xi, yi := t.cartesianToIso(float64(x), float64(y))

			tile := t.Tiles[y][x]

			top := tile.Elevation
			for elevetion := 0; elevetion <= top; elevetion++ {
				e := float64(elevetion)

				op.GeoM.Reset()
				// Move to current isometric position.
				op.GeoM.Translate(xi, yi)

				op.GeoM.Translate(0, -64*e)

				op.GeoM.Scale(t.SX, t.SY)

				op.GeoM.Translate(t.CX, t.CY)

				tile.Draw(screen, op)
			}

			if top == 0 {
				op.GeoM.Reset()
				// Move to current isometric position.
				op.GeoM.Translate(xi, yi)

				op.GeoM.Scale(t.SX, t.SY)

				op.GeoM.Translate(t.CX, t.CY)

				screen.DrawImage(t.TrailHeadSprite, op)
			}

			if top == 9 {
				op.GeoM.Reset()
				// Move to current isometric position.
				op.GeoM.Translate(xi, yi)

				op.GeoM.Translate(0, -64*11)

				op.GeoM.Scale(t.SX, t.SY)

				op.GeoM.Translate(t.CX, t.CY)

				screen.DrawImage(t.FlagSprite, op)
			}
		}

		if animationState == 0 {
			if _maxX < t.Width {
				mapperX[y] = _maxX + 1
			}

			if _maxX == t.Width && y == maxY-1 && maxY < t.Height {
				maxY++
			}
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("WORLD %.3f, %.3f", t.WorldWidth, t.WorldHeight))
}

func (t Topography) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (t Topography) Update() error {
	return nil
}

func Animate2() {
	ebiten.SetWindowTitle("Isometric (Ebitengine Demo)")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	input := utils.ReadInput("day_10/input")

	topography := NewTopography(input)

	topography.CX = float64(ScreenWidth / 2.05)
	topography.CY = float64(ScreenHeight / 5)
	topography.SX = ScreenWidth / topography.WorldWidth
	topography.SY = ScreenHeight / topography.WorldHeight

	if err := ebiten.RunGame(topography); err != nil {
		log.Fatal(err)
	}
	// topography.RateTrailHeadsPart2()

	// total := topography.GetTrailHeadsTotalScore()

	// log.Print(total)

	// return total
}
