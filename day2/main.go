package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/thesepehrm/aoc2022/day2/game"
)

var (
	totalPoints = 0
)

func process(line string) {
	values := strings.Split(line, " ")
	//totalPoints += game.RPS(values[0], values[1])
	totalPoints += game.RPSPart2(values[0], values[1])
}

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		process(fileScanner.Text())
	}

	readFile.Close()
	println(totalPoints)
}
