package advent2019

import (
	"bufio"
	"log"
	"os"
)

func ReadInputLines() []string{
	file, err := os.Open("input.txt")
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