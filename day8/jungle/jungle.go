package jungle

import (
	"strconv"
	"strings"
)

// y,x
type Jungle struct {
	trees              [][]int
	visibleTreesCoords map[string]bool
}

func NewJungle() *Jungle {
	return &Jungle{
		trees:              [][]int{},
		visibleTreesCoords: make(map[string]bool),
	}
}

func (j *Jungle) AddRowOfTrees(ts []int) {
	j.trees = append(j.trees, ts)
}

func (j *Jungle) VisibleTreesCount() int {

	count := 0

	// left and right
	for y, row := range j.trees {
		total := visibleTreesFromOneSide(row, false)
		otherSide := visibleTreesFromOneSide(reverse(row), true)
		total = append(total, otherSide...)
		for _, t := range total {
			j.visibleTreesCoords[toKey(y, t)] = true
		}
	}

	// up and down
	tj := transpose(j.trees)
	for y, row := range tj {
		total := visibleTreesFromOneSide(row, false)
		otherSide := visibleTreesFromOneSide(reverse(row), true)
		total = append(total, otherSide...)
		for _, t := range total {
			j.visibleTreesCoords[toKey(t, y)] = true
		}
	}

	count += len(j.visibleTreesCoords)

	return count

}

func (j Jungle) MaxScenicScore() int {
	max := 0
	for k := range j.visibleTreesCoords {
		raw := strings.Split(k, ",")
		x, _ := strconv.Atoi(raw[1])
		y, _ := strconv.Atoi(raw[0])
		s := j.calculateScore(x, y)

		if s > max {
			max = s
		}

	}

	return max
}

func (j Jungle) calculateScore(x, y int) int {
	bound := len(j.trees)
	tj := transpose(j.trees)

	left := countLowerNumbers(j.trees[y], x)
	right := countLowerNumbers(reverse(j.trees[y]), bound-x-1)
	down := countLowerNumbers(tj[x], y)
	up := countLowerNumbers(reverse(tj[x]), bound-y-1)

	//println(toKey(x, y), j.trees[y][x], left, right, down, up, left*right*down*up)
	return left * right * down * up
}

func visibleTreesFromOneSide(row []int, reversed bool) []int {
	maxHeight := -1
	visibleIndices := []int{}

	for i, t := range row {
		if t > maxHeight {
			maxHeight = t
			if reversed {
				visibleIndices = append(visibleIndices, len(row)-i-1)
			} else {
				visibleIndices = append(visibleIndices, i)
			}

		}
	}

	return visibleIndices
}
