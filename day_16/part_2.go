package main

import (
	"log"
	"time"
)

func Part2() int {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	return 0
}
