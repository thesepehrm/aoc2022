package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/thesepehrm/aoc2022/day1/stack"
)

var (
	currentTotal  = 0
	maxTotal      = 0
	topThreeElves = stack.NewStack(3, []int{0, 0, 0})
)

func process(line string) {
	if line == "" {
		currentTotal = 0
		return
	}

	item, _ := strconv.Atoi(line)
	currentTotal += item

	if maxTotal < currentTotal {
		maxTotal = currentTotal
	}

}

func processPart2(line string) {
	if line == "" {
		for i, max := range topThreeElves.Array() {
			if currentTotal < max {
				continue
			}
			topThreeElves.PushAt(i, currentTotal)
			break
		}
		currentTotal = 0
		return
	}

	item, _ := strconv.Atoi(line)
	currentTotal += item
}

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		//process(fileScanner.Text())
		processPart2(fileScanner.Text())
	}

	readFile.Close()
	//println(maxTotal)
	sum := 0
	for _, s := range topThreeElves.Array() {
		sum += s
	}
	println(sum)
}
