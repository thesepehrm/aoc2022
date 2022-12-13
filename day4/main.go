package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/thesepehrm/aoc2022/day4/processor"
)

var (
	total = 0
)

func process(line string) {
	pairs := strings.Split(line, ",")
	elf1 := strings.Split(pairs[0], "-")
	elf2 := strings.Split(pairs[1], "-")

	result := processor.CheckPartialOverlap(elf1, elf2)

	if result {
		total += 1
	}
}

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		process(fileScanner.Text())
	}

	readFile.Close()
	println(total)
}
