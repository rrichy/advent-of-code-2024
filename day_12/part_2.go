package main

import (
	_ "embed"
	"log"
	"time"
)

func Part2() int {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	garden := NewGarden(input)
	garden.FindRegions()

	sum := 0
	for _, regions := range garden.Plants {
		for _, region := range regions {
			sum += region.Area * region.Sides
		}
	}

	log.Print(sum)

	return sum
}
