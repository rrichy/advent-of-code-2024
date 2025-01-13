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

func main() {
	tests := []struct {
		fileName string
		want1    string
		want2    int
	}{
		{sample2, "5,7,3,0", 117440},
		{sample1, "4,6,3,5,6,3,5,2,1,0", 0},
		{input, "1,6,7,4,3,0,5,0,6", 216148338630253}, // 63687530
	}

	log.Println("========== Part 1 Tests ==========")
	for _, test := range tests {
		got := Part1(test.fileName)
		if got != test.want1 {
			log.Fatalf("Failed Test %s\n\tGot %s\nWant %s\n", test.fileName, got, test.want1)
		}
		log.Printf("Got: %s\n", got)
	}

	log.Println("========== Part 2 Tests ==========")
	for i, test := range tests {
		if i == 1 {
			continue
		}
		got := Part2(test.fileName)
		if got != test.want2 {
			log.Fatalf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want2)
			continue
		}
		log.Printf("Got: %d\n", got)
	}
}
