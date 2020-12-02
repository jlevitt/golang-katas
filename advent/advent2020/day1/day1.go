package day1

import (
	"errors"
	"fmt"

	"github.com/jlevitt/katas/advent/input"
)

func PartTwo(path string) error {
	expenses, err := input.AsInts(input.ReadInputLines(path))
	if err != nil {
		return fmt.Errorf("reading input: %w", err)
	}

	result, err := get2020Expenses(expenses)
	if err != nil {
		return err
	}

	fmt.Printf("Result: %v\n", result)

	return nil
}

func get2020Expenses(expenses []int) (int, error) {
	for i, _ := range expenses {
		for j, _ := range expenses {
			for k, _ := range expenses {
				a := expenses[i]
				b := expenses[j]
				c := expenses[k]

				if a+b+c == 2020 {
					return a * b * c, nil
				}
			}
		}
	}

	return -1, errors.New("2020 expenses not found")
}
