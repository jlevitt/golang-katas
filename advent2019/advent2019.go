package advent2019

import (
	"bufio"
	"github.com/pkg/errors"
	"log"
	"os"
)

func ReadInputLines(path string) []string{
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		result = append(result, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func Min(values []int) (int, error) {
	if len(values) == 0 {
		return 0, errors.New("values is empty")
	}

	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}

	return min, nil
}