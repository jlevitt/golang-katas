package advent2019

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadInput() []int{
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var result []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reading, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, reading)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}