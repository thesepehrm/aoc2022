package processor

import (
	"fmt"
	"strconv"
)

func CheckOverallOverlap(s1 []string, s2 []string) bool {
	r1 := NewRangeFromString(s1)
	r2 := NewRangeFromString(s2)

	maxRange := r1
	minRange := r2

	if r2.Length() > r1.Length() {
		maxRange = r2
		minRange = r1
	}

	return maxRange.LowerBound <= minRange.LowerBound && maxRange.UpperBound >= minRange.UpperBound
}

func CheckPartialOverlap(s1 []string, s2 []string) bool {
	r1 := NewRangeFromString(s1)
	r2 := NewRangeFromString(s2)

	higherRange := r1
	lowerRange := r2

	if r1.LowerBound < r2.LowerBound {
		lowerRange = r1
		higherRange = r2
	}
	return lowerRange.UpperBound >= higherRange.LowerBound
}

type Range struct {
	LowerBound int
	UpperBound int
}

func NewRangeFromString(s []string) Range {
	lowerBound, _ := strconv.Atoi(s[0])
	upperBound, _ := strconv.Atoi(s[1])

	return Range{
		LowerBound: lowerBound,
		UpperBound: upperBound,
	}
}

func (r Range) Length() int {
	return r.UpperBound - r.LowerBound
}

func (r Range) String() string {
	return fmt.Sprintf("%d-%d", r.LowerBound, r.UpperBound)
}
