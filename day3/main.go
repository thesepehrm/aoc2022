package main

import (
	"bufio"
	"os"

	"github.com/thesepehrm/aoc2022/day3/processor"
)

var (
	total = 0
)

func process(line string) {
	lineLength := len(line)
	firstHalf := line[:lineLength/2]
	secondHalf := line[lineLength/2:]
	total += processor.GetPriority(firstHalf, secondHalf)
}

func processPart2(groups ...string) {
	total += processor.FindPriorityInGroups(groups[0], groups[1], groups[2])
}

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		first := fileScanner.Text()
		fileScanner.Scan()
		second := fileScanner.Text()
		fileScanner.Scan()
		third := fileScanner.Text()

		processPart2(first, second, third)
	}

	readFile.Close()
	println(total)
}
