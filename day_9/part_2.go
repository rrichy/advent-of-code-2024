package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rrichy/advent-of-code-2024/utils"
)

func (d *DiskMap) Defrag() {
	for i := range d.Blocks {
		block := d.Blocks[len(d.Blocks)-1-i]
		for _, free := range d.Frees {
			if free.Size == 0 || free.Index > block.Index {
				continue
			}

			if free.Size >= block.Size {
				for j := 0; j < block.Size; j++ {
					d.Cells[free.Index+j] = &block.Id
					d.Cells[block.Index+j] = nil
				}

				block.Index = free.Index
				free.Size -= block.Size
				free.Index += block.Size

				break
			}
		}
	}

}

func Part2() int {
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

	diskMap.Defrag()

	sum := 0
	for i, block := range diskMap.Cells {
		if block != nil {
			sum += i * *block
		}
	}

	fmt.Printf("Sum: %d\n", sum)

	return sum
}
