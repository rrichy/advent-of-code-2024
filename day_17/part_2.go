package main

import (
	"log"
	"time"

	"github.com/rrichy/advent-of-code-2024/utils"
)

func (p *Program) reset(registerA int64) {
	p.RegisterA = registerA
	p.RegisterB = 0
	p.RegisterC = 0
	p.Pointer = 0
	p.Output = []int{}
}

func Part2(input string) int {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	p := NewProgram(input)

	a := int64(0)
	for _i := range p.Instructions {
		i := len(p.Instructions) - _i - 1
		a <<= 3

		for {
			p.reset(a)
			p.run()

			if utils.SliceEqual(p.Output, p.Instructions[i:]) {
				break
			}

			a += 1
		}
	}

	return int(a)
}
