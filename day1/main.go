package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	currentTotal = 0
	maxTotal     = 0
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

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		process(fileScanner.Text())
	}

	readFile.Close()
	println(maxTotal)
}
