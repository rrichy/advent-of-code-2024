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

func (t Topography) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	cx, cy := float64(ScreenWidth/2.05), float64(ScreenHeight/5)
	scaleX := ScreenWidth / t.WorldWidth
	scaleY := ScreenHeight / t.WorldHeight

	for y := 0; y < t.Height; y++ {
		for x := 0; x < t.Width; x++ {
			xi, yi := t.cartesianToIso(float64(x), float64(y))

			tile := t.Tiles[y][x]

			top := tile.Elevation
			for elevetion := 0; elevetion <= top; elevetion++ {
				e := float64(elevetion)

				op.GeoM.Reset()
				// Move to current isometric position.
				op.GeoM.Translate(xi, yi)

				op.GeoM.Translate(0, -64*e)

				op.GeoM.Scale(scaleX, scaleY)

				op.GeoM.Translate(cx, cy)

				tile.Draw(screen, op)
			}

			if top == 0 {
				op.GeoM.Reset()
				// Move to current isometric position.
				op.GeoM.Translate(xi, yi)

				op.GeoM.Scale(scaleX, scaleY)

				op.GeoM.Translate(cx, cy)

				screen.DrawImage(t.TrailHeadSprite, op)
			}

			if top == 9 {
				op.GeoM.Reset()
				// Move to current isometric position.
				op.GeoM.Translate(xi, yi)

				op.GeoM.Translate(0, -64*11)

				op.GeoM.Scale(scaleX, scaleY)

				op.GeoM.Translate(cx, cy)

				screen.DrawImage(t.FlagSprite, op)
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

	if err := ebiten.RunGame(topography); err != nil {
		log.Fatal(err)
	}
	// topography.RateTrailHeadsPart2()

	// total := topography.GetTrailHeadsTotalScore()

	// log.Print(total)

	// return total
}
