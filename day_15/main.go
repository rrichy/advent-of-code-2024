package main

import _ "embed"

//go:embed input
var input string

// 9021
//
//go:embed sample1
var sample1 string

// 1751
//
//go:embed sample2
var sample2 string

// 618
//
//go:embed sample3
var sample3 string

//go:embed sample4
var sample4 string

func main() {
	Part1()
	Part2()
	Animate()
}
