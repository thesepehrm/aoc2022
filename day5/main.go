package main

import (
	"bufio"
	"os"

	"github.com/thesepehrm/aoc2022/day5/crates"
)

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	inputParser := crates.NewInputParser()

	for fileScanner.Scan() {
		inputParser.Parse(fileScanner.Text())
	}

	readFile.Close()

	crane := crates.NewCrane(inputParser.ParsedStacks(), inputParser.ParsedInstructions())
	crane.Run()
	println(crane.Tops())
}
