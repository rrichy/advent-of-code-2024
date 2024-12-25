package resources

import (
	_ "embed"
)

var (
	//go:embed "Resources/Sheep/HappySheep_All.png"
	HappySheepAll []byte

	//go:embed "Terrain/Water/Rocks/Rocks_01.png"
	Rocks1 []byte

	//go:embed "Terrain/Water/Rocks/Rocks_02.png"
	Rocks2 []byte

	//go:embed "Terrain/Water/Rocks/Rocks_03.png"
	Rocks3 []byte

	//go:embed "Terrain/Water/Rocks/Rocks_04.png"
	Rocks4 []byte

	//go:embed "Factions/Knights/Troops/Warrior/Blue/Warrior_Blue.png"
	Warrior []byte

	//go:embed "Terrain/Ground/Tilemap_Flat.png"
	Tilemap_Flat []byte

	//go:embed "Terrain/Ground/Tilemap_Elevation.png"
	Tilemap_Elevation []byte

	//go:embed "Terrain/Water/Water.png"
	Water []byte

	//go:embed "Terrain/Water/Foam/Foam.png"
	Foam []byte
)
