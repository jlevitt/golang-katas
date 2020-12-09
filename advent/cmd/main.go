package main

import (
	"flag"
	"log"
	"strconv"

	year2019day1 "github.com/jlevitt/katas/advent/advent2019/day1"
	year2019day2 "github.com/jlevitt/katas/advent/advent2019/day2"
	year2019day3 "github.com/jlevitt/katas/advent/advent2019/day3"
	year2019day4 "github.com/jlevitt/katas/advent/advent2019/day4"
	year2020day1 "github.com/jlevitt/katas/advent/advent2020/day1"
	year2020day2 "github.com/jlevitt/katas/advent/advent2020/day2"
	year2020day7 "github.com/jlevitt/katas/advent/advent2020/day7"
)

type key struct {
	year, day, part int
}

func main() {
	var year, day, part int
	var path string

	flag.IntVar(&year, "year", 2019, "year")
	flag.IntVar(&day, "day", 1, "day 1 or 2")
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.StringVar(&path, "path", "", "path to the program's input")
	flag.Parse()

	if path == "" {
		path = "day" + strconv.Itoa(day) + ".txt"
	}

	dayPartLookup := map[key]func(string) error{
		key{year: 2019, day: 1, part: 2}: year2019day1.PartTwo,
		key{year: 2019, day: 2, part: 1}: year2019day2.PartOne,
		key{year: 2019, day: 2, part: 2}: year2019day2.PartTwo,
		key{year: 2019, day: 3, part: 1}: year2019day3.PartOne,
		key{year: 2019, day: 3, part: 2}: year2019day3.PartTwo,
		key{year: 2019, day: 4, part: 1}: year2019day4.PartOne,
		key{year: 2020, day: 1, part: 2}: year2020day1.PartTwo,
		key{year: 2020, day: 2, part: 1}: year2020day2.PartOne,
		key{year: 2020, day: 2, part: 2}: year2020day2.PartTwo,
		key{year: 2020, day: 7, part: 1}: year2020day7.PartOne,
	}

	if f, ok := dayPartLookup[key{year, day, part}]; ok {
		err := f(path)
		if err != nil {
			log.Fatalf("Got error: %\n", err)
		}
	} else {
		log.Fatal("advent year/day/part not found - part is either 1 or 2 and day is 1-25")
	}
}
