package main

import (
	"fmt"
	"github.com/jlevitt/katas/advent2019"
	"github.com/jlevitt/katas/advent2019/day2"
	"log"
)

func main() {
	partOne()
}

func partOne() {
	lines := advent2019.ReadInputLines()
	p, err := day2.ParseProgram(lines[0])
	if err != nil {
		log.Fatal(err)
	}

	day2.SetProgramAlarm(p)

	program, err := day2.RunProgram(p)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(program)
}
