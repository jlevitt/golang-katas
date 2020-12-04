package day2

import (
	"fmt"

	"github.com/jlevitt/katas/advent/advent2020/day2/internal/parse"
	"github.com/jlevitt/katas/advent/advent2020/day2/internal/rules"
	"github.com/jlevitt/katas/advent/input"
)

func PartOne(path string) error {
	return run(path, rules.CharRepeated)
}

func PartTwo(path string) error {
	return run(path, rules.CharPositionedAt)
}

func run(path string, r rules.Rule) error {
	lines := input.ReadInputLines(path)
	valid, err := countValidLines(lines, r)
	if err != nil {
		return err
	}

	for _, v := range valid {
		fmt.Printf("%v\n", v)
	}

	fmt.Printf("\nFound %v valid lines:\n", len(valid))

	return nil
}

func countValidLines(lines []string, r rules.Rule) ([]string, error) {
	var valid []string

	for i, l := range lines {
		line, err := parse.NewLine(l)
		if err != nil {
			return nil, fmt.Errorf("parsing line %v: %w", i+1, err)
		}

		if r(line) {
			valid = append(valid, line.Password)
		}
	}
	return valid, nil
}
