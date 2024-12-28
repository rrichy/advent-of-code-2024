package utils

import (
	"bytes"
	"image"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type Coordinate struct {
	X int
	Y int
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func (r *Direction) RotateLeft() {
	switch *r {
	case Up:
		*r = Left
	case Down:
		*r = Right
	case Left:
		*r = Down
	case Right:
		*r = Up
	}
}

func (r *Direction) RotateRight() {
	switch *r {
	case Up:
		*r = Right
	case Down:
		*r = Left
	case Left:
		*r = Up
	case Right:
		*r = Down
	}
}

func SlicePop[T any](slice []T) (T, []T) {
	return slice[len(slice)-1], slice[:len(slice)-1]
}

func SliceMap[T, U any](slice []T, f func(T) U) []U {
	temp := []U{}
	for _, v := range slice {
		temp = append(temp, f(v))
	}

	return temp
}

func SliceReduce[T, K any](slice []T, f func(K, T) K, result K) K {
	for _, v := range slice {
		result = f(result, v)
	}

	return result
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return i
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ReadInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	input := string(data)

	return input
}

func SliceContains[T comparable](slice []T, s T) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func New2DStringMatrix(s string) [][]string {
	temp := [][]string{}
	for _, line := range strings.Split(s, "\n") {
		temp = append(temp, strings.Split(line, ""))
	}

	return temp
}

func ImageMustDecode(data []byte) image.Image {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func GetSpriteSheet(data []byte) *ebiten.Image {
	img := ImageMustDecode(data)
	return ebiten.NewImageFromImage(img)
}

func RandBool() bool {
	return rand.Intn(2) == 1
}
