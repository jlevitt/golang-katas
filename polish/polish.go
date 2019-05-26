package polish

import (
	"github.com/jlevitt/katas/polish/stack"
)

type operator func(float64, float64) float64

var ops = map[rune]operator{
	'+': func(l, r float64) float64 {
		return l + r
	},
	'-': func(l, r float64) float64 {
		return l - r
	},
	'*': func(l, r float64) float64 {
		return l * r
	},
	'/': func(l, r float64) float64 {
		return l / r
	},
}

func Evaluate(expr string) float64 {
	stack := stack.NewStack()

	for _, char := range expr {
		if op, ok := ops[char]; ok {
			r, _ := stack.Pop()
			l, _ := stack.Pop()
			stack.Push(op(l, r))
		} else if char == ' ' {
			continue
		} else {
			stack.Push(float64(char - '0'))
		}
	}

	result, _ := stack.Pop()

	return result
}
