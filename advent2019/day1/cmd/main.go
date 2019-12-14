package main

import (
	"fmt"
	"github.com/jlevitt/katas/advent2019"
	"github.com/jlevitt/katas/advent2019/day1"
)

func main() {
	moduleMasses := advent2019.ReadInput()

	totalFuelRequired := 0
	for _, mass := range moduleMasses {
		totalFuelRequired += day1.FuelRequired(mass)
	}

	fmt.Printf("Total fuel required: %v\n", totalFuelRequired)
}
