package stack

type stack struct {
	values []float64
}

func NewStack() stack {
	return stack{}
}

func (s *stack) Push(value float64) {
	s.values = append(s.values, value)
}

func (s *stack) Pop() (float64, bool) {
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
