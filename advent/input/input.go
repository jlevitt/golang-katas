package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadInputLines(path string) []string {
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

func AsInts(strs []string) ([]int, error) {
	var result []int
	for _, str := range strs {
		i, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}
