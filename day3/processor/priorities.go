package processor

type Item rune

func (i Item) Priority() int {
	if i <= 90 {
		return int(i) - 38
	}
	return int(i) - 96
}

func (i Item) String() string {
	return string(i)
}
