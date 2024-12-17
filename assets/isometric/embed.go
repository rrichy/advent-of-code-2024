package isometric

import (
	_ "embed"
)

var (
	//go:embed "256x192 Tiles.png"
	Tiles []byte

	//go:embed "256x128 Tile Overlays.png"
	TileOverlays []byte

	//go:embed "256x256 Objects.png"
	Objects []byte
)
