package day2

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jlevitt/katas/advent/input"
)

func PartOne(path string) error {
	lines := input.ReadInputLines(path)
	valid, err := countValidLines(lines, containsChar)
	if err != nil {
		return err
	}

	for _, v := range valid {
		fmt.Printf("%v\n", v)
	}

	fmt.Printf("\nFound %v valid lines:\n", len(valid))

	return nil
}

func countValidLines(lines []string, r rule) ([]string, error) {
	var valid []string

	for i, l := range lines {
		line, err := parseLine(l)
		if err != nil {
			return nil, fmt.Errorf("parsing line %v: %w", i, err)
		}

		if r(line) {
			valid = append(valid, line.password)
		}
	}
	return valid, nil
}

type parsedLine struct {
	a, b           int
	char, password string
}

var pattern *regexp.Regexp = regexp.MustCompile("([[:digit:]]+)-([[:digit:]]+) (.): (.+)")

func parseLine(line string) (parsedLine, error) {
	matches := pattern.FindStringSubmatch(line)

	if len(matches) != 5 {
		return parsedLine{}, fmt.Errorf("line didn't match expected format: %v", line)
	}

	min, err := strconv.Atoi(matches[1])
	if err != nil {
		return parsedLine{}, err
	}

	max, err := strconv.Atoi(matches[2])
	if err != nil {
		return parsedLine{}, err
	}

	return parsedLine{
		a:        min,
		b:        max,
		char:     matches[3],
		password: matches[4],
	}, nil
}
