package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	result := Part1()
	assert.Equal(t, result, 12553187650171)
}

func TestPart2(t *testing.T) {
	result := Part2()
	assert.Equal(t, result, 96779702119491)
}