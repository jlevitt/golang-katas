package main

import (
	"flag"
	"github.com/jlevitt/katas/advent2019/day1"
	"github.com/jlevitt/katas/advent2019/day2"
	"github.com/jlevitt/katas/advent2019/day3"
	"github.com/jlevitt/katas/advent2019/day4"
	"log"
)

type key struct {
	day, part int
}

func main() {
	var day, part int
	var path string

	flag.IntVar(&day, "day", 1, "day 1 or 2")
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.StringVar(&path, "path", "input.txt", "path to the program's input")
	flag.Parse()

	dayPartLookup := map[key]func(string){
		key{day: 1, part: 2}: day1.PartTwo,
		key{day: 2, part: 1}: day2.PartOne,
		key{day: 2, part: 2}: day2.PartTwo,
		key{day: 3, part: 1}: day3.PartOne,
		key{day: 3, part: 2}: day3.PartTwo,
		key{day: 4, part: 1}: day4.PartOne,
	}

	if f, ok := dayPartLookup[key{day, part}]; ok {
		f(path)
	} else {
		log.Fatal("advent day/part not found - part is either 1 or 2 and day is 1-25")
	}
}

