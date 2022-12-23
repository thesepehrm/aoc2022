package main

import (
	"bufio"
	"os"

	"github.com/thesepehrm/aoc2022/day7/cli"
	"github.com/thesepehrm/aoc2022/day7/fs"
)

func process(line string) {

}

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sepehrOS := cli.NewCommandLine()

	for fileScanner.Scan() {
		sepehrOS.Parse(fileScanner.Text())
	}

	readFile.Close()
	println(fs.TotalFreeableSpace(sepehrOS.Fs.Root))
	println(fs.SmallestSubdirectoryToDelete(sepehrOS.Fs.Root))

}
