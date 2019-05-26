package stack

type stack struct {
	values []int
}

func NewStack() stack {
	return stack{}
}

func (s *stack) Push(value int) {
	s.values = append(s.values, value)
}

func (s *stack) Pop() (int, bool) {
	if len(s.values) == 0 {
		return 0, false
	}

	value := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]

	return value, true
}

func (s *stack) Size() int {
	return len(s.values)
}
