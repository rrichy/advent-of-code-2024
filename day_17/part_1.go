package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/rrichy/advent-of-code-2024/utils"
)

type Program struct {
	RegisterA    int64
	RegisterB    int64
	RegisterC    int64
	Pointer      int
	Instructions []int
	Output       []int
}

func NewProgram(input string) Program {
	p := Program{
		Output:       []int{},
		Instructions: []int{},
	}

	splits := strings.Split(input, "\n\n")
	stringRegisters := strings.Split(splits[0], "\n")
	registers := []int{}

	for _, register := range stringRegisters {
		splits := strings.Split(register, ": ")
		registers = append(registers, utils.MustAtoi(splits[1]))
	}

	programs := strings.Split(splits[1], ": ")

	for _, instruction := range strings.Split(programs[1], ",") {
		p.Instructions = append(p.Instructions, utils.MustAtoi(instruction))
	}

	p.RegisterA = int64(registers[0])
	p.RegisterB = int64(registers[1])
	p.RegisterC = int64(registers[2])

	return p
}

func (p *Program) getComboOperand(operand int64) int64 {
	if operand == 4 {
		return p.RegisterA
	}

	if operand == 5 {
		return p.RegisterB
	}

	if operand == 6 {
		return p.RegisterC
	}

	return operand
}

func (p *Program) adv(operand int64) {
	p.RegisterA >>= operand
}

func (p *Program) bxl(operand int64) {
	p.RegisterB ^= operand
}

func (p *Program) bst(operand int64) {
	p.RegisterB = operand & 0b111
}

func (p *Program) jnz(operand int64) {
	if p.RegisterA != 0 {
		p.Pointer = int(operand)
	} else {
		p.Pointer += 2
	}
}

func (p *Program) bxc() {
	p.RegisterB ^= p.RegisterC
}

func (p *Program) out(operand int64) {
	p.Output = append(p.Output, int(operand&0b111))
}

func (p *Program) bdv(operand int64) {
	p.RegisterB = p.RegisterA >> operand
}

func (p *Program) cdv(operand int64) {
	p.RegisterC = p.RegisterA >> operand
}

func (p *Program) run() {
	instructionLength := len(p.Instructions)

	for p.Pointer < instructionLength {
		opcode := p.Instructions[p.Pointer]
		literalOperand := int64(p.Instructions[p.Pointer+1])

		switch opcode {
		case 0:
			p.adv(p.getComboOperand(literalOperand))
		case 1:
			p.bxl(literalOperand)
		case 2:
			p.bst(p.getComboOperand(literalOperand))
		case 3:
			p.jnz(p.getComboOperand(literalOperand))
			continue
		case 4:
			p.bxc()
		case 5:
			p.out(p.getComboOperand(literalOperand))
		case 6:
			p.bdv(p.getComboOperand(literalOperand))
		case 7:
			p.cdv(p.getComboOperand(literalOperand))
		default:
		}

		p.Pointer += 2
	}
}

func Part1(input string) string {
	defer func(t time.Time) {
		log.Println("time", time.Since(t))
	}(time.Now())

	p := NewProgram(input)
	p.run()

	output := []string{}
	for _, o := range p.Output {
		output = append(output, strconv.Itoa(o))
	}

	return strings.Join(output, ",")
}
