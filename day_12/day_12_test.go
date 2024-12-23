package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	result := Part1()
	assert.Equal(t, result, 1473620)
}

func TestPart2(t *testing.T) {
	result := Part2()
	assert.Equal(t, result, 902620)
}
