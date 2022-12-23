package jungle

import "fmt"

func countLowerNumbers(arr []int, index int) int {
	count := 0
	if index == 0 {
		return 0
	}
	for i := index - 1; i >= 0; i-- {
		count += 1
		if arr[i] >= arr[index] {
			break
		}
	}

	return count
}

func reverse(input []int) []int {
	s := duplicateSlice(input)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func transpose(input [][]int) [][]int {
	result := make([][]int, len(input[0]))
	for _, row := range input {
		for j, value := range row {
			result[j] = append(result[j], value)
		}
	}
	return result
}

func duplicateSlice(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func toKey(a, b int) string {
	return fmt.Sprintf("%d,%d", a, b)
}
