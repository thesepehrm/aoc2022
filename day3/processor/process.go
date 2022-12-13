package processor

import (
	"strings"
	"sync"
)

func GetPriority(firstBatch string, secondBatch string) int {
	item := Item(findRepeatedCharacter(firstBatch, secondBatch))
	return item.Priority()
}

func FindPriorityInGroups(first, second, third string) int {
	i := findIntersection(first, second)
	item := Item(findRepeatedCharacter(third, i))
	return item.Priority()
}

func findRepeatedCharacter(a string, b string) rune {
	wg := sync.WaitGroup{}
	var repeatedCharacter rune
	var found bool

	for _, v := range a {
		wg.Add(1)
		go func(v rune) {
			defer wg.Done()
			for _, k := range b {
				if found {
					return
				}
				if v == k {
					repeatedCharacter = v
					found = true
				}
			}
		}(v)
	}
	wg.Wait()

	return repeatedCharacter
}

func findIntersection(a string, b string) string {
	x := ""
	for _, v := range a {
		if strings.ContainsRune(b, v) {
			x += string(v)
		}
	}
	return x
}
