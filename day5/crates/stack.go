package crates

type Stack []string

func (s Stack) Pop() (string, Stack) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}

func (s Stack) Push(crate string) Stack {
	return append(Stack{crate}, s...)
}

func (s Stack) MultiplePush(crates Stack) Stack {
	cratesCopy := make(Stack, len(crates))
	copy(cratesCopy, crates)
	return append(cratesCopy, s...)
}

func (s Stack) MultiplePop(times int) (popped Stack, remaining Stack) {
	popped = s[:times]
	remaining = s[times:]
	return
}

func (s Stack) Top() string {
	if len(s) == 0 {
		return ""
	}
	return s[0]
}
