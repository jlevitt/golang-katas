package day2

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jlevitt/katas/advent/input"
)

func PartOne(path string) error {
	lines := input.ReadInputLines(path)
	var valid []string

	for i, l := range lines {
		line, err := parseLine(l)
		if err != nil {
			return fmt.Errorf("parsing line %v: %w", i, err)
		}

		rule := containsCharRule{
			min:  line.min,
			max:  line.max,
			char: line.char,
		}

		if rule.Matches(line.password) {
			valid = append(valid, line.password)
		}
	}

	fmt.Printf("Found %v valid lines:\n", len(valid))

	for _, v := range valid {
		fmt.Printf("%v\n", v)
	}

	return nil
}

type parsedLine struct {
	min, max       int
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
		min:      min,
		max:      max,
		char:     matches[3],
		password: matches[4],
	}, nil
}
