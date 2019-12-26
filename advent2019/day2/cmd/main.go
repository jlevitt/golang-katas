package main

import (
	"fmt"
	"github.com/jlevitt/katas/advent2019"
	"github.com/jlevitt/katas/advent2019/day2"
	"log"
)

func main() {
	lines := advent2019.ReadInputLines()
	program, err := day2.RunProgram(lines[0])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(program)
}
