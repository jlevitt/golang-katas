package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/jlevitt/katas/advent2019"
	"github.com/jlevitt/katas/advent2019/day2"
	"log"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	if part == 1 {
		partOne()
	} else if part == 2 {
		partTwo()
	} else {
		log.Fatal("expected part 1 or 2")
	}
}

func partOne() {
	p := loadProgram()

	day2.SetProgramAlarm(p)

	program, err := day2.RunProgram(p)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(program)
}

func partTwo() {
	initialProgram := loadProgram()

	const expectedOutput = 19690720

	noun, verb, err := searchForOutput(initialProgram, expectedOutput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found noun=%v, verb=%v\n", noun, verb)
}

func loadProgram() day2.Program {
	lines := advent2019.ReadInputLines()
	p, err := day2.ParseProgram(lines[0])
	if err != nil {
		log.Fatal(err)
	}

	return p
}

func searchForOutput(initialProgram day2.Program, expectedOutput int) (int, int, error){
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			p := make(day2.Program, len(initialProgram))
			copy(p, initialProgram)
			day2.SetNounVerb(p, noun, verb)
			_, err := day2.RunProgram(p)
			if err != nil {
				return -1, -1, err
			}
			output := p[0]

			if output == expectedOutput {
				return noun, verb, nil
			}
		}
	}

	return -1, -1, errors.New("expected output not found")
}
