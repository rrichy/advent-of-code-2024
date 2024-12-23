package main

import _ "embed"

//go:embed input
var input string

func main() {
	Part1()
	Part2()
}
