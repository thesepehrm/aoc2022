package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/thesepehrm/aoc2022/day8/jungle"
)

var (
	j = jungle.NewJungle()
)

func process(line string) {
	characters := strings.Split(line, "")
	nums := []int{}
	for _, c := range characters {
		n, _ := strconv.Atoi(string(c))
		nums = append(nums, n)
	}
	j.AddRowOfTrees(nums)
}

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		process(fileScanner.Text())
	}

	readFile.Close()
	println(j.VisibleTreesCount())
	println(j.MaxScenicScore())

}
