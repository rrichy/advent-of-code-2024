package main

import _ "embed"

//go:embed input
var input string

//go:embed sample1
var sample1 string

//go:embed sample2
var sample2 string

//go:embed sample3
var sample3 string

func main() {
	Part1()
	Part2()
}
