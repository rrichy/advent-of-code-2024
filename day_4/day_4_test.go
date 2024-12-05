package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	result := Part1()
	assert.Equal(t, result, 2532)
}

func TestPart2(t *testing.T) {
	result := Part2()
	assert.Equal(t, result, 1941)
}
