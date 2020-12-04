package day2

import "strings"

type rule = func(line parsedLine) bool

func containsChar(line parsedLine) bool {
	count := strings.Count(line.password, line.char)

	return line.a <= count && count <= line.b
}
