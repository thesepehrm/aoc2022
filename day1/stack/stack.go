package stack

type Stack struct {
	data []int
	cap  int
}

func NewStack(cap int, data []int) *Stack {
	return &Stack{
		cap:  cap,
		data: data,
	}
}

func (s *Stack) PushAt(index int, num int) {
	if index >= len(s.data) {
		return
	}

	newArray := []int{}

	for i := 0; i < index; i++ {
		newArray = append(newArray, s.data[i])
	}
	newArray = append(newArray, num)
	newArray = append(newArray, s.data[index:]...)
	s.data = newArray[:s.cap]
}

func (s *Stack) Array() []int {
	return s.data
}
