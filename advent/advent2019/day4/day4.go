package day4

import (
	"fmt"
	"strconv"
)

func PartOne(path string) error {
	min := 193651
	max := 649729

	matches := 0

	for i := min; i < max; i++ {
		digits := getDigits(i)
		if hasRun(digits) && isIncreasing(digits) {
			fmt.Printf("%v\n", i)
			matches++
		}
	}

	fmt.Printf("Total matches: %v\n", matches)

	return nil
}

func PartTwo(path string) error {
	return nil
}

func hasRun(num []int) bool {
	lastNum := -1
	for _, n := range num {
		if n == lastNum {
			return true
		}
		lastNum = n
	}

	return false
}

func isIncreasing(num []int) bool {
	lastNum := -1
	for _, n := range num {
		if n < lastNum {
			return false
		}
		lastNum = n
	}

	return true
}

func getDigits(num int) []int {
	var result []int

	str := strconv.Itoa(num)
	for _, c := range str {
		dStr := string(c)
		digit, _ := strconv.Atoi(dStr)
		result = append(result, digit)
	}

	return result
}
