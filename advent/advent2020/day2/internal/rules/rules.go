package rules

import (
	"strings"

	"github.com/jlevitt/katas/advent/advent2020/day2/internal/parse"
)

type Rule = func(line parse.Line) bool

func CharRepeated(line parse.Line) bool {
	count := strings.Count(line.Password, line.Char)

	return line.A <= count && count <= line.B
}

func CharPositionedAt(line parse.Line) bool {
	if len(line.Password) < line.A || len(line.Password) < line.B {
		return false
	}

	positionAMatches := line.Password[line.A-1:line.A] == line.Char
	positionBMatches := line.Password[line.B-1:line.B] == line.Char

	return (positionAMatches || positionBMatches) && !(positionAMatches && positionBMatches)
}
