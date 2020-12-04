package parse

import (
	"fmt"
	"regexp"
	"strconv"
)

type Line struct {
	A, B           int
	Char, Password string
}

var pattern *regexp.Regexp = regexp.MustCompile("([[:digit:]]+)-([[:digit:]]+) (.): (.+)")

func NewLine(line string) (Line, error) {
	matches := pattern.FindStringSubmatch(line)

	if len(matches) != 5 {
		return Line{}, fmt.Errorf("line didn't match expected format: %v", line)
	}

	min, err := strconv.Atoi(matches[1])
	if err != nil {
		return Line{}, err
	}

	max, err := strconv.Atoi(matches[2])
	if err != nil {
		return Line{}, err
	}

	return Line{
		A:        min,
		B:        max,
		Char:     matches[3],
		Password: matches[4],
	}, nil
}
