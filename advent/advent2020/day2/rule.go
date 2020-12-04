package day2

import "strings"

type rule interface {
	matches(string) bool
}

type containsCharRule struct {
	min, max int
	char     string
}

func (c containsCharRule) Matches(s string) bool {
	count := strings.Count(s, c.char)

	return c.min <= count && count <= c.max
}
