package main

import (
	_ "embed"
	"log"
)

//go:embed input
var input string

//go:embed sample
var sample string

func main() {
	tests := []struct {
		fileName string
		bytes    int
		size     int
		want1    int
		want2    string
	}{
		{sample, 12, 6, 22, "6,1"},
		{input, 1024, 70, 272, "16,44"},
	}

	log.Println("========== Part 1 Tests ==========")
	for _, test := range tests {
		got := Part1(test.fileName, test.bytes, test.size)
		if got != test.want1 {
			log.Fatalf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want1)
		}
		log.Printf("Got: %d\n", got)
	}

	log.Println("========== Part 2 Tests ==========")
	for _, test := range tests {
		got := Part2(test.fileName, test.bytes, test.size)
		if got != test.want2 {
			log.Fatalf("Failed Test %s\n\tGot %s, Want %s\n", test.fileName, got, test.want2)
			continue
		}
		log.Printf("Got: %s\n", got)
	}
}
