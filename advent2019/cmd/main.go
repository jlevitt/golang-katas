package main

import (
	"flag"
	"github.com/jlevitt/katas/advent2019/day1"
	"github.com/jlevitt/katas/advent2019/day2"
	"log"
)

type key struct {
	day, part int
}

func main() {
	var day, part int

	flag.IntVar(&day, "day", 1, "day 1 or 2")
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	dayPartLookup := map[key]func(){
		key{day: 1, part: 2}: day1.PartTwo,
		key{day: 2, part: 1}: day2.PartOne,
		key{day: 2, part: 2}: day2.PartTwo,
	}

	if f, ok := dayPartLookup[key{day, part}]; ok {
		f()
	} else {
		log.Fatal("advent day/part not found - part is either 1 or 2 and day is 1-25")
	}
}

