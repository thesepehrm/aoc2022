package crates

type Instruction []int

func (i Instruction) Crates() int {
	return i[0]
}

func (i Instruction) FromIndex() int {
	return i[1] - 1
}

func (i Instruction) ToIndex() int {
	return i[2] - 1
}
