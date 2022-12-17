package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	MarkerLength = 14
)

var (
	lastIndex = 0
)

func process(line string) {
	marker := ""

	for i, char := range line {
		if len(marker) == MarkerLength {
			lastIndex = i
			break
		}

		_, after, ok := strings.Cut(marker, string(char))

		if ok {
			marker = after
		}

		marker += string(char)
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
	println(lastIndex)
}
