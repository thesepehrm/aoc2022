package crates

import (
	"regexp"
	"strconv"
)

type ReadingState int8

const (
	InitialReading ReadingState = iota
	CraneReading
	InstructionsReading
)

var (
	instructionsRegex = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
)

type Parser struct {
	readingState ReadingState
	numOfStacks  int
	stacks       []Stack
	instructions []Instruction
}

func NewInputParser() *Parser {
	return &Parser{
		readingState: InitialReading,
		numOfStacks:  0,
		stacks:       []Stack{},
		instructions: []Instruction{},
	}
}

func (p *Parser) Parse(line string) {
	if line == "" {
		p.readingState = InstructionsReading
		return
	}

	switch p.readingState {
	case InitialReading:
		p.readingState = CraneReading
		p.numOfStacks = (len(line) + 1) / 4
		p.stacks = generateEmptyStacks(p.numOfStacks)
		fallthrough
	case CraneReading:
		for i := 0; i < len(line); i += 4 {
			letter := string(line[i+1])
			if letter == " " {
				continue
			}
			if _, err := strconv.Atoi(letter); err == nil {
				break
			}
			stackNumber := i / 4
			p.stacks[stackNumber] = append(p.stacks[stackNumber], letter)
		}
	case InstructionsReading:
		matches := instructionsRegex.FindAllStringSubmatch(line, -1)[0][1:]
		intMatches := []int{}
		for _, m := range matches {
			num, _ := strconv.Atoi(m)
			intMatches = append(intMatches, num)
		}

		p.instructions = append(p.instructions, intMatches)
	}

}

func (p Parser) ParsedStacks() []Stack {
	return p.stacks
}

func (p Parser) ParsedInstructions() []Instruction {
	return p.instructions
}

func generateEmptyStacks(size int) []Stack {
	m := []Stack{}
	for i := 0; i < size; i++ {
		m = append(m, Stack{})
	}
	return m
}
