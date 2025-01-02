package main

import (
	_ "embed"
	"log"
)

//go:embed input
var input string

//go:embed sample1
var sample1 string

//go:embed sample2
var sample2 string

//go:embed sample3
var sample3 string

func main() {
	tests := []struct {
		fileName string
		want1    int
		want2    int
	}{
		{sample1, 7036, 45},
		{sample2, 11048, 64},
		{input, 91464, 0},
	}

	log.Println("Day 16: Part 1 Tests")
	for _, test := range tests {
		got := Part1(test.fileName)
		if got != test.want1 {
			log.Fatalf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want1)
		}
		log.Printf("Got: %d\n", got)
	}

	log.Println("Day 16: Part 2 Tests")
	for _, test := range tests {
		got := Part2(test.fileName)
		if got != test.want2 {
			log.Fatalf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want2)
			continue
		}
		log.Printf("Got: %d\n", got)
	}
}
