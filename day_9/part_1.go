package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rrichy/advent-of-code-2024/utils"
)

type Free struct {
	Index int
	Size  int
}

type Block struct {
	Id    int
	Size  int
	Index int
}

type DiskMap struct {
	Cells  []*int
	Frees  []*Free
	Blocks []*Block
}

func (d *DiskMap) AddBlocks(id int, size int) {
	d.Blocks = append(d.Blocks, &Block{Id: id, Size: size, Index: len(d.Cells)})
	for i := 0; i < size; i++ {
		d.Cells = append(d.Cells, &id)
	}
}

func (d *DiskMap) Expand(size int) {
	d.Frees = append(d.Frees, &Free{Index: len(d.Cells), Size: size})
	for i := 0; i < size; i++ {
		d.Cells = append(d.Cells, nil)
	}
}

func (d *DiskMap) FSFragmentation() {
	leftPointer := 0
	rightPointer := len(d.Cells) - 1

	for leftPointer < rightPointer {
		if (d.Cells)[rightPointer] != nil {
			for leftPointer < rightPointer {
				if (d.Cells)[leftPointer] == nil {
					(d.Cells)[leftPointer], (d.Cells)[rightPointer] = (d.Cells)[rightPointer], (d.Cells)[leftPointer]
					break
				}

				leftPointer++
			}
		}

		rightPointer--
	}
}

func (d DiskMap) Print() {
	temp := ""
	for _, block := range d.Cells {

		if block == nil {
			temp += "."
		} else {
			temp += strconv.Itoa(*block)
		}
	}

	fmt.Println(temp)
}

func Part1() int {
	input := utils.ReadInput("day_9/input")

	diskMap := DiskMap{}
	currentId := 0
	diskMapPointer := 0
	for i, block := range strings.Split(input, "") {
		size, _ := strconv.Atoi(block)
		if i%2 == 0 {
			diskMap.AddBlocks(currentId, size)
			currentId++
		} else {
			diskMap.Expand(size)
		}
		diskMapPointer += size
	}

	diskMap.FSFragmentation()

	sum := 0
	for i, block := range diskMap.Cells {
		if block != nil {
			sum += i * *block
		}
	}

	fmt.Printf("Sum: %d\n", sum)

	return sum
}
