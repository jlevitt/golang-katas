package stack

type stack struct {
	values []int
	size   int
}

func NewStack() stack {
	return stack{}
}

func (s *stack) Push(value int) {
	s.size++
	s.values = append(s.values, value)
}

func (s *stack) Pop() (int, bool) {
	if s.size == 0 {
		return 0, false
	}

	value := s.values[s.size-1]
	s.values = s.values[:s.size-1]
	s.size--

	return value, true
}

func (s *stack) Size() int {
	return s.size
}
